package command

import (
    tb "gopkg.in/tucnak/telebot.v2"
    "wlupusbot/src/error"
    "wlupusbot/src/logger"
    "wlupusbot/src/model"
    "wlupusbot/src/sql"
)

var (
	successCheckin = map[string]string{
        languageFlags[EN]: `Checkin success`,
        languageFlags[CN]: `签到成功`,
    }
    alreadyCheckin = map[string]string{
        languageFlags[EN]: `You've checked in already`,
        languageFlags[CN]: `你已经签到过了`,
    }
    internalError = map[string]string{
        languageFlags[EN]: `Internal server error`,
        languageFlags[CN]: `内部服务错误`,
    }
)

func checkin(userName string) int {
    db := sql.DB
    res := createIfNotExist(userName)
    if res == error.InternalError {
        return error.InternalError
    }
    // check existence of table daily_checkin
    var index int
    err := db.QueryRow("select count(*) from daily_checkin").Scan(&index)
    if err != nil {
        logger.Error("Failed to access daily_checkin table:", err)
        return error.InternalError
    }

    // get user
    var status []model.DailyCheckin
    err = db.Select(&status, "select * from daily_checkin where username=?", userName)
    if len(status) == 1 && status[0].CheckinStatus {
        logger.Info("Already Checkin:", userName)
        return error.AlreadyCheckin
    } else {
        db.Exec("insert into daily_checkin (idx, username, checkin_status, greeting_status)"+
            " values (?, ?, ?, ?) on duplicate key update checkin_status=1",
            index+1, userName, 1, 0)
        if len(status) == 0 {
            logger.Info("Create daily_checkin row for", userName)
        } else {
            logger.Info("update daily_checkin row for ", userName)
        }
    }
    return error.Success
}

func RegisterCheckin(bot *tb.Bot) {
    bot.Handle("/checkin", func(m *tb.Message) {
        code := checkin(m.Sender.Username)
        if code == error.Success {
            response, ok := successCheckin[defaultLanguage]
            if !ok {
                response = successCheckin[languageFlags[CN]]
            }
            bot.Reply(m, response)
        } else if code == error.AlreadyCheckin {
            response, ok := alreadyCheckin[defaultLanguage]
            if !ok {
                response = alreadyCheckin[languageFlags[CN]]
            }
            bot.Reply(m, response)
        } else if code == error.InternalError {
            response, ok := internalError[defaultLanguage]
            if !ok {
                response = internalError[languageFlags[CN]]
            }
            bot.Reply(m, response)
        }
    })
}