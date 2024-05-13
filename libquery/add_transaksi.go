package libquery

import (
	"enigma-laundry/database"
	"enigma-laundry/entity"
	"enigma-laundry/validasi"
	"fmt"
	"strings"
)

func AddTransaksi() error {
	db := database.ConnectDb()
	defer db.Close()

	var id int
	err := db.QueryRow("SELECT COUNT(*) FROM transaksi;").Scan(&id)
	if err != nil {
		panic(err)
	}

	id += 1

	var transaksi entity.Transaksi
	fmt.Println("\nMasukkan Data Transaksi:")
	fmt.Print("ID Pelanggan Yang Ingin Bertransaksi: ")
	fmt.Scanln(&transaksi.IdPelanggan)

	_, err = validasi.FindCustomerById(transaksi.IdPelanggan)
	if err != nil {
		return err
	}

	fmt.Print("ID Pegawai: ")
	fmt.Scanln(&transaksi.IdPegawai)

	_, err = validasi.FindEmployeById(transaksi.IdPegawai)
	if err != nil {
		return err
	}

	fmt.Print("Perkiraan Tanggal Selesai (DD-MM-YYYY): ")
	fmt.Scanln(&transaksi.TanggalSelesai)

	tglArray := strings.Split(transaksi.TanggalSelesai, "-")
	reversed := make([]string, len(tglArray))

	for i := 0; i < len(tglArray); i++ {
		reversed[i] = tglArray[len(tglArray)-1-i]
	}

	tgl := strings.Join(reversed, "-")

	var detailTransaksi entity.DetailTransaksi
	fmt.Println("1. Laundry Pakaian\n2. Laundry dan Setrika Pakaian\n3. Laundry Bedcover\n4. Laundry Boneka")
	fmt.Print("Masukkan Nomor Layanan: ")
	fmt.Scanln(&detailTransaksi.IdLayanan)
	if detailTransaksi.IdLayanan <= 0 || detailTransaksi.IdLayanan > 4 {
		fmt.Println("Nomor Layanan Yang Anda Masukkan Tidak Tersedia!")
		return nil
	}

	fmt.Print("Quantity: ")
	fmt.Scanln(&detailTransaksi.Quantity)

	fmt.Print("Status Pembayaran [Lunas / Belum Lunas]: ")
	fmt.Scanln(&transaksi.StatusPembayaran)

	var yrn string
	fmt.Print("Input Transaksi Baru? [Y/N]: ")
	fmt.Scanln(&yrn)
	if yrn == "Y" || yrn == "N" {
		if yrn == "Y" {
			_, err = db.Exec("INSERT INTO transaksi (id, id_pelanggan, id_pegawai, tanggal_selesai, status_pembayaran) VALUES ($1, $2, $3, $4, $5)", id, transaksi.IdPelanggan, transaksi.IdPegawai, tgl, transaksi.StatusPembayaran)
			if err != nil {
				return err
			}

			_, err = db.Exec("INSERT INTO detail_transaksi (id_transaksi, id_layanan, quantity) VALUES ($1, $2, $3)", id, detailTransaksi.IdLayanan, detailTransaksi.Quantity)
			if err != nil {
				return err
			}

			fmt.Println("\nTransaksi Baru Berhasil Ditambahkan!")
		} else {
			fmt.Println("Input transaksi dibatalkan!")
		}

	} else {
		fmt.Println("Hanya terdapat pilihan 'Y' dan 'N'!")
	}
	return nil
}
