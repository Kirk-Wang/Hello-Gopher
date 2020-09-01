package models

type LtUserday struct {
	Id         int `xorm:"not null pk autoincr INT(10)"`
	Uid        int `xorm:"not null default 0 comment('用户ID') unique(uid_day) INT(10)"`
	Day        int `xorm:"not null default 0 comment('日期，如：20180725') unique(uid_day) INT(10)"`
	Num        int `xorm:"not null default 0 comment('次数') INT(10)"`
	SysCreated int `xorm:"not null default 0 comment('创建时间') INT(10)"`
	SysUpdated int `xorm:"not null default 0 comment('修改时间') INT(10)"`
}
