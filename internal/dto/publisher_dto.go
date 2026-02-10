package dto

import "github.com/google/uuid"

type PublishEmbedMessage struct {
	NoteId uuid.UUID `json:"note_id"`
}