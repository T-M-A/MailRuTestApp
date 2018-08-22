package dao_mock

type EventDaoMock struct {
	Store map[string]int
}

func (e *EventDaoMock) AddEvent(eventType string) error {
	if val, exist := e.Store[eventType]; exist {
		e.Store[eventType] = val+1
	} else {
		e.Store[eventType] = 1
	}

	return nil
}

func (e *EventDaoMock) CountByType(eventType string) (int, error) {
	if val, exist := e.Store[eventType]; exist {
		return val, nil
	} else {
		return 0, nil
	}
}

func NewEventDaoMock() *EventDaoMock {
	return &EventDaoMock{
		Store: make(map[string]int),
	}
}