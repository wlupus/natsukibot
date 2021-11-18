package model

type User struct {
    Uid int `db:"uid"`
    UserName string `db:"username"`
    Pneuma int `db:"pneuma"`
}

type DailyCheckin struct {
    Idx int `db:"idx"`
    UserName string `db:"username"`
    CheckinStatus bool `db:"checkin_status"`
    GreetingStatus bool `db:"greeting_status"`
}
