package command

import (
    tb "gopkg.in/tucnak/telebot.v2"
    "wlupusbot/src/logger"
)

const (
    EN int = iota
    CN
    JP
)

var (
    defaultLanguage = "cn"
    languageFlags   = [...]string{"en", "cn", "jp"}
    unsupported = map[string]string{
        languageFlags[EN]: `Unsupported language: `,
        languageFlags[CN]: `不支持该语言: `,
        languageFlags[JP]: `Unsupported language: `,
    }
    switchLanguage = map[string]string{
        languageFlags[EN]: `Switch default language to `,
        languageFlags[CN]: `切换默认语言为 `,
        languageFlags[JP]: `Switch default language to `,
    }
)

func RegisterLanguage(bot *tb.Bot) {
    bot.Handle("/language", func(m *tb.Message){
        check := false
        for _, flag := range languageFlags {
            if flag == m.Payload {
                defaultLanguage = flag
                logger.Info("Switch default language to", flag)
                bot.Reply(m, switchLanguage[flag] + flag)
                check = true
                break
            }
        }
        if !check {
            bot.Reply(m, unsupported[defaultLanguage] + m.Payload)
        }
    })
}
