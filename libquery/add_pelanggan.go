package libquery

import (
	"bufio"
	"enigma-laundry/database"
	"enigma-laundry/entity"
	"enigma-laundry/validasi"
	"fmt"
	"os"
	"strconv"
)

func AddDataPelanggan() error {

	db := database.ConnectDb()
	defer db.Close()

	var NewPelanggan entity.Customers

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("ID Pelanggan Baru: ")
	scanner.Scan()
	NewPelanggan.Id, _ = strconv.Atoi(scanner.Text())
	_, err := validasi.FindCustomerById(NewPelanggan.Id)
	if err != nil {

		fmt.Print("Masukkan Nama Pelanggan Baru: ")
		scanner.Scan()
		NewPelanggan.Name = scanner.Text()

		fmt.Print("Masukkan Nomor HP Pelanggan Baru: ")
		scanner.Scan()
		NewPelanggan.NoHp = scanner.Text()
		var yrn string
		fmt.Print("Apakah anda yakin ingin menambahkan data pelanggan ini (Y/N)? ")
		fmt.Scanln(&yrn)
		if yrn == "Y" || yrn == "N" {
			if yrn == "Y" {
				sqlStatement := "INSERT INTO mst_pelanggan (id, nama_pelanggan, nomor_hp) VALUES ($1, $2, $3);"
				_, err := db.Exec(sqlStatement, NewPelanggan.Id, NewPelanggan.Name, NewPelanggan.NoHp)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Berhasil menambahkan data pelanggan baru!")
				}
			} else {
				fmt.Println("Anda telah membatalkan penambahan pelanggan baru!")
			}
		} else {
			fmt.Println("Hanya terdapat pilihan 'Y' dan 'N'!")
		}

		return nil
	}
	return fmt.Errorf("nama pelanggan dengan id: %d sudah ada", NewPelanggan.Id)
}
