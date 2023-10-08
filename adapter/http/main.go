package http

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sousaprogramador/session-finance-go-lang/adapter/http/actuator"
	"github.com/sousaprogramador/session-finance-go-lang/adapter/http/transaction"
)

func Init() error {
	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	return http.ListenAndServe(":3333", nil)
}
