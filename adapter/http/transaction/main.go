package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sousaprogramador/session-finance-go-lang/model/transaction"
)

var transactions = transaction.Transactions{}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body, _ = io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &transactions)

	fmt.Println(transactions)
}
