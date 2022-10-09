package Utils

import (
	"flag"
	"time"
)

type Configs struct {
	UserDB, PassDB, AddrDB, NameDB, ClusterID, ClientID, ModelSubj string
}

func ParseArgs() (c *Configs) {
	c = &Configs{}
	flag.StringVar(&c.UserDB, "u", "***", "Имя пользовптеля")
	flag.StringVar(&c.NameDB, "d", "postgres", "Имя базы данных")
	flag.StringVar(&c.PassDB, "p", "", "Пароль")
	flag.StringVar(&c.AddrDB, "h", "localhost", "Адрес хоста")
	flag.StringVar(&c.ClusterID, "cid", "test-cluster", "cluster id of NATS-streaming")
	flag.StringVar(&c.ClientID, "cln", "server-1", "client name in NATS-connection")
	flag.StringVar(&c.ModelSubj, "sm", "jsonModel", "Объект канала")
	flag.Parse()
	return c
}

func TryDoIt(t time.Duration, attempts uint8, f func() error) (ok error) {
	ok = f()
	for ok != nil && attempts != 0 {
		time.Sleep(t)
		ok = f()
		attempts--
	}
	return ok
}
