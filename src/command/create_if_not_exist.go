package command

import (
    "wlupusbot/src/error"
    "wlupusbot/src/logger"
    "wlupusbot/src/sql"
)

func createIfNotExist(userName string) int {
    if userName == "" {
        return error.NoUserName
    }
    db := sql.DB
    var index int
    err := db.QueryRow("select count(*) from users").Scan(&index)
    if err != nil {
        logger.Error("Fail to get user number when creating user")
        return error.InternalError
    }
    _, err = db.Exec("insert into users(uid, username, pneuma) values (?, ?, ?)", index + 1, userName, 0)
    if err == nil {
        logger.Info("Create user", index + 1, userName, 0)
    }
    return error.Success
}
