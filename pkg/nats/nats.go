package nats

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

const (
	StreamName     = "REVIEWS"
	StreamSubjects = "REVIEWS.*"

	SubjectNameReviewCreated  = "REVIEWS.rateGiven"
	SubjectNameReviewAnswered = "REVIEWS.rateAnswer"
)

func JetStreamInit() (*nats.Conn, jetstream.JetStream, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	// Create a JetStream management interface
	js, _ := jetstream.New(nc)
	if err != nil {
		nc.Close()
		return nil, nil, fmt.Errorf("failed to create JetStream management interface: %w", err)
	}

	// Create a stream
	_, err = js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     StreamName,
		Subjects: []string{StreamSubjects},
	})
	if err != nil {
		nc.Close()
		return nil, nil, fmt.Errorf("failed to create stream: %w", err)
	}

	return nc, js, nil

	// Connect to NATS
	// nc, err := nats.Connect(nats.DefaultURL)
	// if err != nil {
	// 	return nil, err
	// }

	// Create JetStream Context
	// js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	// if err != nil {
	// 	return nil, err
	// }
	// create a stream (this is an idempotent operation)
	// s, _ := js.CreateStream(ctx, jetstream.StreamConfig{
	// 	Name:     "ORDERS",
	// 	Subjects: []string{"ORDERS.*"},
	// })

	// get stream handle
	// s, _ = js.Stream(ctx, "ORDERS")

	// Create a stream if it does not exist
	// err = CreateStream(js)
	// if err != nil {
	// 	return nil, err
	// }

	// return js, nil
}

func CreateStream(jetStream nats.JetStreamContext) error {
	stream, err := jetStream.StreamInfo(StreamName)

	// stream not found, create it
	if stream == nil {
		log.Printf("Creating stream: %s\n", StreamName)

		_, err = jetStream.AddStream(&nats.StreamConfig{
			Name:     StreamName,
			Subjects: []string{StreamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
