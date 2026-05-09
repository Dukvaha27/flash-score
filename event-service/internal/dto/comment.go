package dto

type CreateCommentRequest struct {
	EventID      *uint `json:"event_id"`
	CommentaryID *uint `json:"commentary_id"`

	Text string `json:"text" binding:"required"`
}

type CommentResponse struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`

	EventID      *uint `json:"event_id"`
	CommentaryID *uint `json:"commentary_id"`

	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}
