package models

import "time"

type Comment struct {
	Id           int64     `json:"id" xorm:"not null pk autoincr"`

	ProjectId    int64     `json:"projectId" xorm:"not null index"`
	UserId       int64     `json:"userId" xorm:"not null index"`
	IsTeacher    bool      `json:"isTeacher" xorm:"not null index default false"`

	Content      string    `json:"content" xorm:"text"`

	CreateAt     time.Time `json:"createAt" xorm:"created"`

}

type CommentReply struct {
	Id           int64     `json:"id" xorm:"not null pk autoincr"`

	CommentId    int64     `json:"commentId" xorm:"not null index"`
	UserId       int64     `json:"userId" xorm:"not null index"`
	IsTeacher    bool      `json:"isTeacher" xorm:"not null index default false"`

	Content      string    `json:"content" xorm:"text"`

	CreateAt     time.Time `json:"createAt" xorm:"created"`
}

