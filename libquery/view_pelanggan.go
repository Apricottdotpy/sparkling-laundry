package libquery

import (
	"enigma-laundry/database"
	"enigma-laundry/entity"
	"fmt"
	"strings"
)

func ViewDataPelanggan() error {
	db := database.ConnectDb()
	defer db.Close()
	var err error

	sqlStatement := "SELECT * FROM mst_pelanggan;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var customers []entity.Customers
	for rows.Next() {
		var customer entity.Customers
		err = rows.Scan(&customer.Id, &customer.Name, &customer.NoHp)
		if err != nil {
			panic(err)
		}
		customers = append(customers, customer)
		for i, customer := range customers {
			fmt.Println("Pelanggan -", i+1)
			fmt.Println("ID Pelanggan: ", customer.Id)
			fmt.Println("Nama Pelanggan: ", customer.Name)
			fmt.Println("Nomor HP: ", customer.NoHp)
			fmt.Println(strings.Repeat("-", 58))
		}
	}
	return nil
}
