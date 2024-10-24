package nats

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type Publisher struct {
	js jetstream.JetStream
}

func NewPublisher(url string) (*Publisher, error) {
	// Conectar a NATS con timeout
	nc, err := nats.Connect(url,
		nats.Timeout(30*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	// Crear interfaz JetStream
	js, err := jetstream.New(nc)
	if err != nil {
		nc.Close()
		return nil, fmt.Errorf("failed to create JetStream management interface: %w", err)
	}

	return &Publisher{
		js: js,
	}, nil
}

// PublishReview publica una review en el stream
// func (p *Publisher) PublishReview(ctx context.Context, review Review) error {
// 	data, err := json.Marshal(review)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal review: %w", err)
// 	}

// 	ack, err := p.js.Publish(ctx, SubjectNameReviewCreated, data)
// 	if err != nil {
// 		return fmt.Errorf("failed to publish review: %w", err)
// 	}

// 	_ = ack
// 	return nil
// }

// // PublishReviewAnswer publica una respuesta a una review
// func (p *Publisher) PublishReviewAnswer(ctx context.Context, reviewID string, answer string) error {
// 	response := struct {
// 		ReviewID  string    `json:"review_id"`
// 		Answer    string    `json:"answer"`
// 		AnswerAt time.Time `json:"answer_at"`
// 	}{
// 		ReviewID:  reviewID,
// 		Answer:    answer,
// 		AnswerAt: time.Now(),
// 	}

// 	data, err := json.Marshal(response)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal review answer: %w", err)
// 	}

// 	ack, err := p.js.Publish(ctx, SubjectNameReviewAnswered, data)
// 	if err != nil {
// 		return fmt.Errorf("failed to publish review answer: %w", err)
// 	}

// 	_ = ack
// 	return nil
// }
