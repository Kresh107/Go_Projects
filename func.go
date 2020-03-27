package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yanzay/tbot/v2"
)

var picks = []string{"rock", "paper", "scissors"}

func init(){
	rand.Seed(time.Now().Unix()) // Initialize global pseudo random generator
}

func makeButtons() *tbot.InlineKeyboardMarkup {
	//create buttons with test and callbackData as value
	btnRock := tbot.InlineKeyboardButton{
		Text:			"Rock",
		CallbackData:	"rock",
	}
	btnPaper :=	tbot.InlineKeyboardButton{
		Text:			"Paper",
		CallbackData:	"paper",
	}
	btnScissors := tbot.InlineKeyboardButton{
		Text:			"Scissors",
		CallbackData:	"scissorss",
	}
	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			[]tbot.InlineKeyboardButton{btnRock, btnPaper, btnScissors},
		},
	}
}

func draw(humanMove string) (msg string) {
	var result string
	botMove := picks[rand.Intn(len(picks))] // Gen a random option for bot

	// Determine the outcome
	switch humanMove {
	case botMove:
		result = "drew"
	case options[botMove]:
		result = "lost"
	default:
		result = "won"
	}
	msg = fmt.Sprintf("You %s! You chose %s and I chose %s.", result, humanMove, botMove)
	return
}