package mysql

import (
	"database/sql"
	"errors"
	"log"
	"qcengine/src/common/database"
	"reflect"
	"strconv"
	"strings"
)

type SqlHeader struct {
	database.DataBase
	config *database.DatabaseConfig
	database *sql.DB
}

type SqlSession struct {
	*SqlHeader
}
func NewSqlSession() *SqlSession {
	instance := new(SqlSession)
	return instance
}
func (sh *SqlHeader) Init(config *database.DatabaseConfig) *SqlHeader {
	sh.config = config
	return sh
}

func (sh *SqlHeader) Connect() error {
	config := sh.config
	url := ""
	if config.UserName != "" {
		url+= config.UserName + ":" + config.Password + "@tcp("
	}
	url += config.HostName
	if config.HostPort != 0 {
		url+= ":" +strconv.Itoa(config.HostPort) + ")"
	}
	if config.DataBaseName != "" {
		url+= "/" + config.DataBaseName + "?charset=utf8"
	}
	db, error := sql.Open("mysql", url)
	if error != nil {
		return error
	}
	sh.database = db
	if sh.database == nil {
		return errors.New("连接到"+config.DataBaseName+"失败")
	}
	return nil
}

func (sh *SqlHeader) GenerateInsertSql(object interface{}) string {
	insertSql := "INSERT INTO "
	targetType := reflect.TypeOf(object).Elem()
	modelName := targetType.Name()
	tableName := strings.ToLower(modelName)
	insertSql += "`" + tableName + "`("
	for i:=0;i<targetType.NumField();i++ {
		jsonName := targetType.Field(i).Tag.Get("sql")
		if i==targetType.NumField()-1 {
			insertSql += "`"+jsonName+"`)"
		} else {
			insertSql += "`"+jsonName+"`,"
		}
	}
	insertSql += " VALUES("
	targetValue := reflect.ValueOf(object).Elem()
	for i:=0;i<targetValue.NumField();i++ {
		f := targetValue.Field(i)
		fieldType := f.Type().Name()
		if fieldType == "string" {
			insertSql += "`"+f.String()+"`"
		} else if fieldType == "int" {
			insertSql += strconv.Itoa(int(f.Int()))
		}
		if i < targetValue.NumField()-1 {
			insertSql += ","
		}
	}
	insertSql += ")"
	return insertSql
}

func (ss *SqlSession) Insert(object interface{}) (interface{}, error) {
	insertSql := ss.GenerateInsertSql(object)
	log.Println(insertSql)
	return nil, nil
}
