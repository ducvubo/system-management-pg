package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"system-management-pg/global"

	"go.uber.org/zap"
	"gorm.io/gen"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlC() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", s)
	checkErrorPanicC(err, "InitMysql initialization error")
	global.Logger.Info("Initializing MySQL Successfully sql")
	global.Mdbc = db

	SetPoolC()
}

func SetPoolC() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func genTableDAOC() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(global.Mdb) 
	g.GenerateModel("go_crm_user")
	g.Execute()
}
func migrateTablesC() {
	err := global.Mdb.AutoMigrate(
	)
	if err != nil {
		fmt.Println("Migrating tables error:", err)
	}
}
