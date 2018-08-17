package handlers

type EventDao interface {
	AddEvent(eventType string) bool
	CountByType(eventType string) int
}