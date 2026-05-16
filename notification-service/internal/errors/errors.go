package errors

import "errors"

var ErrInvalidType = errors.New("Неправильный тип данных")
var ErrNotificationNotFound = errors.New("уведомление не найдено")
var ErrSubscriptionNotFound = errors.New("подписка не найдена")
var ErrSubscriptionTargetRequired = errors.New("необходимо указать объект подписки")
var ErrTeamOrSportRequired = errors.New("необходимо выбрать один объект подписки")
var ErrSubscriptionAlreadyExists = errors.New("подписка уже существует")
var ErrUnauthorized = errors.New("не авторизован")
