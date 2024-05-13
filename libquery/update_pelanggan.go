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

func UpdateDataPelanggan() error {

	db := database.ConnectDb()
	defer db.Close()

	var NewPelanggan entity.Customers

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan ID Pelanggan Yang Ingin Diubah: ")
	scanner.Scan()
	NewPelanggan.Id, _ = strconv.Atoi(scanner.Text())
	pelanggan, err := validasi.FindCustomerById(NewPelanggan.Id)
	if err != nil {
		return err
	}

	fmt.Print("Masukkan Nama Pelanggan Baru: ")
	scanner.Scan()
	NewPelanggan.Name = scanner.Text()
	if NewPelanggan.Name == "" {
		NewPelanggan.Name = pelanggan.Name
	}

	fmt.Print("Masukkan Nomor HP Pelanggan Baru: ")
	scanner.Scan()
	NewPelanggan.NoHp = scanner.Text()
	if NewPelanggan.NoHp == "" {
		NewPelanggan.NoHp = pelanggan.NoHp
	}

	var yrn string
	fmt.Print("Apakah anda yakin ingin mengupdate data pelanggan ini (Y/N)? ")
	fmt.Scanln(&yrn)
	if yrn == "Y" || yrn == "N" {
		if yrn == "Y" {
			sqlStatement := "UPDATE mst_pelanggan SET nama_pelanggan = $2, nomor_hp = $3 WHERE id = $1"
			_, err := db.Exec(sqlStatement, NewPelanggan.Id, NewPelanggan.Name, NewPelanggan.NoHp)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Anda telah berhasil mengubah data pelanggan dengan id", NewPelanggan.Id)
			}

		} else {
			fmt.Println("Perubahan data pelanggan dibatalkan!")
		}
	} else {
		fmt.Println("Hanya terdapat pilihan 'Y' dan 'N'!")
	}

	return nil
}
