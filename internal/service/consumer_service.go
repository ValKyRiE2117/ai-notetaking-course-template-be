package service

import (
	"ai-notetaking-be/internal/dto"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/gofiber/fiber/v2/log"
)

type IConsumerService interface {
	Consume(ctx context.Context) error
}

type consumerService struct {
	pubSub *gochannel.GoChannel	
	topicName string
}

func (cs *consumerService) Consume(ctx context.Context) error {
	messages, err := cs.pubSub.Subscribe(ctx, cs.topicName)
	if err != nil {
		return err
	}

	go func() {
	for msg := range messages {
		err := cs.processMessage(ctx, msg)
		if err != nil {
			log.Error(err)
		}
	}
	}()

	return nil 
}

func (cs *consumerService) processMessage(ctx context.Context, msg *message.Message) error {
	defer msg.Nack()
	defer func() {
		if e := recover(); e != nil {
			log.Error(e)
		}
	}()
	var payload dto.PublishEmbedNoteMessage
	err := json.Unmarshal(msg.Payload, &payload)
	if err != nil {
		return err
	}

	http.NewRequest(
		"POST",
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-3-flash-preview:generateContent",
		
		
	)

	client := &http.Client{}
	client.Do()

	fmt.Println(payload.NoteId)
	msg.Ack()
	return nil
}

func NewConsumerService(pubSub *gochannel.GoChannel, topicName string) IConsumerService {
	return &consumerService{
		pubSub: pubSub,
		topicName: topicName,
	}
}