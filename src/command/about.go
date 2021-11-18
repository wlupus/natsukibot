package command

import (
    tb "gopkg.in/tucnak/telebot.v2"
    "wlupusbot/src/logger"
)

const (
    aboutTextEN string = `Hi, I'm AkiyamaNatsuki. I come from Dokoka, a small planet in the cosmos.
I love listening to classical music, reading novels and, of cause, you.
If you find any bug or have any suggestion, please spare your precious time to contact me.
Feel free to request new features.

Email: akiyamanatsuki@outlook.com

AkiyamaNatsuki 14/11/2021
version: Internal Test alpha`
    aboutTextCN string = `嗨, 我是秋山夏纪, Dokoka星球人.
喜欢古典乐, 小说, 当然还有你٩(๑❛ᴗ❛๑)۶
漏洞反馈 / 建议 / 提出新功能: akiyamanatsuki@outlook.com

秋山夏纪 2021/11/14
version: Internal Test alpha`
    aboutTextJP string = `こんにちは。私は秋山夏紀です。`
)

var (
    aboutText = map[string]string{
        languageFlags[EN]: aboutTextEN,
        languageFlags[CN]: aboutTextCN,
        languageFlags[JP]: aboutTextJP,
    }
)

func RegisterAbout(bot *tb.Bot) {
    bot.Handle("/about", func(m *tb.Message) {
        logger.Info("Call /about:", m.Sender.Username)
        response, ok := aboutText[defaultLanguage]
        if !ok {
            response = aboutText[languageFlags[CN]]
        }
        bot.Reply(m, response)
    })
}
