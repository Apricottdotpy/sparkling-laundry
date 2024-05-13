package menu

import (
	"fmt"
	"os"
	"strings"
)

func MainMenu() {
	var entryMain int
	for {
		fmt.Println(strings.Repeat("=", 58))
		fmt.Println("=====================Enigma Laundry ======================")
		fmt.Printf("1. Data Pelanggan \n2. Data Transaksi \n3. Keluar\n")
		fmt.Println(strings.Repeat("=", 58))

		fmt.Print("Masukkan nomor pilihan anda! ")
		fmt.Scanln(&entryMain)

		switch entryMain {
		case 0:
			fmt.Println("Pilihan hanya tersedia dari 1 sampai dengan 6")
		case 1:
			MenuPelanggan()
		case 2:
			MenuTransaksi()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Pilihan hanya tersedia dari 1 sampai dengan 3")
		}
	}
}
