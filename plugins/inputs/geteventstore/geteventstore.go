package geteventstore

import (
	"github.com/jetbasrawi/go.geteventstore"
	"strings"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type SReader interface {
	Next() bool
	NextVersion(i int)
	Err() error
	Scan(interface{}, interface{})
}

type TelegrafGetEventStore struct {
	Url        string
	Login      string
	Password   string
	StreamName string
	Client     *goes.Client
}

const sampleConfig = `
	## Specify eventstore url
	url="http://127.0.0.1:2113"

	## specify login
	login="admin"
	## specify password
	password="changeit"
	## specify stats stream name
	streamname="$stats-0.0.0.0:2113"
`

func (*TelegrafGetEventStore) SampleConfig() string {
	return sampleConfig
}

func (*TelegrafGetEventStore) Description() string {
	return "Read stats from one eventstore"
}

func (store *TelegrafGetEventStore) createClient(httpClient *http.Client) {
	client, _ := goes.NewClient(httpClient, store.Url)
	client.SetBasicAuth(store.Login, store.Password)
	store.Client= client
}

func (store *TelegrafGetEventStore) getLastEvent() (*goes.EventResponse, error){
	path, _ := store.Client.GetFeedPath(store.StreamName, "backward", -1, 1)
	f, _, err := store.Client.ReadFeed(path)
	if err != nil {
		return nil, err
	}
	numEntries := len(f.Entry)
	index := numEntries - 1
	entry := f.Entry[index]
	url := strings.TrimRight(entry.Link[1].Href, "/")
	u, _, err := store.Client.GetEvent(url)
	return u, err
}

func (store *TelegrafGetEventStore) Scan(eventResponse *goes.EventResponse, e interface{}, m interface{}) error {

	if e != nil {
		data, ok := eventResponse.Event.Data.(*json.RawMessage)
		if !ok {
			return fmt.Errorf("Could not unmarshal the event. Event data is not of type *json.RawMessage")
		}

		if err := json.Unmarshal(*data, e); err != nil {
			return err
		}
	}

	if m != nil && eventResponse.Event.MetaData != nil {
		meta, ok := eventResponse.Event.MetaData.(*json.RawMessage)
		if !ok {
			return fmt.Errorf("Could not unmarshal the event. Event meta data is not of type *json.RawMessage")
		}

		if err := json.Unmarshal(*meta, &m); err != nil {
			return err
		}
	}

	return nil
}

func (store *TelegrafGetEventStore) Gather(acc telegraf.Accumulator) error {
	if store.Client == nil {
		store.createClient(nil)
	}
	fields := make(map[string]interface{})
	tags := make(map[string]string)
	event, err := store.getLastEvent()

	if err != nil {
		return err
	}

	// If the call did not result in an error then an event was returned
	// Create the application types to hold the deserialized even data and meta data
	statsEvent := make(map[string]interface{})

	// Call scan to deserialize the event data and meta data into your types
	err = store.Scan(event, &statsEvent, nil)

	if err != nil {
		// Handle errors that occured during deserialization
		return err
	}

	for key, value := range statsEvent {
		if strings.Contains(key, "es-queue-MainQueue") {
			field := strings.Split(strings.Split(key, "es-queue-")[1], "-")[1]
			if strings.EqualFold(field, "queueName") {
				tags["queueName"] = value.(string)
			} else if strings.EqualFold(field, "lastProcessedMessage") {
				tags["lastProcessedMessage"] = value.(string)
			} else {
				fields[field] = value
			}
		}
	}

	acc.AddFields("es-queue", fields, tags)
	return nil
}

func init() {
	inputs.Add("geteventstore", func() telegraf.Input {
		return &TelegrafGetEventStore{}
	})
}