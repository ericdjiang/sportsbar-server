package stream

type Chat struct {
	GameID   int
	Msg      string
	UserName string
}

type MessageDataType int64

const (
	NewChat MessageDataType = iota
	GameUpdate
	GameUpdates
)

type SocketMessage struct {
	Type MessageDataType
	Data any
}
