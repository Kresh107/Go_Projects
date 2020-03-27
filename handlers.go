package main

import (
	"github.com/yanzay/tbot/v2"
)

// Handle the /start command
func (a *application) startHandler(m *tbot.Message) {
	msg := "This bot plays rock paper scissors. Use /play to play."
	a.client.SendMessage(m.Chat.ID, msg)
}

// Handle the /play command
func (a *application) playHandler(m *tbot.Message) {
	buttons := makeButtons()
	a.client.SendMessage(m.Chat.ID, "Pick an Option", tbot.OptInlineKeyboardMarkup(buttons))
}

// Handle button presses here
func (a *application) callbackHandler(cq *tbot.CallbackQuery) {
	humanMove := cq.Data
	msg := draw(humanMove)
	a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	a.client.SendMessage(cq.Message.Chat.ID, msg)
}