package services

import (
	"context"

	"github.com/Dukvaha27/flash-score/event-service/internal/dto"
)

type MatchClient interface {
	GetMatchByID(ctx context.Context, matchID uint) (*dto.MatchServiceMatchResponse, error)
}
