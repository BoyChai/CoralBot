package action

type TelegramHandle struct{}

func (t TelegramHandle) GetHandlerType() string {
	return "Telegram"
}
