package main

import "fmt"

type barang struct {
	nama      string
	harga     int
	stok      int
	kategori  string
	penjualan int
	hargaJual int
}

type penjualan struct {
	hargaAwal int
	penjualan int
	nama      string
	harga     int
	tHarga    int
	kategori  string
}

type arrBarang [NMAX]barang

type arrPenjualan [NMAX]penjualan

const NMAX int = 5

func transaksi(A *arrPenjualan, T *arrBarang, N *int, N1 *int) {
	var opsi, opsi1, opsiIndex, jmlBarang int

	UIbeliBarang_1()
	fmt.Scan(&opsi)

	for opsi != 2 {
		if *N != 0 {
			for opsi < 1 || opsi > 2 {
				fmt.Println()
				fmt.Println("Maaf pilihan", opsi, "tidak sesuai dengan ")
				fmt.Println("opsi apapun, silahkan memilih opsi lain")
				UIbeliBarang_1()
				fmt.Scan(&opsi)
			}
			if opsi == 1 {
				fmt.Println()
				fmt.Println("=====================  Daftar Barang  =====================")
				display(*T, *N)
				fmt.Println()
				fmt.Print("Masukkan no. barang : ")
				fmt.Scan(&opsiIndex)
				for opsiIndex < 1 || opsiIndex > *N {
					fmt.Println()
					fmt.Println("Maaf, nomor yang anda inputkan tidak ada")
					fmt.Println()
					fmt.Print("Masukkan no barang : ")
					fmt.Scan(&opsiIndex)
				}

				fmt.Print("Masukkan jumlah barang : ")
				fmt.Scan(&jmlBarang)

				fmt.Println()
				fmt.Println("Rincian Transaksi :")
				fmt.Println("Nama :", T[opsiIndex-1].nama)
				fmt.Println("Harga :", T[opsiIndex-1].hargaJual)
				fmt.Println("Kategori :", T[opsiIndex-1].kategori)
				fmt.Println("Total tagihan = Rp.", T[opsiIndex-1].hargaJual*jmlBarang)
				fmt.Println()
				fmt.Println("1. Beli     2. Batal")
				fmt.Print("Masukkan Pilihan : ")
				fmt.Scan(&opsi1)

				if opsi1 == 1 {
					if jmlBarang <= T[opsiIndex-1].stok {
						fmt.Println()
						fmt.Println("Transaksi Berhasil")
						T[opsiIndex-1].penjualan += jmlBarang
						T[opsiIndex-1].stok -= jmlBarang
						A[*N1].hargaAwal = T[opsiIndex-1].harga
						A[*N1].nama = T[opsiIndex-1].nama
						A[*N1].harga = T[opsiIndex-1].hargaJual
						A[*N1].penjualan += jmlBarang
						A[*N1].kategori = T[opsiIndex-1].kategori
						A[*N1].tHarga = A[*N1].penjualan * A[*N1].harga
						displayPenjualan(*A, *N1)
						*N1++
						UIbeliBarang_1()
						fmt.Scan(&opsi)

					} else {
						fmt.Println()
						fmt.Println("Transaksi Gagal")
						fmt.Println("Jumlah stok tidak mencukupi")
						fmt.Println("Silahkan mengurangi jumlah pembelian")
						UIbeliBarang_1()
						fmt.Scan(&opsi)
					}
				} else if opsi1 == 2 {
					fmt.Println()
					fmt.Println("Transaksi Gagal")
					UIbeliBarang_1()
					fmt.Scan(&opsi)
				}
			}

		} else {
			fmt.Println()
			fmt.Println("Tidak terdapat data apapun")
			fmt.Println("Silahkan menambah data")
			UIbeliBarang_1()
			fmt.Scan(&opsi)
		}

	}
}

func addBarang(T *arrBarang, N *int) {
	var opsi1, opsi2, harga, stok int
	var nama, kategori string

	UIaddBarang_1()
	fmt.Scan(&opsi1)
	for opsi1 != 2 {
		for opsi1 < 1 || opsi1 > 2 {
			fmt.Println()
			fmt.Println("Maaf pilihan", opsi1, "tidak sesuai dengan ")
			fmt.Println("opsi apapun, silahkan memilih opsi lain")
			UIaddBarang_1()
			fmt.Scan(&opsi1)
		}

		if opsi1 == 1 {
			if *N < NMAX {
				UIaddBarang_2()
				fmt.Print("Masukkan nama barang : ")
				fmt.Scan(&nama)
				fmt.Print("Masukkan harga barang : ")
				fmt.Scan(&harga)
				fmt.Print("Masukkan banyak stok barang : ")
				fmt.Scan(&stok)
				fmt.Print("Masukkan kategori barang : ")
				fmt.Scan(&kategori)
				fmt.Println()
				fmt.Println("1. Simpan     2. Batal")
				fmt.Print("Masukkan pilihan : ")
				fmt.Scan(&opsi2)

				if opsi2 < 0 || opsi2 > 2 {
					fmt.Println()
					fmt.Println("Data tidak tersimpan")
					fmt.Println("Maaf pilihan", opsi2, "tidak sesuai dengan ")
					fmt.Println("opsi apapun, silahkan memilih opsi lain")
					UIaddBarang_2()
					fmt.Print("Masukkan nama barang : ")
					fmt.Scan(&nama)
					fmt.Print("Masukkan harga barang : ")
					fmt.Scan(&harga)
					fmt.Print("Masukkan banyak stok barang : ")
					fmt.Scan(&stok)
					fmt.Print("Masukkan kategori barang : ")
					fmt.Scan(&kategori)
					fmt.Println()
					fmt.Println("1. Simpan     2. Batal")
					fmt.Print("Masukkan pilihan : ")
					fmt.Scan(&opsi2)
				}

				if opsi2 == 1 {
					if !isMember(*T, *N, nama) {
						fmt.Println()
						fmt.Println("Data berhasil disimpan")
						T[*N].nama = nama
						T[*N].harga = harga
						T[*N].hargaJual = harga + (harga * 10 / 100)
						T[*N].stok = stok
						T[*N].kategori = kategori
						UIaddBarang_1()
						fmt.Scan(&opsi1)
						*N++
					} else {
						fmt.Println()
						fmt.Println("Data tidak tersimpan dikarenkan")
						fmt.Println("data yang anda masukkan sudah ada")
						fmt.Println("silahkan masukkan data lain")
						fmt.Println()
						UIaddBarang_1()
						fmt.Scan(&opsi1)
					}
				} else if opsi2 == 2 {
					fmt.Println()
					fmt.Println("Data tidak tersimpan")
					UIaddBarang_1()
					fmt.Scan(&opsi1)
				}
			} else {
				fmt.Println()
				fmt.Println("Array penuh, tidak dapat menginput ")
				fmt.Println("lebih banyak data")
				UIaddBarang_1()
				fmt.Scan(&opsi1)
			}
		}
	}
}

func editBarang(T *arrBarang, N *int, A *arrPenjualan, N1 *int) {
	var opsi, opsiEdit, opsiPilihan, opsiIndex, i int
	var temp arrBarang

	UIeditBarang_1()
	fmt.Scan(&opsi)
	for opsi != 3 {
		for opsi < 1 || opsi > 3 {
			fmt.Println()
			fmt.Println("Maaf pilihan", opsi, "tidak sesuai dengan ")
			fmt.Println("opsi apapun, silahkan memilih opsi lain")
			UIeditBarang_1()
			fmt.Scan(&opsi)
		}

		if opsi == 1 {
			if *N != 0 {
				fmt.Println()
				fmt.Println()
				fmt.Println("=====================  Daftar Barang  =====================")
				display(*T, *N)
				fmt.Println()
				fmt.Print("Masukkan no barang yang anda edit : ")
				fmt.Scan(&i)
				for i < 1 || i > *N {
					fmt.Println()
					fmt.Println("Maaf, nomor yang anda inputkan tidak ada")
					fmt.Println()
					fmt.Print("Masukkan no barang yang ingin diedit : ")
					fmt.Scan(&i)
				}
				opsiIndex = i - 1
				if opsiIndex <= *N {
					UIeditBarang_2()
					fmt.Scan(&opsiEdit)

					for opsiEdit != 5 {

						for opsiEdit < 1 || opsiEdit > 5 {
							fmt.Println()
							fmt.Println("Maaf pilihan", opsiEdit, "tidak sesuai dengan ")
							fmt.Println("opsi apapun, silahkan memilih opsi lain")
							UIeditBarang_2()
							fmt.Scan(&opsiEdit)
						}

						if opsiEdit == 1 {
							fmt.Print("Nama barang : ")
							fmt.Scan(&temp[opsiIndex].nama)
							UIpilihanEdit()
							fmt.Scan(&opsiPilihan)

							if opsiPilihan == 1 {
								if !isMember(*T, *N, temp[opsiIndex].nama) {
									fmt.Println()
									fmt.Println("Data berhasil disimpan")
									T[opsiIndex].nama = temp[opsiIndex].nama
									UIeditBarang_2()
									fmt.Scan(&opsiEdit)
								} else {
									fmt.Println()
									fmt.Println("Data tidak tersimpan dikarenkan")
									fmt.Println("data yang anda masukkan sudah ada")
									fmt.Println("silahkan masukkan data lain")
									UIeditBarang_2()
									fmt.Scan(&opsiEdit)
								}
							} else if opsiPilihan == 2 {
								fmt.Println()
								fmt.Println("Data tidak tersimpan")
								UIeditBarang_2()
								fmt.Scan(&opsiEdit)
							}

						} else if opsiEdit == 2 {
							fmt.Print("Harga barang : ")
							fmt.Scan(&temp[opsiIndex].harga)
							UIpilihanEdit()
							fmt.Scan(&opsiPilihan)

							if opsiPilihan == 1 {
								fmt.Println()
								fmt.Println("Data berhasil disimpan")
								T[opsiIndex].harga = temp[opsiIndex].harga

								UIeditBarang_2()
								fmt.Scan(&opsiEdit)

							} else if opsiPilihan == 2 {
								fmt.Println()
								fmt.Println("Data tidak tersimpan")
								UIeditBarang_2()
								fmt.Scan(&opsiEdit)
							}
						} else if opsiEdit == 3 {
							fmt.Print("Stok barang : ")
							fmt.Scan(&temp[opsiIndex].stok)
							UIpilihanEdit()
							fmt.Scan(&opsiPilihan)

							if opsiPilihan == 1 {
								fmt.Println()
								fmt.Println("Data berhasil disimpan")
								T[opsiIndex].stok = temp[opsiIndex].stok

								UIeditBarang_2()
								fmt.Scan(&opsiEdit)

							} else if opsiPilihan == 2 {
								fmt.Println()
								fmt.Println("Data tidak tersimpan")
								UIeditBarang_2()
								fmt.Scan(&opsiEdit)
							}
						} else if opsiEdit == 4 {
							fmt.Print("Kategori barang : ")
							fmt.Scan(&temp[opsiIndex].kategori)
							UIpilihanEdit()
							fmt.Scan(&opsiPilihan)

							if opsiPilihan == 1 {
								fmt.Println()
								fmt.Println("Data berhasil disimpan")
								T[opsiIndex].kategori = temp[opsiIndex].kategori

								UIeditBarang_2()
								fmt.Scan(&opsiEdit)

							} else if opsiPilihan == 2 {
								fmt.Println()
								fmt.Println("Data tidak tersimpan")
								UIeditBarang_2()
								fmt.Scan(&opsiEdit)
							}
						}
					}
					UIeditBarang_1()
					fmt.Scan(&opsi)

				}
			} else {
				fmt.Println()
				fmt.Println("Tidak terdapat data apapun")
				fmt.Println("Silahkan menambah data")
				UIeditBarang_1()
				fmt.Scan(&opsi)
			}

		} else if opsi == 2 {
			if *N1 != 0 {
				fmt.Println()
				fmt.Println()
				fmt.Println("====================  Daftar Transaksi  ===================")
				displayPenjualan(*A, *N1)
				fmt.Println()
				fmt.Print("Masukkan no barang yang anda edit : ")
				fmt.Scan(&i)
				for i < 1 || i > *N1 {
					fmt.Println()
					fmt.Println("Maaf, nomor yang anda inputkan tidak ada")
					fmt.Println()
					fmt.Print("Masukkan no transaksi yang ingin diedit : ")
					fmt.Scan(&i)
				}
				opsiIndex = i - 1
				if opsiIndex <= *N1 {
					UIeditTransaksi_2()
					fmt.Scan(&opsiEdit)

					for opsiEdit != 3 {
						for opsiEdit < 1 || opsiEdit > 5 {
							fmt.Println()
							fmt.Println("Maaf pilihan", opsiEdit, "tidak sesuai dengan ")
							fmt.Println("opsi apapun, silahkan memilih opsi lain")
							UIeditTransaksi_2()
							fmt.Scan(&opsiEdit)
						}
						if opsiEdit == 1 {
							fmt.Print("Harga barang (bukan harga jual) : ")
							fmt.Scan(&temp[opsiIndex].harga)
							UIpilihanEdit()
							fmt.Scan(&opsiPilihan)

							if opsiPilihan == 1 {
								fmt.Println()
								fmt.Println("Data berhasil disimpan")
								A[opsiIndex].harga = temp[opsiIndex].harga + (temp[opsiIndex].harga*10)/100
								A[opsiIndex].tHarga = A[opsiIndex].penjualan * A[opsiIndex].harga
								UIeditTransaksi_2()
								fmt.Scan(&opsiEdit)
							} else if opsiPilihan == 2 {
								fmt.Println()
								fmt.Println("Data tidak tersimpan")
								UIeditTransaksi_2()
								fmt.Scan(&opsiEdit)
							}
						} else if opsiEdit == 2 {
							fmt.Print("Kategori barang : ")
							fmt.Scan(&temp[opsiIndex].kategori)
							UIpilihanEdit()
							fmt.Scan(&opsiPilihan)

							if opsiPilihan == 1 {
								fmt.Println()
								fmt.Println("Data berhasil disimpan")
								A[opsiIndex].kategori = temp[opsiIndex].kategori

								UIeditTransaksi_2()
								fmt.Scan(&opsiEdit)

							} else if opsiPilihan == 2 {
								fmt.Println()
								fmt.Println("Data tidak tersimpan")
								UIeditTransaksi_2()
								fmt.Scan(&opsiEdit)
							}
						}

					}
					UIeditBarang_1()
					fmt.Scan(&opsi)
				}
			} else {
				fmt.Println()
				fmt.Println("Tidak terdapat data apapun")
				fmt.Println("Silahkan menambah data")
				UIeditBarang_1()
				fmt.Scan(&opsi)
			}
		}

	}

}

func lihatTransaksi(A *arrPenjualan, N1 *int, T arrBarang, N int) {
	var opsi, opsi1, i, j, modal, modalTerjual, pendapatan int
	var pilihan, pilihan1 string

	UIlihatPenjualan_1()
	fmt.Scan(&opsi)

	for opsi != 4 {
		for opsi < 1 || opsi > 4 {
			fmt.Println()
			fmt.Println("Maaf pilihan", opsi, "tidak sesuai dengan ")
			fmt.Println("opsi apapun, silahkan memilih opsi lain")
			UIlihatPenjualan_1()
			fmt.Scan(&opsi)
		}

		if opsi == 1 {
			if *N1 != 0 {
				fmt.Println()
				fmt.Println("=====================  Daftar Barang  =====================")
				displayPenjualan(*A, *N1)
				fmt.Println()
				UIlihatPenjualan_2()
				fmt.Scan(&opsi1)

				for opsi1 != 3 {
					for opsi1 < 1 || opsi1 > 3 {
						fmt.Println()
						fmt.Println("Maaf pilihan", opsi1, "tidak sesuai dengan ")
						fmt.Println("opsi apapun, silahkan memilih opsi lain")
						UIlihatPenjualan_2()
						fmt.Scan(&opsi1)
					}
					if opsi1 == 1 {
						UIurutPenjualan_1()
						fmt.Scan(&pilihan)
						fmt.Println("(terbesar/terkecil)")
						fmt.Print("Masukkan Pilihan : ")
						fmt.Scan(&pilihan1)
						mengurutkanTransaksi(A, *N1, pilihan, pilihan1)
						fmt.Println()
						fmt.Println("=====================  Daftar Barang  =====================")
						displayPenjualan(*A, *N1)
						UIlihatPenjualan_2()
						fmt.Scan(&opsi1)
					} else if opsi1 == 2 {
						fmt.Println()
						fmt.Println("=====================  Daftar Barang  =====================")
						mengurutkanTransaksi(A, *N1, "penjualan", "terbesar")
						displayPenjualanTOP5(*A, *N1)
						UIlihatPenjualan_2()
						fmt.Scan(&opsi1)
					}
				}
				UIlihatPenjualan_1()
				fmt.Scan(&opsi)
			} else {
				fmt.Println()
				fmt.Println("Tidak terdapat data apapun")
				fmt.Println("Silahkan menambah data")
				UIlihatPenjualan_1()
				fmt.Scan(&opsi)
			}
		} else if opsi == 2 {
			if N != 0 {
				fmt.Println()
				fmt.Println("====================  Modal Penjualan  ====================")
				for i = 0; i < N; i++ {
					modal += T[i].stok * T[i].harga
				}
				for j = 0; j < *N1; j++ {
					modalTerjual += A[j].penjualan * A[j].hargaAwal
				}
				fmt.Println()
				fmt.Println("Data modal didapat dari data seluruh barang")
				fmt.Println("Jumlah modal : ", modal+modalTerjual)
				modal = 0
				modalTerjual = 0

				UIlihatPenjualan_1()
				fmt.Scan(&opsi)
			} else {
				fmt.Println()
				fmt.Println("Tidak terdapat data apapun")
				fmt.Println("Silahkan menambah data")
				UIlihatPenjualan_1()
				fmt.Scan(&opsi)
			}
		} else if opsi == 3 {
			if *N1 != 0 {
				fmt.Println()
				fmt.Println("==================  Transaksi Penjualan  ==================")
				for i = 0; i < N; i++ {
					modal += T[i].stok * T[i].harga
				}
				for j = 0; j < *N1; j++ {
					modalTerjual += A[j].penjualan * A[j].hargaAwal
					pendapatan += A[j].penjualan * A[j].harga
				}
				fmt.Println()
				fmt.Println("Data pendapatan didapat dari data seluruh transaksi")
				fmt.Println("Jumlah pendapatan kotor : ", pendapatan)
				fmt.Println("Jumlah pendapatan bersih : ", pendapatan-(modal+modalTerjual))
				pendapatan = 0
				modal = 0
				modalTerjual = 0
				UIlihatPenjualan_1()
				fmt.Scan(&opsi)
			} else {
				fmt.Println()
				fmt.Println("Tidak terdapat data apapun")
				fmt.Println("Silahkan menambah data")
				UIlihatPenjualan_1()
				fmt.Scan(&opsi)
			}
		}
	}

}

func lihatBarang(T *arrBarang, N *int) {
	var opsi, opsi1, idx int
	var pilihan, pilihan1, nama string

	UIlihatBarang_1()
	fmt.Scan(&opsi)

	for opsi != 2 {
		for opsi < 1 || opsi > 2 {
			fmt.Println()
			fmt.Println("Maaf pilihan", opsi, "tidak sesuai dengan ")
			fmt.Println("opsi apapun, silahkan memilih opsi lain")
			UIlihatBarang_1()
			fmt.Scan(&opsi)
		}

		if opsi == 1 {
			if *N != 0 {
				fmt.Println()
				fmt.Println("=====================  Daftar Barang  =====================")
				display(*T, *N)
				fmt.Println()
				UIlihatBarang_2()
				fmt.Scan(&opsi1)

				for opsi1 != 3 {
					for opsi1 < 1 || opsi1 > 3 {
						fmt.Println()
						fmt.Println("Maaf pilihan", opsi1, "tidak sesuai dengan ")
						fmt.Println("opsi apapun, silahkan memilih opsi lain")
						UIlihatBarang_2()
						fmt.Scan(&opsi1)
					}
					if opsi1 == 1 {
						UIurutBarang_1()
						fmt.Scan(&pilihan)
						fmt.Println("(terbesar/terkecil)")
						fmt.Print("Masukkan Pilihan : ")
						fmt.Scan(&pilihan1)
						mengurutkan(T, *N, pilihan, pilihan1)
						fmt.Println()
						fmt.Println("=====================  Daftar Barang  =====================")
						display(*T, *N)
						UIlihatBarang_2()
						fmt.Scan(&opsi1)
					} else if opsi1 == 2 {
						mengurutkanSelection(T, N)
						fmt.Println()
						fmt.Println("=====================  Daftar Barang  =====================")
						display(*T, *N)
						fmt.Println()
						fmt.Print("Cari Nama Barang : ")
						fmt.Scan(&nama)
						if isMemberIndexBiner(*T, *N, nama) != -1 {
							idx = isMemberIndexBiner(*T, *N, nama)

							fmt.Println()
							fmt.Println("Rincian Barang : ")
							fmt.Println("Nama : ", T[idx].nama)
							fmt.Println("Harga : ", T[idx].harga)
							fmt.Println("Stok : ", T[idx].stok)
							fmt.Println("Kategori : ", T[idx].kategori)
							fmt.Println("Penjualan : ", T[idx].penjualan)
							UIlihatBarang_2()
							fmt.Scan(&opsi1)

						} else {
							fmt.Println()
							fmt.Println("Maaf, data tersebut tidak ada")
							fmt.Println("Silahkan cari data yang lain")
							UIlihatBarang_2()
							fmt.Scan(&opsi1)
						}
					}
				}
				UIlihatBarang_1()
				fmt.Scan(&opsi)
			} else {
				fmt.Println()
				fmt.Println("Tidak terdapat data apapun")
				fmt.Println("Silahkan menambah data")
				UIlihatBarang_1()
				fmt.Scan(&opsi)
			}
		}
	}

}

func hapusData(T *arrBarang, A *arrPenjualan, N *int, N1 *int) {
	var opsi, opsi1, opsiIndexBarang, opsiIndexPenjualan int

	UIhapusData_1()
	fmt.Scan(&opsi)

	for opsi != 3 {
		for opsi < 0 || opsi > 3 {
			fmt.Println()
			fmt.Println("Maaf pilihan", opsi, "tidak sesuai dengan ")
			fmt.Println("opsi apapun, silahkan memilih opsi lain")
			UIhapusData_1()
			fmt.Scan(&opsi)
		}
		if opsi == 1 {
			if *N != 0 {
				fmt.Println()
				fmt.Println("=====================  Daftar Barang  =====================")
				display(*T, *N)
				fmt.Println()
				fmt.Print("Masukkan no barang yang ingin dihapus : ")
				fmt.Scan(&opsiIndexBarang)
				for opsiIndexBarang < 1 || opsiIndexBarang > *N {
					fmt.Println()
					fmt.Println("Maaf, nomor yang anda inputkan tidak ada")
					fmt.Println()
					fmt.Print("Masukkan no barang yang ingin dihapus : ")
					fmt.Scan(&opsiIndexBarang)
				}
				fmt.Println("1. Hapus     2. Batal")
				fmt.Print("Masukkan Pilihan : ")
				fmt.Scan(&opsi1)

				if opsi1 == 1 {
					fmt.Println()
					fmt.Println("Data Berhasil dihapus")
					hapusDataBarang(T, N, opsiIndexBarang)
					UIhapusData_1()
					fmt.Scan(&opsi)
				} else if opsi1 == 2 {
					fmt.Println()
					fmt.Println("Data Tidak Terhapus")
					UIhapusData_1()
					fmt.Scan(&opsi)
				}

			} else {
				fmt.Println()
				fmt.Println("Tidak terdapat data apapun")
				fmt.Println("Silahkan menambah data")
				UIhapusData_1()
				fmt.Scan(&opsi)
			}
		} else if opsi == 2 {
			if *N1 != 0 {
				fmt.Println()
				fmt.Println("====================  Daftar Transaksi  ===================")
				displayPenjualan(*A, *N1)
				fmt.Println()
				fmt.Print("Masukkan no barang yang ingin dihapus : ")
				fmt.Scan(&opsiIndexPenjualan)
				for opsiIndexPenjualan < 1 || opsiIndexPenjualan > *N1 {
					fmt.Println()
					fmt.Println("Maaf, nomor yang anda inputkan tidak ada")
					fmt.Println()
					fmt.Print("Masukkan no barang yang ingin dihapus : ")
					fmt.Scan(&opsiIndexPenjualan)
				}
				fmt.Println("1. Hapus     2. Batal")
				fmt.Print("Masukkan Pilihan : ")
				fmt.Scan(&opsi1)

				if opsi1 == 1 {
					fmt.Println()
					fmt.Println("Data Berhasil dihapus")
					hapusDataPenjualan(A, N1, opsiIndexPenjualan)
					UIhapusData_1()
					fmt.Scan(&opsi)
				} else if opsi1 == 2 {
					fmt.Println()
					fmt.Println("Data Tidak Terhapus")
					UIhapusData_1()
					fmt.Scan(&opsi)
				}

			} else {
				fmt.Println()
				fmt.Println("Tidak terdapat data apapun")
				fmt.Println("Silahkan menambah data")
				UIhapusData_1()
				fmt.Scan(&opsi)
			}
		}

	}
}

func hapusDataBarang(T *arrBarang, N *int, index int) {

	for i := index - 1; i < *N-1; i++ {
		T[i] = T[i+1]
	}
	*N--
}

func hapusDataPenjualan(A *arrPenjualan, N1 *int, index int) {
	for i := index - 1; i < *N1-1; i++ {
		A[i] = A[i+1]
	}
	*N1--
}

func isMember(T arrBarang, N int, s string) bool {
	for i := 0; i < N; i++ {
		if T[i].nama == s {
			return true
		}
	}
	return false

}

func isMemberIndex(T arrBarang, N int, s string) int {
	var idx int = -1

	for i := 0; i < N; i++ {
		if T[i].nama == s {
			idx = i
		}
	}
	return idx

}

func isMemberIndexBiner(T arrBarang, N int, s string) int {
	var found int = -1
	var med int
	var kr int = 0
	var kn int = N - 1

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if s > T[med].nama {
			kn = med - 1
		} else if s < T[med].nama {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found

}

func menu(A *arrPenjualan, T *arrBarang, N *int, N1 *int) {
	var opsi int

	header()
	fmt.Scan(&opsi)
	for opsi != 7 {
		for opsi < 1 || opsi > 7 {
			fmt.Println()
			fmt.Println("Maaf pilihan", opsi, "tidak sesuai dengan ")
			fmt.Println("opsi apapun, silahkan memilih opsi lain")
			header()
			fmt.Scan(&opsi)
		}
		if opsi == 1 {
			transaksi(A, T, N, N1)
		} else if opsi == 2 {
			addBarang(T, N)

		} else if opsi == 3 {
			editBarang(T, N, A, N1)
		} else if opsi == 4 {
			lihatTransaksi(A, N1, *T, *N)
		} else if opsi == 5 {
			lihatBarang(T, N)
		} else if opsi == 6 {
			hapusData(T, A, N, N1)
		}
		header()
		fmt.Scan(&opsi)

	}
	fmt.Println()
	fmt.Println("Terimakasih telah menggunakan layanan kami :) ")

}

func mengurutkan(T *arrBarang, N int, s string, s1 string) {
	var i, j int
	var temp barang

	if s == "harga" {
		if s1 == "terbesar" {
			i = 1
			for i <= N-1 {
				j = i
				temp = T[j]
				for j > 0 && temp.harga > T[j-1].harga {
					T[j] = T[j-1]
					j--
				}
				T[j] = temp
				i++
			}
		} else if s1 == "terkecil" {
			i = 1
			for i <= N-1 {
				j = i
				temp = T[j]
				for j > 0 && temp.harga < T[j-1].harga {
					T[j] = T[j-1]
					j--
				}
				T[j] = temp
				i++
			}
		}

	} else if s == "stok" {
		if s1 == "terbesar" {
			i = 1
			for i <= N-1 {
				j = i
				temp = T[j]
				for j > 0 && temp.stok > T[j-1].stok {
					T[j] = T[j-1]
					j--
				}
				T[j] = temp
				i++
			}
		} else if s1 == "terkecil" {
			i = 1
			for i <= N-1 {
				j = i
				temp = T[j]
				for j > 0 && temp.stok < T[j-1].stok {
					T[j] = T[j-1]
					j--
				}
				T[j] = temp
				i++
			}
		}
	} else if s == "nama" {
		if s1 == "terbesar" {
			i = 1
			for i <= N-1 {
				j = i
				temp = T[j]
				for j > 0 && temp.nama > T[j-1].nama {
					T[j] = T[j-1]
					j--
				}
				T[j] = temp
				i++
			}
		} else if s1 == "terkecil" {
			i = 1
			for i <= N-1 {
				j = i
				temp = T[j]
				for j > 0 && temp.nama < T[j-1].nama {
					T[j] = T[j-1]
					j--
				}
				T[j] = temp
				i++
			}
		}
	} else if s == "kategori" {
		if s1 == "terbesar" {
			i = 1
			for i <= N-1 {
				j = i
				temp = T[j]
				for j > 0 && temp.kategori > T[j-1].kategori {
					T[j] = T[j-1]
					j--
				}
				T[j] = temp
				i++
			}
		} else if s1 == "terkecil" {
			i = 1
			for i <= N-1 {
				j = i
				temp = T[j]
				for j > 0 && temp.kategori < T[j-1].kategori {
					T[j] = T[j-1]
					j--
				}
				T[j] = temp
				i++
			}
		}
	} else if s == "penjualan" {
		if s1 == "terbesar" {
			i = 1
			for i <= N-1 {
				j = i
				temp = T[j]
				for j > 0 && temp.penjualan > T[j-1].penjualan {
					T[j] = T[j-1]
					j--
				}
				T[j] = temp
				i++
			}
		} else if s1 == "terkecil" {
			i = 1
			for i <= N-1 {
				j = i
				temp = T[j]
				for j > 0 && temp.penjualan < T[j-1].penjualan {
					T[j] = T[j-1]
					j--
				}
				T[j] = temp
				i++
			}
		}
	}

}

func mengurutkanTransaksi(A *arrPenjualan, N1 int, s string, s1 string) {
	var i, j int
	var temp penjualan

	if s == "nama" {
		if s1 == "terbesar" {
			i = 1
			for i <= N1-1 {
				j = i
				temp = A[j]
				for j > 0 && temp.nama > A[j-1].nama {
					A[j] = A[j-1]
					j--
				}
				A[j] = temp
				i++
			}
		} else if s1 == "terkecil" {
			i = 1
			for i <= N1-1 {
				j = i
				temp = A[j]
				for j > 0 && temp.nama < A[j-1].nama {
					A[j] = A[j-1]
					j--
				}
				A[j] = temp
				i++
			}
		}
	} else if s == "kategori" {
		if s1 == "terbesar" {
			i = 1
			for i <= N1-1 {
				j = i
				temp = A[j]
				for j > 0 && temp.kategori > A[j-1].kategori {
					A[j] = A[j-1]
					j--
				}
				A[j] = temp
				i++
			}
		} else if s1 == "terkecil" {
			i = 1
			for i <= N1-1 {
				j = i
				temp = A[j]
				for j > 0 && temp.kategori < A[j-1].kategori {
					A[j] = A[j-1]
					j--
				}
				A[j] = temp
				i++
			}
		}
	} else if s == "totalharga" {
		if s1 == "terbesar" {
			i = 1
			for i <= N1-1 {
				j = i
				temp = A[j]
				for j > 0 && temp.harga > A[j-1].harga {
					A[j] = A[j-1]
					j--
				}
				A[j] = temp
				i++
			}
		} else if s1 == "terkecil" {
			i = 1
			for i <= N1-1 {
				j = i
				temp = A[j]
				for j > 0 && temp.harga < A[j-1].harga {
					A[j] = A[j-1]
					j--
				}
				A[j] = temp
				i++
			}
		}
	} else if s == "penjualan" {
		if s1 == "terbesar" {
			i = 1
			for i <= N1-1 {
				j = i
				temp = A[j]
				for j > 0 && temp.penjualan > A[j-1].penjualan {
					A[j] = A[j-1]
					j--
				}
				A[j] = temp
				i++
			}
		} else if s1 == "terkecil" {
			i = 1
			for i <= N1-1 {
				j = i
				temp = A[j]
				for j > 0 && temp.penjualan < A[j-1].penjualan {
					A[j] = A[j-1]
					j--
				}
				A[j] = temp
				i++
			}
		}
	}
}

func mengurutkanSelection(T *arrBarang, N *int) {
	var i, j, idx_min int
	var temp barang

	i = 1
	for i <= *N-1 {
		idx_min = i - 1
		j = i
		for j < *N {
			if T[idx_min].nama < T[j].nama {
				idx_min = j
			}
			j++
		}
		temp = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = temp
		i++
	}
}

func displayPenjualanTOP5(A arrPenjualan, N1 int) {
	for i := 0; i < 5; i++ {
		if A[i].penjualan != 0 {
			fmt.Println(i+1, ".", "Nama :", A[i].nama, ",Harga :", A[i].harga, ",kategori :", A[i].kategori, ",Penjualan :", A[i].penjualan, ",Total Harga :", A[i].tHarga)
		}
	}
}

func display(T arrBarang, N int) {
	for i := 0; i < N; i++ {
		fmt.Println(i+1, ".", "Nama :", T[i].nama, ",Harga :", T[i].harga, ",Stok :", T[i].stok, ",kategori :", T[i].kategori, ",Penjualan :", T[i].penjualan)
	}
}

func displayPenjualan(A arrPenjualan, N1 int) {
	for i := 0; i < N1; i++ {
		fmt.Println(i+1, ".", "Nama :", A[i].nama, ",Harga :", A[i].harga, ",kategori :", A[i].kategori, ",Penjualan :", A[i].penjualan, ",Total Harga :", A[i].tHarga)
	}
}

func main() {
	var arrayBarang arrBarang
	var arrayPenjualan arrPenjualan
	var jumlah, jumlah1 int

	menu(&arrayPenjualan, &arrayBarang, &jumlah, &jumlah1)

}

func UIaddBarang_1() {
	fmt.Println()
	fmt.Println("===================  Menu Tambah Data  ====================")
	fmt.Println("1. Tambah data barang")
	fmt.Println("2. Keluar")
	fmt.Print("Masukkan pilihan : ")
}

func UIeditBarang_1() {
	fmt.Println()
	fmt.Println("====================  Menu Edit Data  =====================")
	fmt.Println("1. Edit data barang")
	fmt.Println("2. Edit data transaksi")
	fmt.Println("3. Keluar")
	fmt.Print("Masukkan pilihan : ")
}

func UIaddBarang_2() {
	fmt.Println()
	fmt.Println("====================  Masukkan Data   =====================")
}

func UIeditBarang_2() {
	fmt.Println()
	fmt.Println("==================  Masukkan Data Baru  ===================")
	fmt.Println("1. Nama barang")
	fmt.Println("2. Harga barang")
	fmt.Println("3. Stok barang")
	fmt.Println("4. Kategori barang")
	fmt.Println("5. Keluar")
	fmt.Print("Masukkan pilihan : ")
}

func UIeditTransaksi_2() {
	fmt.Println()
	fmt.Println("=============  Masukkan Data Transaksi Baru  ===============")
	fmt.Println("1. Harga barang")
	fmt.Println("2. Kategori barang")
	fmt.Println("3. Keluar")
	fmt.Print("Masukkan pilihan : ")
}

func UIlihatBarang_1() {
	fmt.Println()
	fmt.Println("===================  Menu Lihat Barang  ===================")
	fmt.Println("1. Lihat Barang")
	fmt.Println("2. Keluar")
	fmt.Print("Masukkan pilihan : ")
}

func UIlihatBarang_2() {
	fmt.Println()
	fmt.Println("1. Urutkan")
	fmt.Println("2. Cari Barang")
	fmt.Println("3. Keluar")
	fmt.Print("Masukkan Pilihan : ")

}

func UIlihatPenjualan_1() {
	fmt.Println()
	fmt.Println("=================  Menu Lihat Transaksi  ==================")
	fmt.Println("1. Lihat Transaksi")
	fmt.Println("2. Modal Penjualan")
	fmt.Println("3. Pendapatan")
	fmt.Println("4. Keluar")
	fmt.Print("Masukkan pilihan : ")
}

func UIlihatPenjualan_2() {
	fmt.Println()
	fmt.Println("1. Urutkan")
	fmt.Println("2. TOP 5 Penjualan")
	fmt.Println("3. Keluar")
	fmt.Print("Masukkan Pilihan : ")

}

func UIhapusData_1() {
	fmt.Println()
	fmt.Println("===================  Menu Hapus Barang  ===================")
	fmt.Println("1. Hapus Data Barang")
	fmt.Println("2. Hapus Data Penjualan")
	fmt.Println("3. Keluar")
	fmt.Print("Masukkan pilihan : ")
}

func UIurutBarang_1() {
	fmt.Println()
	fmt.Println("====================   Urutkan Barang  ====================")
	fmt.Println("Diurutkan berdasarkan : ")
	fmt.Println("(nama/harga/stok/kategori/penjualan)")
	fmt.Print("Masukkan Pilihan : ")
}

func UIurutPenjualan_1() {
	fmt.Println()
	fmt.Println("===================   Urutkan Transaksi  ==================")
	fmt.Println("Diurutkan berdasarkan : ")
	fmt.Println("(nama/totalharga/kategori/penjualan)")
	fmt.Print("Masukkan Pilihan : ")
}

func UIbeliBarang_1() {
	fmt.Println()
	fmt.Println("===================  Menu Beli Barang  ====================")
	fmt.Println("1. Beli barang")
	fmt.Println("2. Keluar")
	fmt.Print("Masukkan pilihan : ")
}

func UIpilihanEdit() {
	fmt.Println()
	fmt.Println("1. Simpan     2. Batal")
	fmt.Print("Masukkan pilihan : ")
}

func header() {
	fmt.Println()
	fmt.Println()
	fmt.Println("===========================================================")
	fmt.Println("|                                                         |")
	fmt.Println("|                  Toko HahaHihiCumlaude                  |")
	fmt.Println("|                                                         |")
	fmt.Println("===========================================================")
	fmt.Println("1. Beli barang")
	fmt.Println("2. Tambah barang")
	fmt.Println("3. Edit Data")
	fmt.Println("4. Daftar transaksi")
	fmt.Println("5. Daftar barang")
	fmt.Println("6. Hapus Data")
	fmt.Println("7. Exit")
	fmt.Println("===========================================================")

	fmt.Print("Pilihlah salah satu opsi diatas : ")

}
