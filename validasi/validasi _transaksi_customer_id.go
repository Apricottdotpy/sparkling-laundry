package validasi

import (
	"enigma-laundry/database"
	"enigma-laundry/entity"
	"fmt"
)

func CheckTransactionByCustomerId(id int) (entity.Transaksi, error) {
	db := database.ConnectDb()
	defer db.Close()

	var transaksi entity.Transaksi
	sqlStatement := "SELECT * FROM transaksi WHERE id_pelanggan = $1;"

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&transaksi.Id, &transaksi.IdPelanggan, &transaksi.IdPegawai, &transaksi.TanggalMasuk, &transaksi.TanggalSelesai, &transaksi.StatusPembayaran)
	if err != nil {
		return transaksi, fmt.Errorf("customer dengan id %d belum pernah melakukan transaksi", id)
	}
	return transaksi, nil
}
