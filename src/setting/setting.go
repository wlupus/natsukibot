package setting

import (
    "github.com/go-ini/ini"
    "wlupusbot/src/logger"
)

var (
    TelebotToken string

    MysqlUserName string
    MysqlPassword string

    CheckinReward int
    cfg *ini.File
)

func init() {
    var err error
    cfg, err = ini.Load("src/conf/app.ini")
    if err != nil {
        logger.Fatal("Fail to parse config file:", err)
    }
    loadToken()
    loadMysql()
    loadBot()
}

func loadMysql() {
    sec, err := cfg.GetSection("mysql")
    if err != nil {
        logger.Fatal("Fail to get section 'mysql':", err)
    }
    MysqlUserName = sec.Key("username").String()
    MysqlPassword = sec.Key("password").String()
}

func loadToken() {
    sec, err := cfg.GetSection("telegram")
    if err != nil {
        logger.Fatal("Fail to get section 'telegram':", err)
    }
    TelebotToken = sec.Key("bot_token").String()
}

func loadBot() {
    sec, err := cfg.GetSection("bot")
    if err != nil {
        logger.Error("Fail to get section 'bot':", err)
    } else {
        CheckinReward = sec.Key("checkin_reward").MustInt(0)
    }
}