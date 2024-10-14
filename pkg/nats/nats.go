package nats

import (
	"github.com/nats-io/nats.go"
)

const (
	StreamName     = "REVIEWS"
	StreamSubjects = "REVIEWS.*"

	SubjectNameReviewCreated  = "REVIEWS.rateGiven"
	SubjectNameReviewAnswered = "REVIEWS.rateAnswer"
)

func JetStreamInit() (nats.JetStreamContext, error) {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	// Create JetStream Context
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, err
	}

	return js, nil
}

// Return a nats connection
// func JetStreamInit(cfg *config.Config) (nats.JetStreamContext, error) {

// 	// Connect to NATS
// 	nc, err := nats.Connect(cfg.Nats.URL)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create JetStream Context
// 	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create a stream if it does not exist
// 	err = CreateStream(js)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return js, nil

// }

// func CreateStream(jetStream nats.JetStreamContext) error {
// 	stream, err := jetStream.StreamInfo(StreamName)

// 	// stream not found, create it
// 	if stream == nil {
// 		log.Printf("Creating stream: %s\n", StreamName)

// 		_, err = jetStream.AddStream(&nats.StreamConfig{
// 			Name:     StreamName,
// 			Subjects: []string{StreamSubjects},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
