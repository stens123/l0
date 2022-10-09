package Backend

import (
	"awesomeProject/srcs/Backend/JsonStruct"
	pq "awesomeProject/srcs/Backend/Postgresql"
	"awesomeProject/srcs/Backend/Utils"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

type CommonBackend struct {
	DataBase    pq.Postgresql
	ConnectStan *StanManager
	JModelSlice JsonStruct.JsonSlice
}

func BackEnd(Configs *Utils.Configs) *CommonBackend {
	var backend CommonBackend
	backend.DataBase.Connect(Configs, time.Second*3)
	backend.JModelSlice = JsonStruct.NewJsonSlice()
	ReadFromDataBase(&backend)
	backend.ConnectStan = NewConnect(Configs)
	ModelSubscribe(&backend, Configs.ModelSubj)
	return &backend
}

func (c *CommonBackend) Close() {
	c.DataBase.Disconnect()
	c.ConnectStan.UnscribeAll()
}

func ReadFromDataBase(bk *CommonBackend) {
	var jsonData []byte
	query, err := bk.DataBase.GetRaw().Query("SELECT model FROM models;")
	if err != nil {
		log.Panic(err)
		return
	}
	for query.Next() {
		err = query.Scan(&jsonData)
		if err != nil {
			log.Panic(err)
		}
		bk.JModelSlice.Lock()
		err := bk.JModelSlice.AddFromData(jsonData)
		bk.JModelSlice.Unlock()
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func ModelSubscribe(bk *CommonBackend, subject string) {
	fmt.Println("\033[36m"+"Количество элементов в кэше ранее", len(bk.JModelSlice.GetSlice()), "\033[0m")
	bk.ConnectStan.NewSubscribe(&subject, func(msg *stan.Msg) {
		js, ok := JsonStruct.ParseBytes(msg.Data)
		if ok != nil {
			fmt.Println("\u001B[31m" + "Неверный формат JSON" + "\033[0m")
			fmt.Println(ok.Error())
			return
		}

		_, ok = bk.DataBase.GetRaw().Exec("INSERT INTO models (model) VALUES ($1)", msg.Data)
		if ok != nil {
			log.Println(ok)
			return
		}
		bk.JModelSlice.Lock()
		defer bk.JModelSlice.Unlock()
		bk.JModelSlice.Add(js)
		fmt.Println("\033[32m"+"Элементов в кэше: ", len(bk.JModelSlice.GetSlice()), "\b"+"\033[0m")
	})
}
