package services

import "github.com/Dukvaha27/flash-score/event-service/internal/models"

func canManageMatchContent(role string) bool {
	return role == "commentator" || role == "admin"
}

func canDeleteUserContent(ownerID uint, currentUserID uint, role string) bool {
	return ownerID == currentUserID || role == "admin"
}

func hasExactlyOneTarget(eventID *uint, commentaryID *uint) bool {
	return (eventID != nil && commentaryID == nil) ||
		(eventID == nil && commentaryID != nil)
}

func isValidEventType(eventType string) bool {
	switch models.EventType(eventType) {
	case models.Goal,
		models.YellowCard,
		models.RedCard,
		models.Substitution,
		models.Penalty,
		models.VarDecision,
		models.HalfStart,
		models.HalfEnd,
		models.FullTime,
		models.Injury,
		models.Timeout:
		return true
	default:
		return false
	}
}

func isValidReactionType(reactionType string) bool {
	switch models.ReactionsType(reactionType) {
	case models.Like,
		models.Fire,
		models.Shock,
		models.Sad,
		models.Laugh:
		return true
	default:
		return false
	}
}
