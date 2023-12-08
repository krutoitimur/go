package main

import (
	"beastmode/packages/db"
	"beastmode/packages/routes"

	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	//database connect
	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("Error in connection database:", err)
	}

	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println("Error closing database:", err)
		}
	}(db.GetDB())

	//routes
	r := mux.NewRouter()

	r.PathPrefix("/auth").Handler(routes.Auth(db.GetDB()))
	r.PathPrefix("/registration").Handler(routes.Reg(db.GetDB()))
	//serve
	fmt.Print("Server is running on port :8080")
	http.ListenAndServe(":8080", r)
}
