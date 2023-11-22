package main

import (
	"database/sql"
	"fmt"
	"github.com/danyukod/chave-pix-utils/pkg/events"
	"github.com/danyukod/wallet-core-go/internal/database"
	"github.com/danyukod/wallet-core-go/internal/event"
	"github.com/danyukod/wallet-core-go/internal/gateway"
	accountUsecase "github.com/danyukod/wallet-core-go/internal/usecase/create_account"
	clientUsecase "github.com/danyukod/wallet-core-go/internal/usecase/create_client"
	transactionUsecase "github.com/danyukod/wallet-core-go/internal/usecase/create_transaction"
	"github.com/danyukod/wallet-core-go/internal/web"
	"github.com/danyukod/wallet-core-go/internal/web/webserver"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	clientGateway := gateway.NewClientGateway(*clientDb)
	accountGateway := gateway.NewAccountGateway(*accountDb)
	transactionGateway := gateway.NewTransactionGateway(*transactionDb)

	createClientInteract := clientUsecase.NewCreateClientInteract(clientGateway)
	createAccountInteract := accountUsecase.NewCreateAccountInteract(accountGateway, clientGateway)
	createTransactionInteract := transactionUsecase.NewCreateTransactionInteract(transactionGateway, accountGateway, eventDispatcher, transactionCreatedEvent)

	webServer := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(createClientInteract)
	accountHandler := web.NewWebAccountHandler(createAccountInteract)
	transactionHandler := web.NewWebTransactionHandler(createTransactionInteract)

	webServer.AddHandler("/clients", clientHandler.CreateClient)
	webServer.AddHandler("/accounts", accountHandler.CreateAccount)
	webServer.AddHandler("/transactions", transactionHandler.CreateTransaction)

	err = webServer.Start()
	if err != nil {
		return
	}
}
