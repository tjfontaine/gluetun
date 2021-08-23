package surfshark

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/qdm12/gluetun/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_addServersFromAPI(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		hts            hostToServer
		responseStatus int
		responseBody   io.ReadCloser
		expected       hostToServer
		err            error
	}{
		"fetch API error": {
			responseStatus: http.StatusNoContent,
			err:            errors.New("HTTP status code not OK: 204 No Content"),
		},
		"success": {
			hts: hostToServer{
				"existinghost": {Hostname: "existinghost"},
			},
			responseStatus: http.StatusOK,
			responseBody: ioutil.NopCloser(strings.NewReader(`[
				{"connectionName":"host1","region":"region1","country":"country1","location":"location1"},
				{"connectionName":"host1","region":"region1","country":"country1","location":"location1","pubKey":"some key"},
				{"connectionName":"host2","region":"region2","country":"country1","location":"location2","pubKey":"some key"},
				{"connectionName":"host3","region":"region3","country":"country3","location":"location3"}
			]`)),
			expected: map[string]models.SurfsharkServer{
				"existinghost": {Hostname: "existinghost"},
				"host1": {
					OpenVPN:   true,
					Wireguard: true,
					Region:    "region1",
					Country:   "country1",
					City:      "location1",
					Hostname:  "host1",
					TCP:       true,
					UDP:       true,
					WgPubKey:  "some key",
				},
				"host2": {
					OpenVPN:   true,
					Wireguard: true,
					Region:    "region2",
					Country:   "country1",
					City:      "location2",
					Hostname:  "host2",
					TCP:       true,
					UDP:       true,
					WgPubKey:  "some key",
				},
				"host3": {
					OpenVPN:  true,
					Region:   "region3",
					Country:  "country3",
					City:     "location3",
					Hostname: "host3",
					TCP:      true,
					UDP:      true,
				}},
		},
	}
	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			client := &http.Client{
				Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
					assert.Equal(t, http.MethodGet, r.Method)
					assert.Equal(t, r.URL.String(), "https://my.surfshark.com/vpn/api/v4/server/clusters")
					return &http.Response{
						StatusCode: testCase.responseStatus,
						Status:     http.StatusText(testCase.responseStatus),
						Body:       testCase.responseBody,
					}, nil
				}),
			}

			err := addServersFromAPI(ctx, client, testCase.hts)

			assert.Equal(t, testCase.expected, testCase.hts)
			if testCase.err != nil {
				require.Error(t, err)
				assert.Equal(t, testCase.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_fetchAPI(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		responseStatus int
		responseBody   io.ReadCloser
		data           []serverData
		err            error
	}{
		"http response status not ok": {
			responseStatus: http.StatusNoContent,
			err:            errors.New("HTTP status code not OK: 204 No Content"),
		},
		"nil body": {
			responseStatus: http.StatusOK,
			err:            errors.New("failed unmarshaling response body: EOF"),
		},
		"no server": {
			responseStatus: http.StatusOK,
			responseBody:   ioutil.NopCloser(strings.NewReader(`[]`)),
			data:           []serverData{},
		},
		"success": {
			responseStatus: http.StatusOK,
			responseBody: ioutil.NopCloser(strings.NewReader(`[
				{"connectionName":"host1","region":"region1","country":"country1","location":"location1"},
				{"connectionName":"host2","region":"region2","country":"country1","location":"location2","pubKey":"some key"}
			]`)),
			data: []serverData{
				{
					Region:   "region1",
					Country:  "country1",
					Location: "location1",
					Host:     "host1",
				},
				{
					Region:   "region2",
					Country:  "country1",
					Location: "location2",
					Host:     "host2",
					PubKey:   "some key",
				},
			},
		},
	}
	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			client := &http.Client{
				Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
					assert.Equal(t, http.MethodGet, r.Method)
					assert.Equal(t, r.URL.String(), "https://my.surfshark.com/vpn/api/v4/server/clusters")
					return &http.Response{
						StatusCode: testCase.responseStatus,
						Status:     http.StatusText(testCase.responseStatus),
						Body:       testCase.responseBody,
					}, nil
				}),
			}

			data, err := fetchAPI(ctx, client)

			assert.Equal(t, testCase.data, data)
			if testCase.err != nil {
				require.Error(t, err)
				assert.Equal(t, testCase.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
