package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"os"
)

func readJson(filename string) []uint8 {
	plan, _ := ioutil.ReadFile(filename)
	return plan
}
func main() {
	nc, err := nats.Connect("wbx-world-nats-stage.dp.wb.ru,wbx-world-nats-stage.dl.wb.ru")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to nats. nats: %v\n", err)
		os.Exit(1)
	}
	sc, err := stan.Connect("world-nats-stage", "me", stan.NatsConn(nc))
	items, _ := ioutil.ReadDir("C:\\Users\\ASUS\\GolandProjects\\wb_microservice_zero\\data")
	for _, item := range items {
		coc := readJson("C:\\Users\\ASUS\\GolandProjects\\wb_microservice_zero\\data\\" + item.Name())
		sc.Publish("go.test", coc)
	}
}
