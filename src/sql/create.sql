create table users(
    uid int unsigned primary key,
    username varchar(50) unique not null,
    pneuma int unsigned default 0 not null
);

create table daily_checkin(
    idx int unsigned primary key,
    username varchar(50) unique not null,
    checkin_status int unsigned default 0,
    greeting_status int unsigned default 0
);
