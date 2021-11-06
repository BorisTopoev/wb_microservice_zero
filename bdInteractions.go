package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"main/envContainer"
	"main/structs"
	"os"
)

func getConnect() *pgx.Conn {
	var envContainer = envContainer.Get_env()
	conn, err := pgx.Connect(context.Background(), envContainer.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
func insertPayment(payment structs.Payment) {
	if payment.Transaction == "" {
		return
	}
	conn := getConnect()
	row := conn.QueryRow(context.Background(),
		"INSERT INTO payment (transaction, currency, provider, amount, paymentdt, bank, deliverycost, goodstotal) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		payment.Transaction, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal)
	fmt.Println(row)
	return
}
func insertItem(items []structs.Items) {
	for _, element := range items {
		if element.ChrtID == 0 {
			return
		}
		conn := getConnect()
		row := conn.QueryRow(context.Background(),
			"INSERT INTO items (chrtid, Price, Rid, Name, Sale, Size, TotalPrice, NmID, Brand) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
			element.ChrtID, element.Price, element.Rid, element.Name, element.Sale, element.Size, element.TotalPrice, element.NmID, element.Brand)
		fmt.Println(row)
		var id uint64
		err := row.Scan(&id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to INSERT: %v\n", err)
		}
	}
	return
}
func insertOrder(order structs.Order) {
	if order.OrderUID == "" {
		return
	}
	conn := getConnect()
	row := conn.QueryRow(context.Background(),
		"INSERT INTO orders (OrderUID, Entry, InternalSignature, Payment, Locale, CustomerID, TrackNumber, DeliveryService, Shardkey, SmID) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		order.OrderUID, order.Entry, order.InternalSignature, order.Payment.Transaction, order.Locale, order.CustomerID, order.TrackNumber, order.DeliveryService, order.Shardkey, order.SmID)
	fmt.Println(row)
	var id uint64
	err := row.Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to INSERT: %v\n", err)
	}
	return
}
func insertOrderItems(order structs.Order) {
	if order.OrderUID == "" {
		return
	}
	for _, element := range order.Items {
		if element.ChrtID == 0 {
			return
		}
		conn := getConnect()
		row := conn.QueryRow(context.Background(),
			"INSERT INTO orderItems (orderuid, chrtid) VALUES ($1, $2)",
			order.OrderUID, element.ChrtID)
		fmt.Println(row)
		var id uint64
		err := row.Scan(&id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to INSERT: %v\n", err)
		}
	}

}
