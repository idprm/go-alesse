package database

import (
	"time"

	"github.com/idprm/go-alesse/src/pkg/model"
)

var configs = []model.Config{
	{
		Name:  "AUTO_MESSAGE_SENDBIRD",
		Value: "Hi, Saya @v1 silahkan jelaskan keluhan kamu",
	},
	{
		Name:  "NOTIF_MESSAGE_DOCTOR",
		Value: "Hi *@v1*, User *@v2* menunggu konfirmasi untuk konsultasi online. Klik disini untuk memulai chat @v3",
	},
	{
		Name:  "NOTIF_MESSAGE_USER",
		Value: "Hi *@v1* pembayaran Anda sudah terkonfirmasi. Untuk Chat dengan *@v2* klik disini @v3 (Add to contact agar link bisa diklik)",
	},
	{
		Name:  "NOTIF_OTP_USER",
		Value: "Berikut adalah kode OTP kamu : *@v1* untuk mulai konsultasi dokter di yellowclinic.com",
	},
	{
		Name:  "PRICE",
		Value: "20000",
	},
	{
		Name:  "DISCOUNT",
		Value: "10000",
	},
	{
		Name:  "PAGE_FINISH",
		Value: "<p>Hi @v1 pastikan Anda sudah membayar Rp @v2 untuk berkonsultasi dengan @v3</p><p>(Abaikan apabila sudah membayar)</p><p>Kami akan mengirimkan Whatsapp notifikasi apabila pembayaran sudah terkonfirmasi</p>",
	},
	{
		Name:  "PAGE_UNFINISH",
		Value: "-",
	},
	{
		Name:  "PAGE_ERROR",
		Value: "-",
	},
}

var doctors = []model.Doctor{
	{
		Username:             "dr-ernita",
		Name:                 "dr. Ernita Rosyanti Dewi",
		Photo:                "dr-ernita.png",
		Type:                 "Dokter Umum",
		Number:               "STR 3121100220145699",
		Experience:           5,
		GraduatedFrom:        "Universitas Indonesia, 2007",
		ConsultationSchedule: "06.00 - 23.00 WIB",
		PlacePractice:        "Jakarta Timur, DKI Jakarta",
		Phone:                "6281776736076",
		Start:                time.Now(),
		End:                  time.Now(),
	},
	{
		Username:             "dr-ayu",
		Name:                 "dr. Ayu A. Istiana",
		Photo:                "dr-ayu.png",
		Type:                 "Dokter Umum",
		Number:               "STR 3121100220145544",
		Experience:           7,
		GraduatedFrom:        "Universitas Yarsi, 2013",
		ConsultationSchedule: "06.00 - 23.00 WIB",
		PlacePractice:        "Bogor, Jawa Barat",
		Phone:                "6281212480644",
		Start:                time.Now(),
		End:                  time.Now(),
	},
	{
		Username:             "dr-peter",
		Name:                 "dr. Peter Fernando",
		Photo:                "dr-peter.png",
		Type:                 "Dokter Umum",
		Number:               "STR 6111100120221435",
		Experience:           3,
		GraduatedFrom:        "Universitas Tanjungpura, 2019",
		ConsultationSchedule: "06.00 - 23.00 WIB",
		PlacePractice:        "Ngabang, Kalimantan Timur",
		Phone:                "6281776736076",
		Start:                time.Now(),
		End:                  time.Now(),
	},
}

var clinics = []model.Clinic{
	{
		Name:     "Klinik Cepat Sehat Indonesia",
		Photo:    "yellowclinic.png",
		Address:  "Jl. Peternakan No.13, RT.5/RW.1, Kp. Tengah, Kec. Kramat jati, Jakarta, Daerah Khusus Ibukota Jakarta 13540",
		Phone:    "6281281881802",
		IsActive: true,
	},
	{
		Name:     "Klinik Cepat Sehat Indonesia",
		Photo:    "yellowclinic.png",
		Address:  "Jl. Peternakan No.13, RT.5/RW.1, Kp. Tengah, Kec. Kramat jati, Jakarta, Daerah Khusus Ibukota Jakarta 13540",
		Phone:    "6281281881802",
		IsActive: true,
	},
	{
		Name:     "Klinik Cepat Sehat Indonesia",
		Photo:    "yellowclinic.png",
		Address:  "Jl. Peternakan No.13, RT.5/RW.1, Kp. Tengah, Kec. Kramat jati, Jakarta, Daerah Khusus Ibukota Jakarta 13540",
		Phone:    "6281281881802",
		IsActive: true,
	},
}
