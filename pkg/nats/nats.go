package nats

import (
	"fmt"
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

	// Conectar a NATS con timeout
	nc, err := nats.Connect(nats.DefaultURL,
		nats.Timeout(30*time.Second),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	// Crear interfaz JetStream usando el contexto
	js, err := jetstream.New(nc)
	if err != nil {
		nc.Close()
		return nil, nil, fmt.Errorf("failed to create JetStream management interface: %w", err)
	}

	return nc, js, nil
}

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
