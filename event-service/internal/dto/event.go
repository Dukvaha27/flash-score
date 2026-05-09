package dto

type CreateEventRequest struct {
	MatchID   uint   `json:"match_id" binding:"required"`
	EventType string `json:"event_type" binding:"required"`
	Minute    int    `json:"minute" binding:"required,min=0"`

	TeamID   *uint `json:"team_id"`
	PlayerID *uint `json:"player_id"`

	Text string `json:"text"`
}

type EventResponse struct {
	ID        uint   `json:"id"`
	MatchID   uint   `json:"match_id"`
	EventType string `json:"event_type"`
	Minute    int    `json:"minute"`

	TeamID   *uint `json:"team_id"`
	PlayerID *uint `json:"player_id"`

	Text      string `json:"text"`
	CreatedBy uint   `json:"created_by"`
	CreatedAt string `json:"created_at"`
}
