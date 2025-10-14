package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var db *sql.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
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
