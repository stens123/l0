package Postgresql

import (
	"awesomeProject/srcs/Backend/Utils"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Postgresql struct {
	connStr string
	open    *sql.DB
}

func (p *Postgresql) Connect(cnf *Utils.Configs, s time.Duration) {
	var ok error
	p.connStr = "postgresql://" + cnf.UserDB + ":" + cnf.PassDB + "@" + cnf.AddrDB + "/" + cnf.NameDB + "?sslmode=disable"
	ok = Utils.TryDoIt(s, 10, func() error {
		p.open, ok = sql.Open("postgres", p.connStr)
		return ok
	})
	if ok == nil {
		fmt.Println("\033[34m" + "DataBase" + "\033[32m" + " connected" + "\033[0m")
	} else {
		log.Panic(ok)
	}
}

func (p *Postgresql) Disconnect() {
	err := p.open.Close()
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("\033[34m"+"DataBase", "\033[31m"+"disconnected"+"\033[0m")
	}
}

func (p *Postgresql) GetRaw() *sql.DB {
	return p.open
}
