package validasi

import (
	"enigma-laundry/database"
	"fmt"
	"strings"
)

func CheckTransactionByCustomerName(name string) (bool, error) {
	db := database.ConnectDb()
	defer db.Close()

	var count int
	sqlStatement := "SELECT COUNT(*) FROM transaksi JOIN mst_pelanggan ON transaksi.id_pelanggan = mst_pelanggan.id WHERE LOWER(mst_pelanggan.nama_pelanggan) LIKE $1 || '%';"

	row := db.QueryRow(sqlStatement, strings.ToLower(name))
	err := row.Scan(&count)
	if err != nil {
		return false, fmt.Errorf("customer dengan nama %s belum pernah melakukan transaksi", name)
	}

	return count > 0, nil
}
