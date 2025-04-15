// package initialize

// import (
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	"system-management-pg/global"

// 	_ "github.com/go-sql-driver/mysql"
// 	"go.uber.org/zap"
// )

// // checkErrorPanic ghi log và panic nếu có lỗi
// func checkErrorPanic(err error, errString string) {
// 	if err != nil {
// 		global.Logger.Error(errString, zap.Error(err))
// 		panic(err)
// 	}
// }

// func InitMysqlC() {
// 	m := global.Config.Mysql
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true&loc=Local&charset=utf8mb4",
// 		m.Username, m.Password, m.Host, m.Port, m.Dbname)

// 	db, err := sql.Open("mysql", dsn)
// 	checkErrorPanic(err, "InitMysqlc initialization error")

// 	err = db.Ping()
// 	checkErrorPanic(err, "MySQLC ping error")

// 	global.Logger.Info("Initializing MySQLC Successfully")
// 	global.Mdbc = db

// 	SetPool()
// }

// func SetPool() {
// 	m := global.Config.Mysql
// 	sqlDb := global.Mdbc

// 	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
// 	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
// 	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
// }

package initialize

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"system-management-pg/global"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// MysqlDBSingleton quản lý kết nối database theo mô hình Singleton
type MysqlDBSingleton struct {
	db *sql.DB
}

// instance là biến toàn cục lưu trữ instance duy nhất
var (
	instance *MysqlDBSingleton
	once     sync.Once
)

// checkErrorPanic ghi log và panic nếu có lỗi
func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

// GetMysqlInstance trả về instance duy nhất của database
func GetMysqlInstance() *sql.DB {
	once.Do(func() {
		instance = &MysqlDBSingleton{}
		instance.initDB()
	})
	return instance.db
}

// initDB khởi tạo kết nối database
func (m *MysqlDBSingleton) initDB() {
	mysqlConfig := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Dbname)

	db, err := sql.Open("mysql", dsn)
	checkErrorPanic(err, "MySQL initialization error")

	// Kiểm tra kết nối
	err = db.Ping()
	checkErrorPanic(err, "MySQL ping error")

	// Thiết lập connection pool
	m.configurePool(db)

	m.db = db
	global.Mdbc = db
	global.Logger.Info("Initialized MySQL successfully with Singleton pattern")
}

// configurePool thiết lập các tham số cho connection pool
func (m *MysqlDBSingleton) configurePool(db *sql.DB) {
	mysqlConfig := global.Config.Mysql

	db.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	db.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(mysqlConfig.ConnMaxLifetime) * time.Second)
	// db.SetConnMaxIdleTime(time.Duration(mysqlConfig.MaxIdleTime) * time.Second)
}

// Close đóng kết nối database
func (m *MysqlDBSingleton) Close() error {
	if m.db != nil {
		return m.db.Close()
	}
	return nil
}

// InitMysqlC khởi tạo database với Singleton pattern
func InitMysqlC() {
	global.Mdbc = GetMysqlInstance()
}