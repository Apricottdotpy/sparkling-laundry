package menu

import (
	"enigma-laundry/libquery"
	"fmt"
	"os"
	"strings"
)

func MenuTransaksi() {
	var entry int
	for {
		fmt.Println(strings.Repeat("=", 58))
		fmt.Println("=====================Enigma Laundry ======================")
		fmt.Printf("1. Lihat Data Transaksi \n2. Tambah Transaksi Baru \n3. Kembali Ke Menu Sebelumnya\n4. Keluar\n")
		fmt.Println(strings.Repeat("=", 58))

		fmt.Print("Masukkan nomor pilihan anda! ")
		fmt.Scanln(&entry)

		switch entry {
		case 0:
			fmt.Println("Pilihan hanya tersedia dari 1 sampai dengan 4")
		case 1:
			MenuViewTransaksi()
		case 2:
			err := libquery.AddTransaksi()
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			MainMenu()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Pilihan hanya tersedia dari 1 sampai dengan 4")
		}
	}
}
