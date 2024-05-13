package validasi

import (
	"enigma-laundry/database"
	"enigma-laundry/entity"
	"fmt"
)

func FindTransactionById(id int) (entity.Transaksi, error) {
	db := database.ConnectDb()

	var transaksi entity.Transaksi

	sqlStatement := "SELECT * FROM transaksi WHERE id = $1;"

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&transaksi.Id, &transaksi.IdPelanggan, &transaksi.IdPegawai, &transaksi.TanggalMasuk, &transaksi.TanggalSelesai, &transaksi.StatusPembayaran)
	if err != nil {
		return transaksi, fmt.Errorf("transaksi dengan id %d tidak dapat ditemukan", id)
	}
	return transaksi, nil
}
