package main

import (
    tb "gopkg.in/tucnak/telebot.v2"
    "time"
    "wlupusbot/src/command"
    "wlupusbot/src/logger"
    "wlupusbot/src/schedule"
    "wlupusbot/src/setting"
)

func initializeSchedule() {
    schedule.RefreshCheckin()
}

func registerHandle(bot *tb.Bot) {
    command.RegisterHelp(bot)
    command.RegisterAbout(bot)
    command.RegisterLanguage(bot)
    command.RegisterCheckin(bot)
    command.RegisterGetPneuma(bot)
}

func main() {
    b, err := tb.NewBot(tb.Settings{
        URL: "https://api.telegram.org",
        Token: setting.TelebotToken,
        Poller: &tb.LongPoller{Timeout: 10 * time.Second},
    })
    if err != nil {
        logger.Fatal(err)
        return
    }
    initializeSchedule()
    registerHandle(b)
    b.Start()
}
