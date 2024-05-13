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

func DeleteDataPelanggan() error {

	var deletedId entity.Customers
	var yrn string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan ID pelanggan yang ingin dihapus: ")
	scanner.Scan()
	deletedId.Id, _ = strconv.Atoi(scanner.Text())
	_, err := validasi.FindCustomerById(deletedId.Id)
	if err != nil {
		return err
	}

	fmt.Print("Apakah anda yakin ingin menghapus data pelanggan dengan ID ", deletedId.Id, "? [Y/N] ")
	fmt.Scanln(&yrn)
	if yrn == "Y" || yrn == "N" {
		if yrn == "Y" {
			db := database.ConnectDb()
			defer db.Close()

			sqlStatement := "DELETE FROM mst_pelanggan WHERE id = $1;"

			_, err = db.Exec(sqlStatement, deletedId.Id)
			if err != nil {
				panic(err)
			} else {
				fmt.Println("Anda telah berhasil menghapus data pelanggan dengan id", deletedId.Id)
			}
		} else {
			fmt.Println("Penghapusan data pelanggan dibatalkan")
		}
	} else {
		fmt.Println("Hanya terdapat pilihan 'Y' dan 'N'!")
	}

	return nil
}
