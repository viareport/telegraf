package geteventstore

import (
	"testing"
	"github.com/telegraf/testutil"
	"io/ioutil"
	"strings"
	"net/http"
)

type responseMock struct {
	body       string
	accept     string
	statusCode int
}

type transportMock struct {
	responses map[string]responseMock
}

func newTransportMock(transport transportMock) http.RoundTripper {
	return &transport
}

func (t *transportMock) addResponse(path string, response responseMock) {
	t.responses[path] = response
}

func (t *transportMock) RoundTrip(r *http.Request) (*http.Response, error) {
	resp := t.responses[r.URL.Path]
	res := &http.Response{
		Header:     make(http.Header),
		Request:    r,
		StatusCode: resp.statusCode,
	}
	res.Header.Set("Content-Type", resp.accept)
	res.Body = ioutil.NopCloser(strings.NewReader(resp.body))
	return res, nil
}

func (t *transportMock) CancelRequest(_ *http.Request) {
}

func TestGetEventStore(t *testing.T) {
	tes := newGetEventStoreWithReader()

	var acc testutil.Accumulator
	if err := tes.Gather(&acc); err != nil {
		t.Fatal(err)
	}

	tags := map[string]string{
		"queueName":            "MainQueue",
		"lastProcessedMessage": "Schedule",
	}

	acc.AssertContainsTaggedFields(t, "es-queue", esQueueExpected, tags)
}

func newGetEventStoreWithReader() *TelegrafGetEventStore {
	es := &TelegrafGetEventStore{
		Url:        "http://192.168.1.30:2113",
		Login:      "admin",
		Password:   "changeit",
		StreamName: "$stats-0.0.0.0:2113",
	}
	resp1 := responseMock{
		body:       lastEvents,
		accept:     "application/atom+xml",
		statusCode: http.StatusOK,
	}
	resp2 := responseMock{
		body:       statsResponse,
		accept:     "application/vnd.eventstore.atom+json",
		statusCode: http.StatusOK,
	}
	tt := transportMock{make(map[string]responseMock)}
	tt.addResponse("/streams/$stats-0.0.0.0:2113/head/backward/1", resp1)
	tt.addResponse("/streams/$stats-0.0.0.0:2113/16570", resp2)
	transport := newTransportMock(tt)
	client := &http.Client{Transport: transport}
	es.createClient(client)
	return es
}
