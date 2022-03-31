package stream

type Chat struct {
	GameID    int
	Msg       string
	UserName  string
	MessageID string
}

type MessageDataType int64

const (
	NewChat MessageDataType = iota
	GameUpdate
	GameUpdates
)

type SocketMessage struct {
	MessageType MessageDataType
	Data        any
}

type NewChatSocketMessage struct {
	MessageType MessageDataType
	Data        Chat
}
