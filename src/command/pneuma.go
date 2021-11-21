package command

import (
    tb "gopkg.in/tucnak/telebot.v2"
    "wlupusbot/src/error"
    "wlupusbot/src/logger"
    "wlupusbot/src/model"
    "wlupusbot/src/sql"
)

var (
    successGetPneuma = map[string]string{
        languageFlags[EN]: `The number of your pneuma is: `,
        languageFlags[CN]: `你的pneuma余量为: `,
    }
)

func getPneuma(userName string) (int, int) {
    db := sql.DB
    res := createIfNotExist(userName)
    if res == error.InternalError {
        return error.InternalError, 0
    }
    if res == error.NoUserName {
        return error.NoUserName, 0
    }

    var user model.User
    db.Get(&user, "select * from users where username=?", userName)
    logger.Info("quest pneuma number:", userName)
    return error.Success, user.Pneuma
}

func RegisterGetPneuma(bot *tb.Bot) {
    bot.Handle("/pneuma", func(m *tb.Message) {
        code, value := getPneuma(m.Sender.Username)
        if code == error.InternalError {
            response, ok := internalError[defaultLanguage]
            if !ok {
                response = internalError[languageFlags[CN]]
            }
            bot.Reply(m, response)
        } else if code == error.NoUserName {
            response, ok := noUserName[defaultLanguage]
            if !ok {
                response = noUserName[languageFlags[CN]]
            }
            bot.Reply(m, response)
        } else if code == error.Success {
            response, ok := successGetPneuma[defaultLanguage]
            if !ok {
                response = successGetPneuma[languageFlags[CN]]
            }
            response += string(value)
            bot.Reply(m, response)
        }
    })
}