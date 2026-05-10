package services

import (
	"sort"
	"time"

	"github.com/Dukvaha27/flash-score/event-service/internal/dto"
	"github.com/Dukvaha27/flash-score/event-service/internal/repositories"
)

type TimelineService interface {
	GetByMatchID(matchID uint) ([]dto.TimelineItemResponse, error)
}

type timelineService struct {
	eventRepo      repositories.MatchEventRepository
	commentaryRepo repositories.CommentaryRepository
}

func NewTimelineService(
	eventRepo repositories.MatchEventRepository,
	commentaryRepo repositories.CommentaryRepository,
) TimelineService {
	return &timelineService{
		eventRepo:      eventRepo,
		commentaryRepo: commentaryRepo,
	}
}

func (s *timelineService) GetByMatchID(matchID uint) ([]dto.TimelineItemResponse, error) {
	events, err := s.eventRepo.ListByMatchID(matchID)
	if err != nil {
		return nil, err
	}

	commentaries, err := s.commentaryRepo.ListByMatchID(matchID)
	if err != nil {
		return nil, err
	}

	items := make([]dto.TimelineItemResponse, 0, len(events)+len(commentaries))

	for _, event := range events {
		eventType := string(event.EventType)

		items = append(items, dto.TimelineItemResponse{
			Type:      "event",
			ID:        event.ID,
			MatchID:   event.MatchID,
			Minute:    event.Minute,
			EventType: &eventType,
			TeamID:    event.TeamID,
			PlayerID:  event.PlayerID,
			Text:      event.Text,
			CreatedBy: event.CreatedBy,
			CreatedAt: event.CreatedAt.Format(time.RFC3339),
		})
	}

	for _, commentary := range commentaries {
		isPinned := commentary.IsPinned

		items = append(items, dto.TimelineItemResponse{
			Type:      "commentary",
			ID:        commentary.ID,
			MatchID:   commentary.MatchID,
			Minute:    commentary.Minute,
			Text:      commentary.Text,
			IsPinned:  &isPinned,
			CreatedBy: commentary.CreatedBy,
			CreatedAt: commentary.CreatedAt.Format(time.RFC3339),
		})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].Minute != items[j].Minute {
			return items[i].Minute < items[j].Minute
		}

		return items[i].CreatedAt < items[j].CreatedAt
	})

	return items, nil
}
