package chat

type Message struct {
    Message string `json:"message"`
    Nickname string `json:"nickname"`
}

func NewMessage(nickname, message string) *Message {
    return &Message{
        message,
        nickname,
    }
}
