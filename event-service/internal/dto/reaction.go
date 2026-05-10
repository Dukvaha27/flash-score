package dto

type UpsertReactionRequest struct {
	EventID      *uint  `json:"event_id"`
	CommentaryID *uint  `json:"commentary_id"`
	Type         string `json:"type" binding:"required"`
}
