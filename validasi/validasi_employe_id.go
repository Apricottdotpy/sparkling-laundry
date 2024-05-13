package validasi

import (
	"enigma-laundry/database"
	"enigma-laundry/entity"
	"fmt"
)

func FindEmployeById(id int) (entity.Pegawai, error) {
	db := database.ConnectDb()

	var pegawai entity.Pegawai

	sqlStatement := "SELECT * FROM mst_pegawai WHERE id = $1;"

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&pegawai.Id, &pegawai.Name)
	if err != nil {
		return pegawai, fmt.Errorf("pegawai dengan id %d tidak dapat ditemukan", id)
	}
	return pegawai, nil
}
