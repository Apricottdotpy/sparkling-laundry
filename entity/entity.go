package entity

type Customers struct {
	Id   int
	Name string
	NoHp string
}

type Pegawai struct {
	Id   int
	Name string
}

type Layanan struct {
	Id          int
	Pelayanan string
	Harga       int
}

type Transaksi struct {
	Id               int
	IdPelanggan      int
	IdPegawai        int
	TanggalMasuk     string
	TanggalSelesai    string
	StatusPembayaran string
}

type DetailTransaksi struct {
	Id          int
	IdTransaksi int
	IdLayanan   int
	Quantity    int
	TotalHarga  int
}
