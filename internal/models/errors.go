package models

import "errors"

var (
	ErrNoRecord           = errors.New("models: подходящей записи не найдено")
	ErrInvalidCredentials = errors.New("models: неверные учетные данные")
	ErrDuplicateEmail     = errors.New("models: дублирование электронной почты")
)
