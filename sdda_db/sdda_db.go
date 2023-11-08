package main

import (
	"fmt"
	"os"
	"log"
	"database/sql"
	"github.com/lib/pq"
)

connectionString := "postgressql://sdda-app:<password>@172.20.0.2/sdda-app?sslmode=disable"

func placeholder() {
	fmt.Print("Wait")
}
