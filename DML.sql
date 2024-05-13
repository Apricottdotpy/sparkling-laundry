
INSERT INTO mst_pelanggan (nama_pelanggan, nomor_hp)
VALUES 
    ('Budi', '081234567890'),
    ('Ani', '081298765432'),
    ('Dewi', '082112345678'),
    ('Joko', '082198765432'),
    ('Tono', '085612345678'),
    ('Mira', '085698765432'),
    ('Rudi', '087712345678'),
    ('Linda', '087798765432'),
    ('Yanto', '089612345678'),
    ('Rini', '089698765432');

INSERT INTO mst_pegawai (nama_pegawai)
VALUES 
    ('Mirna'),
    ('Devi'),
    ('Anton');

INSERT INTO layanan (pelayanan, satuan, harga)
VALUES 
    ('Laundry Pakaian', 'KG', 5000),
    ('Laundry dan Setrika Pakaian', 'KG', 7000),
    ('Laundry Bedcover', 'Buah', 50000),
    ('Laundry Boneka', 'Buah', 25000);

INSERT INTO transaksi (id_pelanggan, id_pegawai, tanggal_masuk, tanggal_selesai, status_pembayaran)
VALUES 
    (1, 1, '2024-03-25', '2024-03-27', 'Lunas'),
    (2, 2, '2024-03-26', '2024-03-28', 'Lunas'),
    (3, 3, '2024-03-27', '2024-03-29', 'Belum Lunas');

INSERT INTO detail_transaksi (id_transaksi, id_layanan, quantity)
VALUES 
    (1, 1, 1),
    (1, 2, 2),
    (1, 3, 3),
    (2, 1, 2),
    (2, 2, 3),
    (2, 4, 2),
    (3, 2, 1),
    (3, 3, 1);