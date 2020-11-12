package model

type User struct {
	Id   string `bson:"_id,omitempty" json:"id,omitempty" form:"id,omitempty"`
	Name string `bson:"name,omitempty" json:"name,omitempty" form:"id,omitempty"`
}
