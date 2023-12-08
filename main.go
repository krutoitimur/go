package main

import (
	"beastmode/packages/routes"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	//database connect

	db, err := sqlx.Connect("mysql", "root:@/danil")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println("Error closing database:", err)
		}
	}(db)

	//routes
	r := mux.NewRouter()

	// Определите маршруты здесь
	r.PathPrefix("/auth").Handler(routes.Auth(db))
	r.PathPrefix("/registration").Handler(routes.Reg(db))
	//serve
	fmt.Print("Server is running on port :8080")
	http.ListenAndServe(":8080", r)
}
