package codes

import (
	"google.golang.org/grpc/codes"
)

// StatusMessage represent string message for code
var StatusMessage = map[codes.Code]string{
	Success:              "Berhasil",
	SuccessCreated:       "Berhasil, Data Tersimpan",
	SuccessNoContent:     "Berhasil, Data tidak ditemukan",
	InvalidArgument:      "Parameter tidak valid",
	Unauthorized:         "Username atau password kamu salah",
	Forbidden:            "Akses tidak dibolehkan atau kamu tidak memiliki akses",
	NotFound:             "Data tidak ditemukan",
	Cancelled:            "Permintaan dibatalkan",
	RequestTimeout:       "Permintaan melebihi batas waktu",
	InactiveAccount:      "Akun tidak aktif",
	InvalidToken:         "Akses tidak valid karena token tidak cocok atau login sudah kadaluarsa",
	InvalidAPIKey:        "API key tidak valid",
	InvalidSession:       "Sesi tidak valid atau sudah berakhir",
	ResourceExhausted:    "Sudah mencapai batas limit",
	InvalidSubdomain:     "Nama Toko kamu salah",
	InactiveSubdomain:    "Toko belum diaktifkan",
	SuspendedSubdomain:   "Toko kamu diblokir, silahkan hubungi Qasir melalui email hello@qasir.id",
	InvalidTransaction:   "Data transaksi tidak valid atau tidak sesuai",
	DuplicateTransaction: "Data transaksi duplikat",
	InternalError:        "Error dari server",
}
