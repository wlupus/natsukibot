package sql

import (
    "wlupusbot/src/logger"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "wlupusbot/src/setting"
)

var DB *sqlx.DB

func init() {
    expression := setting.MysqlUserName + ":" + setting.MysqlPassword + "@(127.0.0.1:3306)/wlupusbot"
    db, err := sqlx.Connect("mysql", expression)
    if err != nil {
        logger.Fatal(err)
    }
    DB = db
}
