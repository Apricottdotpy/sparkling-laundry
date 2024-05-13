package validasi

import (
	"enigma-laundry/database"
	"enigma-laundry/entity"
	"fmt"
)

func FindCustomerByName(name string) (entity.Customers, error) {
	db := database.ConnectDb()

	var customer entity.Customers

	sqlStatement := "SELECT * FROM mst_pelanggan WHERE LOWER(mst_pelanggan.nama_pelanggan) LIKE $1 || '%';"

	row := db.QueryRow(sqlStatement, name)
	err := row.Scan(&customer.Id, &customer.Name, &customer.NoHp)
	if err != nil {
		return customer, fmt.Errorf("customer dengan nama %s tidak dapat ditemukan", name)
	}
	return customer, nil
}
