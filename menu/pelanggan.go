package menu

import (
	"enigma-laundry/libquery"
	"fmt"
	"os"
	"strings"
)

func MenuPelanggan() {

	var entry int
	for {
		fmt.Println(strings.Repeat("=", 58))
		fmt.Println("=====================Enigma Laundry ======================")
		fmt.Printf("1. Lihat Data Pelanggan \n2. Tambah Data Pelanggan Baru \n3. Perbarui Data Pelanggan\n4. Hapus Data Pelanggan\n5. Kembali Ke Menu Sebelumnya\n6. Keluar\n")
		fmt.Println(strings.Repeat("=", 58))

		fmt.Print("Masukkan nomor pilihan anda! ")
		fmt.Scanln(&entry)

		switch entry {
		case 0:
			fmt.Println("Pilihan hanya tersedia dari 1 sampai dengan 6")
		case 1:
			err := libquery.ViewDataPelanggan()
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			err := libquery.AddDataPelanggan()
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			err := libquery.UpdateDataPelanggan()
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			err := libquery.DeleteDataPelanggan()
			if err != nil {
				fmt.Println(err)
			}
		case 5:
			MainMenu()
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Pilihan hanya tersedia dari 1 sampai dengan 6")
		}
	}

}
