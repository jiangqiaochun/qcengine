package model

type User struct {
	Id   string `bson:"_id" json:"id" form:"id"`
	Name string `bson:"name" json:"name" form:"id"`
}
