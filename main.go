package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"main/envContainer"
	"main/structs"
	"os"
)

func main() {
	var envContainer = envContainer.Get_env()
	NatsConnect(envContainer)
	return
	customItem := structs.Items{123, 12331, "23423", "sleeve", 12314, "M", 228, 12, "trasher"}
	customPayment := structs.Payment{
		Transaction:  "s",
		Currency:     "s",
		Provider:     "s",
		Amount:       2,
		PaymentDt:    3,
		Bank:         "asda",
		DeliveryCost: 1110,
		GoodsTotal:   15678,
	}
	customOrder := structs.Order{
		OrderUID:          "eeee",
		Entry:             "qweqwe",
		InternalSignature: "asdasd",
		Payment:           customPayment,
		Items:             []structs.Items{customItem},
		Locale:            "asdasdas",
		CustomerID:        "asdqwe",
		TrackNumber:       "wqeac",
		DeliveryService:   "eqweqwe",
		Shardkey:          "asdas",
		SmID:              1,
	}
	fmt.Println(customOrder)

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// internalsignature, payment, items, locale, customerid, tracknumber, deliveryservice, shardkey, smid

	row := conn.QueryRow(context.Background(),
		"INSERT INTO order (, currency) VALUES ($1, $2)",
		"transnum1", customPayment.Currency)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to INSERT: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	_, err = conn.Query(context.Background(), "drop table wtf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(conn)
	fmt.Println(err)
}
