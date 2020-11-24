package servicepackage

import (
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"qcengine/src/common/database"
	"qcengine/src/common/database/mongodb"
	"strconv"
)

var config = new(Config)

type Config struct {
	XMLName        xml.Name                 `xml:"config"`
	DatabaseConfig *database.DatabaseConfig `xml:"iDatabase"`
}

type Service struct {
	//iDatabase *iDatabase.DatabaseConfig
	iDatabase database.IDataBase
}

func  init() {
	var appPackage = flag.String("appPackage", "resource/config.xml", "配置文件")
	flag.Parse()
	if *appPackage == "" {
		log.Println("读取配置文件失败")
	}
	file, error := os.Open(*appPackage)
	if error != nil {
		log.Println(error)
	}
	defer file.Close()
	data, error := ioutil.ReadAll(file)
	if error != nil {
		log.Println(error.Error())
	}
	error = xml.Unmarshal(data, &config)
	if error != nil {
		log.Println(error.Error())
	}
}

func (this *Service) DataBase() database.IDataBase {
	if this.iDatabase == nil {
		newConfig := new(database.DatabaseConfig)
		initialConfig := config.DatabaseConfig
		if initialConfig == nil {
			return nil
		}
		databaseType := initialConfig.Type
		if databaseType == "" {
			return nil
		}
		hostName := initialConfig.HostName
		if hostName == "" {
			hostName = os.Getenv("DATABASE_HOST")
		}
		if hostName == "" {
			hostName = "127.0.0.1"
		}
		newConfig.HostName = hostName
		hostPort := initialConfig.HostPort
		if hostPort == 0 {
			port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
			hostPort = port
		}
		if hostPort == 0 {
			hostPort = 27017
		}
		newConfig.HostPort = hostPort
		userName := initialConfig.UserName
		if userName == "" {
			userName = os.Getenv("DATABASE_USERNAME")
		}
		newConfig.UserName = userName
		password := initialConfig.Password
		if password == "" {
			password = os.Getenv("DATABASE_PASSWORD")
		}
		newConfig.Password = password
		databaseName := initialConfig.DataBaseName
		if databaseName == "" {
			databaseName = os.Getenv("DATABASE_NAME")
		}
		newConfig.DataBaseName = databaseName
		mongosession := mongodb.NewMongoDataBase(newConfig)
		mongosession.Connect()
		this.iDatabase = mongosession
	}
	return this.iDatabase
}


