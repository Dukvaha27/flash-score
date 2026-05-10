package services

import "errors"

var (
	ErrForbidden                = errors.New("forbidden")
	ErrInvalidEventType         = errors.New("invalid event type")
	ErrInvalidReaction          = errors.New("invalid reaction type")
	ErrInvalidTarget            = errors.New("target must be either event or commentary")
	ErrMatchNotLive             = errors.New("match is not live")
	ErrTeamRequired             = errors.New("team_id is required for goal event")
	ErrTeamNotInMatch           = errors.New("team does not participate in this match")
	ErrMatchClientNotConfigured = errors.New("match client is not configured")
)
