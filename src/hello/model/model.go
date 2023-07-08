package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type SystemEvents struct {
	DeviceReportedTime string `db:"DeviceReportedTime"`
	FromHost           string `db:"FromHost"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "rsyslog:david@tcp(127.0.0.1:3306)/Syslog")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	Db = database
	//defer Db.Close()
}

func get_log_info() {
	var info []SystemEvents
	err := Db.Select(&info, "select min(DeviceReportedTime) as DeviceReportedTime, FromHost from SystemEvents group by FromHost;")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select ok:", info)
}
