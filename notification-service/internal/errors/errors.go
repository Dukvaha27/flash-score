package errors

import "errors"

var ErrNotificationNotFound = errors.New("Уведомление не найдена")
var ErrSubscriptionNotFound = errors.New("Подписка не найдена")
var ErrSubscriptionTargetRequired = errors.New("Необходимо указать объект подписки")
var ErrTeamOrSportRequired = errors.New("необходимо выбрать один объект подписки")
var ErrSubscriptionAlreadyExists = errors.New("Подписка уже существует")
