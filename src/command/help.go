package command

import (
    tb "gopkg.in/tucnak/telebot.v2"
    "wlupusbot/src/logger"
)

const (
    helpTextEN string = `commands list:
/help show this help message
/about about me
/language change default language e.g. /language en
/pneuma get the value of pneuma
language list: en, cn, jp
`
    helpTextJP string = `コマンド リスト：
/help ヘルプメッセージ
/about 私のこと
/language change default language e.g. /language en
/pneuma get the value of pneuma
language リスト: en, cn, jp
`
    helpTextCN string = `指令列表: 
/help 显示本条帮助信息
/about 关于
/language 更改默认语言 例如: /language cn
/pneuma 查询pneuma余量
语言可选值: en, cn, jp
`
)
var (
    helpText = map[string]string{
        languageFlags[EN]: helpTextEN,
        languageFlags[CN]: helpTextCN,
        languageFlags[JP]: helpTextJP,
    }
)

func RegisterHelp(bot *tb.Bot) {
    bot.Handle("/help", func(m *tb.Message) {
        logger.Info("Call /help:", m.Sender.Username)
        response, ok := helpText[defaultLanguage]
        if !ok {
            response = helpText[languageFlags[CN]]
        }
        bot.Reply(m, response)
    })
}