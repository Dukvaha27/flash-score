package dto

type MatchServiceMatchResponse struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`

	HomeTeamID uint `json:"home_team_id"`
	AwayTeamID uint `json:"away_team_id"`

	HomeScore int `json:"home_score"`
	AwayScore int `json:"away_score"`
}

type MatchEventCreatedMessage struct {
	MatchID     uint   `json:"match_id"`
	EventType   string `json:"event_type"`
	Minute      int    `json:"minute"`
	Description string `json:"description"`

	TeamID   *uint `json:"team_id"`
	PlayerID *uint `json:"player_id"`
}

type MatchGoalMessage struct {
	MatchID uint `json:"match_id"`

	TeamID   uint  `json:"team_id"`
	PlayerID *uint `json:"player_id"`

	Minute   int `json:"minute"`
	NewScore int `json:"new_score"`
}
