package dto

type CreateCommentaryRequest struct {
	MatchID uint   `json:"match_id" binding:"required"`
	Minute  int    `json:"minute" binding:"required,min=0"`
	Text    string `json:"text" binding:"required"`

	IsPinned bool `json:"is_pinned"`
}

type UpdateCommentaryRequest struct {
	Minute int    `json:"minute" binding:"required,min=0"`
	Text   string `json:"text" binding:"required"`
}

type CommentaryResponse struct {
	ID        uint   `json:"id"`
	MatchID   uint   `json:"match_id"`
	Minute    int    `json:"minute"`
	Text      string `json:"text"`
	IsPinned  bool   `json:"is_pinned"`
	CreatedBy uint   `json:"created_by"`
	CreatedAt string `json:"created_at"`
}
