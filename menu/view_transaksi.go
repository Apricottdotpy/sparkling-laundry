package menu

import (
	"enigma-laundry/libquery"
	"fmt"
	"os"
	"strings"
)

func MenuViewTransaksi() {
	var entry int
	for {
		fmt.Println(strings.Repeat("=", 58))
		fmt.Println("=====================Enigma Laundry ======================")
		fmt.Printf("Cari transaksi berdasarkan\n1. ID Pelanggan \n2. Nama Pelanggan \n3. ID Transaksi\n4. Kembali Ke Menu Sebelumnya\n5. Keluar\n")
		fmt.Println(strings.Repeat("=", 58))

		fmt.Print("Masukkan nomor pilihan anda! ")
		fmt.Scanln(&entry)

		switch entry {
		case 0:
			fmt.Println("Pilihan hanya tersedia dari 1 sampai dengan 5")
		case 1:
			err := libquery.ViewTransaksiPelangganId()
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			err := libquery.ViewTransaksiPelangganNama()
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			err := libquery.ViewTransaksiById()
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			MenuTransaksi()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Pilihan hanya tersedia dari 1 sampai dengan 5")
		}
	}
}
