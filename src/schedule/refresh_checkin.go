package schedule

import (
    "github.com/robfig/cron/v3"
    "wlupusbot/src/logger"
    "wlupusbot/src/sql"
)

const (
    createDailyCheckin string = `create table daily_checkin(
    idx int unsigned primary key,
    username varchar(50) unique not null,
    checkin_status int unsigned default 0,
    greeting_status int unsigned default 0
);`
)
func RefreshCheckin() {
    db := sql.DB
    c := cron.New(cron.WithSeconds())
    expression := "0 0 21 * * *"
    c.AddFunc(expression, func() {
        _, err := db.Exec("drop table daily_checkin")
        if err != nil {
            logger.Error("Failed to drop daily_checkin:", err)
        } else {
            logger.Info("drop table daily_checkin")
        }
        _, err = db.Exec(createDailyCheckin)
        if err != nil {
            logger.Error("Failed to create daily_checkin:", err)
        } else {
            logger.Info("create daily_checkin")
        }
    })
    c.Start()
}