package main

import (
	"bookstore/model"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("My_NSQ_Topic", "My_NSQ_Channel", decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer")
	}
	//c.MaxInFlight defaults to 1

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("NSQ message received:")
		// log.Println(string(message.Body))
		b := model.Book{}
		err := json.Unmarshal(message.Body, &b)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("id:", b.ID)
		fmt.Println("name:", b.Name)
		return nil
	}))

	err = c.ConnectToNSQD("127.0.0.1:32772")
	if err != nil {
		log.Panic("Could not connect")
	}
	log.Println("Awaiting messages from NSQ topic \"My NSQ Topic\"...")
	wg.Wait()
}
