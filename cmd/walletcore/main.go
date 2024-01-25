package main

import (
	"context"
	"database/sql"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/danyukod/chave-pix-utils/pkg/events"
	"github.com/danyukod/chave-pix-utils/pkg/kafka"
	"github.com/danyukod/chave-pix-utils/pkg/uow"
	"github.com/danyukod/wallet-core-go/internal/database"
	"github.com/danyukod/wallet-core-go/internal/event"
	"github.com/danyukod/wallet-core-go/internal/event/handler"
	"github.com/danyukod/wallet-core-go/internal/gateway"
	accountUsecase "github.com/danyukod/wallet-core-go/internal/usecase/create_account"
	clientUsecase "github.com/danyukod/wallet-core-go/internal/usecase/create_client"
	transactionUsecase "github.com/danyukod/wallet-core-go/internal/usecase/create_transaction"
	"github.com/danyukod/wallet-core-go/internal/web"
	"github.com/danyukod/wallet-core-go/internal/web/webserver"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	err = eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	if err != nil {
		panic(err)
	}
	transactionCreatedEvent := event.NewTransactionCreated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	clientGateway := gateway.NewClientGateway(*clientDb)
	accountGateway := gateway.NewAccountGateway(*accountDb)

	createClientInteract := clientUsecase.NewCreateClientInteract(clientGateway)
	createAccountInteract := accountUsecase.NewCreateAccountInteract(accountGateway, clientGateway)
	createTransactionInteract := transactionUsecase.NewCreateTransactionInteract(uow, eventDispatcher, transactionCreatedEvent)

	webServer := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(createClientInteract)
	accountHandler := web.NewWebAccountHandler(createAccountInteract)
	transactionHandler := web.NewWebTransactionHandler(createTransactionInteract)

	webServer.AddHandler("/clients", clientHandler.CreateClient)
	webServer.AddHandler("/accounts", accountHandler.CreateAccount)
	webServer.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	err = webServer.Start()
	if err != nil {
		return
	}
}
