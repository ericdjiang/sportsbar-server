package stream

type EventDataType int64

const (
	NewChat EventDataType = iota
	GameUpdate
	GameUpdates
)

type Marshal interface {
	ToJSON()
}

type EventData struct {
	Type EventDataType
	Data Marshal
}
