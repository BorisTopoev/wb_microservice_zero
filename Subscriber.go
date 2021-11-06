package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"main/envContainer"
	"main/structs"
	"os"
	"time"
)

func processMsg(m *stan.Msg) {
	var order structs.Order
	fmt.Println("Received a message")
	json.Unmarshal(m.Data, &order)
	fmt.Println(order)
	payment := order.Payment
	//fmt.Println(payment)
	insertItem(order.Items)
	insertPayment(payment)
	insertOrder(order)
	insertOrderItems(order)

}
func NatsConnect(envContainer envContainer.EnvContainer) {
	nc, err := nats.Connect("wbx-world-nats-stage.dp.wb.ru,wbx-world-nats-stage.dl.wb.ru")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to nats. nats: %v\n", err)
		os.Exit(1)
	}

	sc, _ := stan.Connect("world-nats-stage", "meme", stan.NatsConn(nc))
	sc.Subscribe("go.test", func(m *stan.Msg) {
		go processMsg(m)
	})
	time.Sleep(4 * time.Minute)
}
