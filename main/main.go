package main

import tg_done_bot "github.com/leonid-shevtsov/tg-done-bot"

func main() {
	tg_done_bot.RunBot(tg_done_bot.DBConnect())
}
