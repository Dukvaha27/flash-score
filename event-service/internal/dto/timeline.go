package dto

type TimelineItemResponse struct {
	Type string `json:"type"`

	ID      uint `json:"id"`
	MatchID uint `json:"match_id"`
	Minute  int  `json:"minute"`

	EventType *string `json:"event_type,omitempty"`

	TeamID   *uint `json:"team_id,omitempty"`
	PlayerID *uint `json:"player_id,omitempty"`

	Text      string `json:"text"`
	IsPinned  *bool  `json:"is_pinned,omitempty"`
	CreatedBy uint   `json:"created_by"`
	CreatedAt string `json:"created_at"`
}
