package mysql

import (
	"WebApp/settings"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return err
	}
	if err = db.Ping(); err != nil {
		fmt.Printf("connect to mysql failed, err:%v\n", err)
		return err
	}
	return nil
}

func Close() {
	_ = db.Close()
}
