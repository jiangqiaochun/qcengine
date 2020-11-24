package mysql

import (
	"log"
	"qcengine/src/model"
	"testing"
)

func TestSqlHeader_GenerateInsertSql(t *testing.T) {
	user := new(model.User)
	user.Name = "jiang"
	user.Age = 12
	sql := NewSqlSession()
	insetSql := sql.GenerateInsertSql(user)
	log.Println(insetSql)
}
