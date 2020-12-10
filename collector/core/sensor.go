package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BGrewell/perspective/helpers"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Poll(host string, port int, interval int, ctx context.Context) (eventChan chan *helpers.SensorEvent, err error) {
	eventChan = make(chan *helpers.SensorEvent, 1000)
	url := fmt.Sprintf("http://%s:%d/events", host, port) //todo: setup ssl
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "perspective-collector")
	go poll(client, req, interval, eventChan, ctx)
	return eventChan, nil
}

func poll(client http.Client, r *http.Request, interval int, eventChan chan *helpers.SensorEvent, ctx context.Context) {
	for {
		select {
		case <- time.After(time.Duration(interval) * time.Second):
			// poll for events
			response, err := client.Do(r)
			if err != nil {
				log.Printf("error: failed to get %v: %v\n", r.URL, err)
				continue
			}
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("error: failed to read body %v: %v\n", r.URL, err)
				continue
			}
			er := helpers.EventsResponse{}
			err = json.Unmarshal(body, &er)
			if err != nil {
				log.Printf("error: failed to unmarshal events %v: %v\n", r.URL, err)
				continue
			}
			for _, event := range er.Events {
				eventChan <- event
			}
		case <- ctx.Done():
			// finish polling
			return
		}
	}
}
