package main

import (
	"database/sql"
	"fmt"

	"github.com/devmatheuus/pfa-go/internal/infra/database"
	"github.com/devmatheuus/pfa-go/internal/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repo)

	input := usecase.OrderInputDTO{
		ID:    "2",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Exec(&input)

	if err != nil {
		panic(err)
	}

	fmt.Println(output.FinalPrice)
}
