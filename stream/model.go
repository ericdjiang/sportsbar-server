package stream

type Chat struct {
	GameID    int
	Msg       string
	UserName  string
	MessageID string
}

type MessageDataType int64

type Unknown interface{}

const (
	NewChat MessageDataType = iota
	GameUpdate
	GameUpdates
)

type SocketMessage struct {
	MessageType MessageDataType
	Data        Unknown
}

type NewChatSocketMessage struct {
	MessageType MessageDataType
	Data        Chat
}
