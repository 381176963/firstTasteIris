package config

import (
	"sync"
	"time"

	"firstTasteIris/backend/files"
	"firstTasteIris/backend/transformer"
	"github.com/kataras/iris/v12"
	gf "github.com/snowlyg/gotransformer"
)

type config struct {
	Tc  *transformer.Conf
	Isc iris.Configuration
}

var cfg *config

var once sync.Once // sync.Once能确保实例化对象Do方法在多线程环境只运行一次,内部通过互斥锁实现

func GetAppUrl() string {
	return getTc().App.Url
}

func getTc() *transformer.Conf {
	return singleton().Tc
}

func singleton() *config {
	once.Do(func() { // sync.Once能确保实例化对象Do方法在多线程环境只运行一次,内部通过互斥锁实现
		path := files.GetAbsPath("./config/config")
		isc := iris.TOML(path) // 加载配置文件
		tc := getTfConf(isc)
		cfg = &config{Tc: tc, Isc: isc}
	})
	return cfg
}

func getTfConf(isc iris.Configuration) *transformer.Conf {
	app := transformer.App{}
	g := gf.NewTransform(&app, isc.Other["App"], time.RFC3339)
	_ = g.Transformer()

	db := transformer.Mysql{}
	g.OutputObj = &db
	g.InsertObj = isc.Other["Mysql"]
	_ = g.Transformer()

	mongodb := transformer.Mongodb{}
	g.OutputObj = &mongodb
	g.InsertObj = isc.Other["Mongodb"]
	_ = g.Transformer()

	redis := transformer.Redis{}
	g.OutputObj = &redis
	g.InsertObj = isc.Other["Redis"]
	_ = g.Transformer()

	sqlite := transformer.Sqlite{}
	g.OutputObj = &sqlite
	g.InsertObj = isc.Other["Sqlite"]
	_ = g.Transformer()

	admin := transformer.Admin{}
	g.OutputObj = &admin
	g.InsertObj = isc.Other["Admin"]
	_ = g.Transformer()

	testData := transformer.TestData{}
	g.OutputObj = &testData
	g.InsertObj = isc.Other["TestData"]
	_ = g.Transformer()

	return &transformer.Conf{
		App:      app,
		Mysql:    db,
		Mongodb:  mongodb,
		Redis:    redis,
		Sqlite:   sqlite,
		Admin:    admin,
		TestData: testData,
	}
}

func GetIrisConf() iris.Configuration {
	return singleton().Isc
}

func GetAppLoggerLevel() string {
	return getTc().App.LoggerLevel
}

func GetDeployment() string {
	return getTc().App.Deployment
}

func GetApiVersion() string {
	return getTc().App.ApiVersion
}

func GetAppDriverType() string {
	return getTc().App.DriverType
}

func GetSqliteTConnect() string {
	return files.GetAbsPath(getTc().Sqlite.TConnect)
}

func GetSqliteConnect() string {
	return files.GetAbsPath(getTc().Sqlite.Connect)
}

func GetMysqlConnect() string {
	return getTc().Mysql.Connect
}

func GetMysqlTName() string {
	return getTc().Mysql.TName
}

func GetMysqlName() string {
	return getTc().Mysql.Name
}