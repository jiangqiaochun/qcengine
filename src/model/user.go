package model

type User struct {
	Id   string `bson:"_id,omitempty" json:"id,omitempty" form:"id,omitempty" sql:"id"`
	Name string `bson:"name,omitempty" json:"name,omitempty" form:"name,omitempty" sql:"name"`
	Age  int    `bson:"age,omitempty" json:"age,omitempty" form:"age,omitempty" sql:"age"`
}
