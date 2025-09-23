package telegram

type ChatStruct struct {
	ID   int64
	Lang string
}

var Chat ChatStruct
var DeleteMessageID int // id сообщения которое будет удалено