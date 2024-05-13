package libquery

import (
	"enigma-laundry/database"
	"enigma-laundry/entity"
	"enigma-laundry/validasi"
	"fmt"
	"strings"
)

func ViewTransaksiPelangganNama() error {
	var err error
	var namaPelanggan string
	fmt.Print("Masukkan Nama Pelanggan yang Ingin Dicari: ")
	fmt.Scanln(&namaPelanggan)

	_, err = validasi.FindCustomerByName(namaPelanggan)
	if err != nil {
		return err
	}

	_, err = validasi.CheckTransactionByCustomerName(namaPelanggan)
	if err != nil {
		return err
	}

	db := database.ConnectDb()
	defer db.Close()

	namaPelanggan = strings.ToLower(namaPelanggan)

	sqlStatement := `
        SELECT 
            mst_pelanggan.nama_pelanggan, 
            layanan.pelayanan, 
            detail_transaksi.quantity, 
            mst_pegawai.nama_pegawai, 
            transaksi.tanggal_masuk, 
            layanan.harga * detail_transaksi.quantity AS total_harga 
        FROM 
            mst_pelanggan 
        JOIN 
            transaksi ON mst_pelanggan.id = transaksi.id_pelanggan 
        JOIN 
            detail_transaksi ON transaksi.id = detail_transaksi.id_transaksi 
        JOIN 
            layanan ON detail_transaksi.id_layanan = layanan.id 
        JOIN 
            mst_pegawai ON transaksi.id_pegawai = mst_pegawai.id 
        WHERE 
            LOWER(mst_pelanggan.nama_pelanggan) LIKE '%' || $1 || '%';
    `

	rows, err := db.Query(sqlStatement, namaPelanggan)
	if err != nil {
		return err
	}
	defer rows.Close()

	var customers []entity.Customers
	var pegawai []entity.Pegawai
	var layanan []entity.Layanan
	var transaksi []entity.Transaksi
	var detailTransaksi []entity.DetailTransaksi

	for rows.Next() {
		var customer entity.Customers
		var employee entity.Pegawai
		var service entity.Layanan
		var transaction entity.Transaksi
		var detail entity.DetailTransaksi

		err = rows.Scan(
			&customer.Name,
			&service.Pelayanan,
			&detail.Quantity,
			&employee.Name,
			&transaction.TanggalMasuk,
			&detail.TotalHarga,
		)
		if err != nil {
			return err
		}

		customers = append(customers, customer)
		pegawai = append(pegawai, employee)
		layanan = append(layanan, service)
		transaksi = append(transaksi, transaction)
		detailTransaksi = append(detailTransaksi, detail)
	}

	for i, customer := range customers {
		fmt.Println("Pelanggan -", i+1)
		fmt.Println("Nama Pelanggan: ", customer.Name)
		fmt.Println("Pelayanan: ", layanan[i].Pelayanan)
		fmt.Println("Quantity: ", detailTransaksi[i].Quantity)
		fmt.Println("Nama Pegawai: ", pegawai[i].Name)
		fmt.Println("Tanggal Masuk: ", transaksi[i].TanggalMasuk)
		fmt.Println("Total Harga: ", detailTransaksi[i].TotalHarga)
		fmt.Println(strings.Repeat("-", 58))
	}

	return nil
}
