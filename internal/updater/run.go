package updater

import (
	"context"

	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/models"
)

type Runner interface {
	Run(ctx context.Context, done chan<- struct{})
}

func (l *Loop) Run(ctx context.Context, done chan<- struct{}) {
	defer close(done)

	select {
	case <-l.start:
	case <-ctx.Done():
		return
	}

	for ctx.Err() == nil {
		updateCtx, updateCancel := context.WithCancel(ctx)

		serversCh := make(chan models.AllServers)
		errorCh := make(chan error)

		go func() {
			servers, err := l.updater.UpdateServers(updateCtx)
			if err != nil {
				errorCh <- err
				return
			}
			serversCh <- servers
		}()

		if l.userTrigger {
			l.userTrigger = false
			l.running <- constants.Running
		} else { // crash
			l.backoffTime = defaultBackoffTime
			l.statusManager.SetStatus(constants.Running)
		}

		stayHere := true
		for stayHere {
			select {
			case <-ctx.Done():
				updateCancel()
				<-errorCh
				close(errorCh)
				close(serversCh)
				return
			case <-l.start:
				l.userTrigger = true
				l.logger.Info("starting")
				stayHere = false
			case <-l.stop:
				l.userTrigger = true
				l.logger.Info("stopping")
				updateCancel()
				<-errorCh
				l.stopped <- struct{}{}
			case servers := <-serversCh:
				l.setAllServers(servers)
				if err := l.storage.FlushToFile(servers); err != nil {
					l.logger.Error(err.Error())
				}
				l.statusManager.SetStatus(constants.Completed)
				l.logger.Info("Updated servers information")
			case err := <-errorCh:
				updateCancel()
				close(serversCh)
				close(errorCh)
				l.statusManager.SetStatus(constants.Crashed)
				l.logAndWait(ctx, err)
				stayHere = false
			}
		}
		updateCancel()
	}
}
