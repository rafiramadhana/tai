package main

import (
	"bookstore/model"
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:32772", config)
	if err != nil {
		log.Panic(err)
	}
	// err = p.Publish("My_NSQ_Topic", []byte("sample NSQ message"))
	b, err := json.Marshal(model.Book{
		ID:   "1",
		Name: "Filosofi Teras",
	})
	if err != nil {
		log.Panic(err)
	}
	err = p.Publish("My_NSQ_Topic", b)
	if err != nil {
		log.Panic(err)
	}
}
