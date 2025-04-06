package domain

import "time"

type CreateScheduleDto struct {
	UserId    int64
	Memo      string
	StartTime time.Time
	EndTime   time.Time
}

type DeleteScheduleDto struct {
	UserId int64
	MemoId int64
}

type RegisterUserDto struct {
}
