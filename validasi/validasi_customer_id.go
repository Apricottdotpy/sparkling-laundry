package validasi

import (
	"enigma-laundry/database"
	"enigma-laundry/entity"
	"fmt"
)

func FindCustomerById(id int) (entity.Customers, error) {
	db := database.ConnectDb()

	var customer entity.Customers

	sqlStatement := "SELECT * FROM mst_pelanggan WHERE id = $1;"

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&customer.Id, &customer.Name, &customer.NoHp)
	if err != nil {
		return customer, fmt.Errorf("customer dengan id %d tidak dapat ditemukan", id)
	}
	return customer, nil
}
