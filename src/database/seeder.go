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
		Value: "Hello pasien *@patient*, Apabila ada pertanyaan silakan hubungi nomor ini *@phone*",
	},
	{
		Name:  "NOTIF_MESSAGE_SPECIALIST",
		Value: "Hi *@v1*, Dokter Umum *@v2* menunggu konfirmasi untuk konsultasi online. Klik disini untuk memulai chat @v3",
	},
	{
		Name:  "NOTIF_OTP_USER",
		Value: "Berikut adalah kode OTP kamu : *@v1* untuk mulai konsultasi dokter di dottoro-ta.com",
	},
	{
		Name:  "NOTIF_OTP_ADMIN",
		Value: "Berikut adalah kode OTP kamu : *@v1*",
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
	//** SEED DATA NOTIF WA */
	{
		Name:  "NOTIF_DOCTOR_TO_PHARMACY",
		Value: "Hello Admin Farmasi *@health_center* terdapat pengajuan resep obat dari *@doctor* untuk pasien *@patient* Cek disini *@link*",
	},
	{
		Name:  "NOTIF_PHARMACY_TO_COURIER",
		Value: "Hello Kurir *@courier*, terdapat pemintaan pengantaran obat dari Farmasi *@pharmacy* untuk pasien *@patient*. Cek disini *@link*",
	},
	{
		Name:  "NOTIF_COURIER_TO_PHARMACY",
		Value: "Hello Admin Farmasi *@health_center*, Kurir *@courier* sudah menyelesaikan pengantaran obat ke pasien *@patient*",
	},
	{
		Name:  "NOTIF_PHARMACY_TO_PATIENT",
		Value: "Hello pasien *@patient* obat Anda sedang disiapkan oleh Farmasi *@health_center*",
	},
	{
		Name:  "NOTIF_COURIER_TO_PATIENT",
		Value: "Hello pasien *@patient* obat Anda sedang diantarkan oleh Kurir *@health_center*",
	},
	{
		Name:  "NOTIF_HOMECARE_TO_PATIENT_PROGRESS",
		Value: "Hello pasien *@patient*, tim Homecare *@health_center* akan datang kerumah Anda dalam waktu 15 menit - 1 jam. Apabila ada pertanyaan silakan hubungi nomor ini *@phone*",
	},
	{
		Name:  "NOTIF_HOMECARE_TO_PATIENT_DONE",
		Value: "Hello pasien *@patient*, layanan homecare dari tim Homecare *@health_center* sudah selesai dilakukan. Semoga Anda lekas sembuh. *@link*",
	},
	{
		Name:  "NOTIF_DOCTOR_TO_HOMECARE",
		Value: "Hello Admin Homecare *@health_center*, terdapat permintaan layanan homecare dari *@doctor* untuk pasien *@patient* Cek disini *@link*",
	},
	{
		Name:  "NOTIF_HOMECARE_TO_HEALTHOFFICE",
		Value: "Hello Admin Dinkes, tim homecare *@health_center* sudah menyelesaikan layanan homecare untuk pasien *@patient*",
	},
	{
		Name:  "NOTIF_FEEDBACK_TO_PATIENT",
		Value: "Hello pasien *@patient*, Semoga Anda lekas sembuh. Seberapa puaskah Anda dengan layanan puskesmas *@health_center* ? *@link*",
	},
}

var roles = []model.Role{
	{
		Name: "Pasien",
	},
	{
		Name: "Dokter Umum",
	},
	{
		Name: "Dokter Spesialis",
	},
	{
		Name: "Perawat",
	},
	{
		Name: "Farmasi",
	},
	{
		Name: "Kurir",
	},
}

var categories = []model.Category{
	{
		Code:     "chat",
		Name:     "Chat",
		IsActive: true,
	},
	{
		Code:     "pharmacy",
		Name:     "e-Resep",
		IsActive: true,
	},
	{
		Code:     "homecare",
		Name:     "Homecare",
		IsActive: true,
	},
}

var statuses = []model.Status{
	{
		Name:        "OTP_TO_USER",
		ValueSystem: "Kirim kode OTP",
		ValueUser:   "Mengirim kode OTP @otp kepada pasien @patient",
		ValueNotif:  "Berikut adalah kode OTP kamu : *@otp* untuk mulai konsultasi dokter di @link/auth/verify, Simpan Nomor ke Kontak agar link bisa di Klik",
		ValuePush:   "-",
	},
	{
		Name:        "OTP_TO_ADMIN",
		ValueSystem: "Kirim kode OTP",
		ValueUser:   "Mengirim kode OTP @otp",
		ValueNotif:  "Berikut adalah kode OTP kamu : *@otp*",
		ValuePush:   "-",
	},
	{
		Name:        "MESSAGE_TO_DOCTOR",
		ValueSystem: "Request chat",
		ValueUser:   "Pasien @patient mengajukan Konsultasi Dokter @doctor",
		ValueNotif:  "Hi *@doctor*, User *@patient* menunggu konfirmasi untuk konsultasi online. Klik disini untuk memulai chat *@link*, Simpan Nomor ke Kontak agar link bisa di Klik",
		ValuePush:   "Hi @doctor, User @patient menunggu konfirmasi untuk konsultasi online.",
	},
	{
		Name:        "MESSAGE_TO_SPECIALIST",
		ValueSystem: "Request chat specialist",
		ValueUser:   "Dokter umum @doctor mengajukan Konsultasi dengan Dokter spesialis @specialist",
		ValueNotif:  "Hi *@specialist*, Dokter Umum *@doctor* menunggu konfirmasi untuk konsultasi online. Klik disini untuk memulai chat *@link*, Simpan Nomor ke Kontak agar link bisa di Klik",
		ValuePush:   "Hi @specialist, Dokter Umum @doctor menunggu konfirmasi untuk konsultasi online.",
	},
	{
		Name:        "MESSAGE_TO_USER",
		ValueSystem: "Auto message pasien",
		ValueUser:   "Mengirim pesan kontak @phone kepada pasien @patient",
		ValueNotif:  "Hello pasien *@patient*, Apabila ada pertanyaan silakan hubungi nomor ini @phone",
		ValuePush:   "Hello pasien @patient, Apabila ada pertanyaan silakan hubungi nomor ini @phone",
	},
	{
		Name:        "DOCTOR_TO_PHARMACY",
		ValueSystem: "Request e-Resep",
		ValueUser:   "Dokter @doctor meresepkan e-Resep @number",
		ValueNotif:  "Hello Admin Farmasi *@health_center* terdapat pengajuan resep obat dari *@doctor* untuk pasien *@patient* Cek disini *@link*, Simpan Nomor ke Kontak agar link bisa di Klik",
		ValuePush:   "Hello Admin Farmasi @health_center terdapat pengajuan resep obat dari @doctor untuk pasien @patient",
	},
	{
		Name:        "PHARMACY_TO_COURIER",
		ValueSystem: "Request kurir",
		ValueUser:   "e-Resep @number telah dibuat. Menunggu Kurir @courier untuk mengambil obat.",
		ValueNotif:  "Hello Kurir *@courier*, terdapat pemintaan pengantaran obat dari Farmasi *@pharmacy* untuk pasien *@patient*. Cek disini *@link*, Simpan Nomor ke Kontak agar link bisa di Klik",
		ValuePush:   "Hello Kurir @courier, terdapat pemintaan pengantaran obat dari Farmasi @pharmacy untuk pasien @patient",
	},
	{
		Name:        "COURIER_TO_PHARMACY",
		ValueSystem: "Obat telah diterima Pasien",
		ValueUser:   "Pasien @patient telah menerima obat dari Kurir @courier.",
		ValueNotif:  "Hello Admin Farmasi *@health_center*, Kurir *@courier* sudah menyelesaikan pengantaran obat ke pasien *@patient*",
		ValuePush:   "Hello Admin Farmasi @health_center, Kurir @courier sudah menyelesaikan pengantaran obat ke pasien @patient",
	},
	{
		Name:        "PHARMACY_TO_PATIENT",
		ValueSystem: "e-Resep sedang dibuat",
		ValueUser:   "Farmasi Puskesmas @health_center telah memverifikasi dan sedang menyiapkan e-Resep @number.",
		ValueNotif:  "Hello pasien *@patient* obat Anda sedang disiapkan oleh Farmasi *@health_center*",
		ValuePush:   "Hello pasien @patient obat Anda sedang disiapkan oleh Farmasi @health_center",
	},
	{
		Name:        "COURIER_TO_PATIENT",
		ValueSystem: "Kurir menuju alamat pasien",
		ValueUser:   "Kurir telah mengambil obat di Farmasi Puskesmas @health_center dan menuju alamat Pasien @patient.",
		ValueNotif:  "Hello pasien *@patient* obat Anda sedang diantarkan oleh Kurir *@health_center*",
		ValuePush:   "Hello pasien @patient obat Anda sedang diantarkan oleh Kurir @health_center",
	},
	{
		Name:        "HOMECARE_TO_PATIENT_PROGRESS",
		ValueSystem: "Tim Homecare menuju alamat pasien",
		ValueUser:   "Tim Homecare Puskesmas @health_center telah memverifikasi Permintaan Homecare Dokter @doctor dan sedang menuju alamat Pasien @patient.",
		ValueNotif:  "Hello pasien *@patient*, tim Homecare *@health_center* akan datang kerumah Anda dalam waktu 15 menit - 1 jam. Apabila ada pertanyaan silakan hubungi nomor ini *@phone*",
		ValuePush:   "Hello pasien @patient, tim Homecare @health_center akan datang kerumah Anda dalam waktu 15 menit - 1 jam. Apabila ada pertanyaan silakan hubungi nomor ini @phone",
	},

	{
		Name:        "HOMECARE_TO_PATIENT_DONE",
		ValueSystem: "Pasien telah dikunjungi",
		ValueUser:   "Tim Homecare Puskesmas @health_center telah mengunjungi dan memeriksa Pasien @patient.",
		ValueNotif:  "Hello pasien *@patient*, layanan homecare dari tim Homecare *@health_center* sudah selesai dilakukan. Semoga Anda lekas sembuh. *@link*, Simpan Nomor ke Kontak agar link bisa di Klik",
		ValuePush:   "Hello pasien @patient, layanan homecare dari tim Homecare @health_center sudah selesai dilakukan. Semoga Anda lekas sembuh.",
	},
	{
		Name:        "HOMECARE_TO_HEALTHOFFICE",
		ValueSystem: "Homecare selesai",
		ValueUser:   "Aktivitas Homecare Puskesmas @health_center untuk pasien @patient telah selesai",
		ValueNotif:  "Hello Admin Dinkes, tim homecare *@health_center* sudah menyelesaikan layanan homecare untuk pasien *@patient*",
		ValuePush:   "Hello Admin Dinkes, tim homecare @health_center sudah menyelesaikan layanan homecare untuk pasien @patient",
	},
	{
		Name:        "DOCTOR_TO_HOMECARE",
		ValueSystem: "Request homecare",
		ValueUser:   "Dokter @doctor mengajukan Layanan Homecare untuk Pasien @patient",
		ValueNotif:  "Hello Admin Homecare *@health_center*, terdapat permintaan layanan homecare dari *@doctor* untuk pasien *@patient* Cek disini *@link*, Simpan Nomor ke Kontak agar link bisa di Klik",
		ValuePush:   "Hello Admin Homecare @health_center, terdapat permintaan layanan homecare dari @doctor untuk pasien @patient",
	},
	{
		Name:        "DOCTOR_TO_PATIENT_HOMECARE",
		ValueSystem: "Notif schedule homecare",
		ValueUser:   "Dokter @doctor telah menjadwalkan kunjungan Homecare pada tanggal @visit_at pukul @hour untuk Pasien @patient",
		ValueNotif:  "Dokter *@doctor* telah menjadwalkan kunjungan Homecare pada tanggal *@visit_at* pukul *@hour* untuk Pasien *@patient*",
		ValuePush:   "Dokter @doctor telah menjadwalkan kunjungan Homecare pada tanggal @visit_at pukul @hour untuk Pasien @patient",
	},
	{
		Name:        "FEEDBACK_TO_PATIENT",
		ValueSystem: "Feedback user",
		ValueUser:   "Mengirim link feedback kepada pasien @patient",
		ValueNotif:  "Hello pasien *@patient*, Semoga Anda lekas sembuh. Seberapa puaskah Anda dengan layanan puskesmas *@health_center* ? Klik Link ini untuk feedback *@link*, Simpan Nomor ke Kontak agar link bisa di Klik",
		ValuePush:   "Hello pasien @patient, Semoga Anda lekas sembuh. Seberapa puaskah Anda dengan layanan puskesmas @health_center",
	},
	{
		Name:        "CHAT_PROGRESS",
		ValueSystem: "Sedang chat",
		ValueUser:   "Pasien @patient memulai konsultasi dengan Dokter @doctor",
		ValueNotif:  "-",
		ValuePush:   "-",
	},
	{
		Name:        "CHAT_DONE",
		ValueSystem: "Chat selesai",
		ValueUser:   "Dokter @doctor mengakhiri konsultasi chat dengan Pasien @patient",
		ValueNotif:  "-",
		ValuePush:   "Dokter @doctor mengakhiri konsultasi chat dengan Pasien @patient",
	},
}

var healthcenters = []model.Healthcenter{
	{
		Name:     "ANDALAS",
		Code:     "1070781",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Sangir Lr.209 No.6, Kel. Melayu, Kec. Wajo",
		IsActive: true,
	},
	{
		Name:     "ANTANG",
		Code:     "1070795",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Antang Raya No.43, Kel. Antang, Kec. Manggala",
		IsActive: true,
	},
	{
		Name:     "ANTANG PERUMNAS",
		Code:     "1070796",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Jl. Lasuloro Kel. Manggalo, Kec. Manggala",
		IsActive: false,
	},
	{
		Name:     "ANTARA",
		Code:     "1070802",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Komp.BTN Antara Blok B No.6, Kec. Tamalanrea",
		IsActive: false,
	},
	{
		Name:     "BALLAPARANG",
		Code:     "1071285",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Nikel III No 1 Makassar",
		IsActive: false,
	},
	{
		Name:     "BANGKALA",
		Code:     "1070798",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Kec. Manggala",
		IsActive: false,
	},
	{
		Name:     "BARA-BARAYA",
		Code:     "1070776",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Jl. Abubakar Lambogo, Kel. Bara-Baraya, Kec. Makassar",
		IsActive: false,
	},
	{
		Name:     "BARANG LOMPO",
		Code:     "1070785",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Jl. Pulau Barraang Lompo Rw 1, Kec. Ujung Tanah",
		IsActive: false,
	},
	{
		Name:     "BAROMBONG",
		Code:     "1070769",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Perjanjian Buangaya No.13, Kel. Barombong, Kec. Tamalate",
		IsActive: false,
	},
	{
		Name:     "BATUA",
		Code:     "1070791",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Jl. Abd Daeng Sirua No.338 Kel. Tello Baru, Kec. Panakkukang",
		IsActive: false,
	},
	{
		Name:     "BIRA",
		Code:     "1070803",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Prof. Ir. Sutami No.128, Kec. Tamalanrea",
		IsActive: false,
	},
	{
		Name:     "BULUROKENG",
		Code:     "1070805",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Kec. Biringkanaya",
		IsActive: false,
	},
	{
		Name:     "CENDRAWASIH",
		Code:     "1070768",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Cendrawasih No.404 Kel. Sambung Jawa, Kec. Mamajang",
		IsActive: false,
	},
	{
		Name:     "DAHLIA",
		Code:     "1070764",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Seroja No.3, Kel. Kampung Buyang, Kec. Mariso",
		IsActive: false,
	},
	{
		Name:     "JONGAYA",
		Code:     "1070770",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Jl. Andi Tonro No.70A, Kel. Jongaya, Kec. Tamalate",
		IsActive: false,
	},
	{
		Name:     "JUMPANDANG BARU",
		Code:     "1070788",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Jl. Ade Irma Nasution Blok C No.1, Kel. Rappo Jawa, Kec. Tallo",
		IsActive: false,
	},
	{
		Name:     "KALUKU BODOA",
		Code:     "1070789",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Butta Caddi No.15 Kel. Kaluku Bodoa, Kec. Tallo",
		IsActive: false,
	},
	{
		Name:     "KAPASA",
		Code:     "1070801",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Komp. BTN Angkatan Laut Kec. Biring Kanaya",
		IsActive: false,
	},
	{
		Name:     "KARUWISI",
		Code:     "1070792",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Urip Sumiharjo Lr.2, Kel. Karuwisi, Kec. Panakkukang",
		IsActive: false,
	},
	{
		Name:     "KASSI-KASSI",
		Code:     "1070773",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Jl. Tamalate I No.43, Kel. Kassi-kassi, Kec. Rappocini",
		IsActive: false,
	},
	{
		Name:     "LAYANG",
		Code:     "1070782",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Tinumbu Lr.148 No.1-2, Kel. Layang, Kec. Bontoala",
		IsActive: false,
	},
	{
		Name:     "MACCINI SAWAH",
		Code:     "1070778",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Maccini Raya Kel. Maccini, Kec. Makassar",
		IsActive: false,
	},
	{
		Name:     "MACCINI SOMBALA",
		Code:     "1070772",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Kec. Tamalate",
		IsActive: false,
	},
	{
		Name:     "MAKKASAU",
		Code:     "1070779",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Ratulangi, Kec. Ujung Pandang",
		IsActive: false,
	},
	{
		Name:     "MALIMONGAN BARU",
		Code:     "1070783",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Pontiku Lr.5 No.1, Kel. Malimongan Baru, Kec. Bontoala",
		IsActive: false,
	},
	{
		Name:     "MAMAJANG",
		Code:     "1070767",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Jl. Baji Minasa No.10 Kel. Mamajang Luar, Kec. Mamajang",
		IsActive: false,
	},
	{
		Name:     "MANGASA",
		Code:     "1070775",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Monumen Emmy Saelan Komp BTN M 11, Kec. Rappocini",
		IsActive: false,
	},
	{
		Name:     "MARADEKAYA",
		Code:     "1070777",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Sungai Saddang Baru Lr.5 No.27 Kel. Mandekaya, Kec. Makassar",
		IsActive: false,
	},
	{
		Name:     "MINASA UPA",
		Code:     "1070774",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Jl. Minasa Upa Raya No.18, Kel. Gunung Sari, Kec. Rappocini",
		IsActive: false,
	},
	{
		Name:     "PACCERAKKANG",
		Code:     "1070806",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Kec. Tamalanrea",
		IsActive: false,
	},
	{
		Name:     "PAMPANG",
		Code:     "1070793",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Pampang III No. 28 A, Kec. Panakkukang",
		IsActive: false,
	},
	{
		Name:     "PANAMBUNGAN",
		Code:     "1070765",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Rajawali Lr.13B No.13, Kel. Panambungan, Kec. Mariso",
		IsActive: false,
	},
	{
		Name:     "PATTINGALLOANG",
		Code:     "1070784",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Jl. Barukang VI No.15, Kel. Pattingalloang Lama, Kec. Ujung Tanah",
		IsActive: false,
	},
	{
		Name:     "PERTIWI",
		Code:     "1070766",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Cendrawasih III No. 2, Kel. Mariso, Kec. Mariso",
		IsActive: false,
	},
	{
		Name:     "PULAU KODINGARENG",
		Code:     "1070787",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Pulau Kodingareng Kec. Ujung Tanah",
		IsActive: false,
	},
	{
		Name:     "RAPPOKALLING",
		Code:     "1070790",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Daeng Ragge Kel. Rappokalling, Kec. Tallo",
		IsActive: false,
	},
	{
		Name:     "SUDIANG",
		Code:     "1070800",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Gowa Raya, Kec. Biring Kanaya",
		IsActive: false,
	},
	{
		Name:     "SUDIANG RAYA",
		Code:     "1070799",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Makasar Raya, Kec. Biring Kanaya",
		IsActive: false,
	},
	{
		Name:     "TABARINGAN",
		Code:     "1070786",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Tinumbu Lr.154 Kel. Ujung Tanah, Kec. Ujung Tanah",
		IsActive: false,
	},
	{
		Name:     "TAMALANREA",
		Code:     "1070804",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Kesejahteraan Timur I BTP Blok B. Kec. Tamalanrea",
		IsActive: false,
	},
	{
		Name:     "TAMALANREA JAYA",
		Code:     "1071287",
		Photo:    "healthcenter.png",
		Type:     "Rawat Inap",
		Address:  "Jl. Perintis Kemerdekaan IV No 9, Tamalanrea, Makassar",
		IsActive: false,
	},
	{
		Name:     "TAMALATE",
		Code:     "1070771",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Dg.Tata I (Komp.Tabaria Blok G 8 Kel. Pr Tambung, Kec. Tamalate",
		IsActive: false,
	},
	{
		Name:     "TAMAMAUNG",
		Code:     "1070794",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Abdul Daeng Sirua No.158, Kec. Panakkukang",
		IsActive: false,
	},
	{
		Name:     "TAMANGAPA",
		Code:     "1070797",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Tamangapa Raya No.26 H, Kec. Manggala",
		IsActive: false,
	},
	{
		Name:     "TARAKANG",
		Code:     "1070780",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Kondingareng Lr.181 No.5, Kel. Malimongan Tua, Kec. Wajo",
		IsActive: false,
	},
	{
		Name:     "TODDOPULI",
		Code:     "1071286",
		Photo:    "healthcenter.png",
		Type:     "Non Rawat Inap",
		Address:  "Jl. Toddopuli Raya No 96 Makassar",
		IsActive: false,
	},
}

var medicines = []model.Medicine{
	{
		Name:     "Albendazole 400 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Type:     "Saset",
		Name:     "Alkohol 70 % 1000 ml",
		Unit:     "Botol",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Alkohol Swab",
		Type:     "Saset",
		Unit:     "Saset",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Allopurinol 100 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Ambroksol 30 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Amlodipin 10 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Amlodipin 5 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Amoksisillin 500 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Antasida DOEN I tablet kunyah, kombinasi : Aluminium Hidroksida 200 mg Magnesium Hidroksida 200 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Antihaemorroid suppositoria",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Antimigren: Ergotamin Tartrat 1 mg + Kofein 50 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Asam Askorbat (Vit C) 250 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Asiklovir 200 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Asiklovir 400 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Asiklovir krim 5 %",
		Type:     "Cream",
		Unit:     "Tube",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Atorvastatin 10 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Betametason krim 0.1 %",
		Type:     "Cream",
		Unit:     "Tube",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Bisakodil suppositoria 10 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "BRAUN INTROCAN G24",
		Type:     "Needle",
		Unit:     "Needle",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "BUFACETINE Kloramfenikol salep kulit 2 %",
		Type:     "Cream",
		Unit:     "Tube",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Deksametason 0.5 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Dermafix",
		Type:     "Saset",
		Unit:     "Saset",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Doksisiklin 100 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Domperidon 10 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Fitomenadion (Vitamin K1) 10 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Furosemide 40 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Garam Oralit",
		Type:     "Saset",
		Unit:     "Saset",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Glimepiride 2 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Guaifenesin (GG)",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Haloperidol 1.5 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Hidrokortison krim",
		Type:     "Cream",
		Unit:     "Tube",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Hiosin Butilbromida",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Kalsium Laktat (Kalk) 500 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Kaptopril 50 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Klorfeniramina Maleat (CTM) 4 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Kotrimoksazole 480 mg (Adult)",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Loperamid 2 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Metformin 500 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Metilprednisolon 4 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Mikonazole Nitrat Salep 2%",
		Type:     "Cream",
		Unit:     "Tube",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "NaCl 0.9% 500 ml larutan infus",
		Type:     "Cream",
		Unit:     "Botol",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Natrium Diklofenak 25 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Omeprazole 20 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Ondansetron 4 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Parasetamol 500 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Piridoksin (Vitamin B6) 25 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Salbutamol 2 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Salep 2-4 , Kombinasi : Asam Salisilat 2% + Belerang endap 4%",
		Type:     "Cream",
		Unit:     "Tube",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Setirizin 10 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Simvastatin 10 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Simvastatin 20 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Siprofloksasin 500 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Bisakodil Suppositoria 5 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Salep Whitfield (Antifungi salep komb : As Benzoat 6 % + As. Salisilat 3%)",
		Type:     "Cream",
		Unit:     "Tube",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Asam Asetilsalisilat 80 mg (Miniaspi 80)",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Metronidazole 500 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Tablet Tambah Darah",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Lubricant",
		Type:     "Cream",
		Unit:     "Tube",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Acetyl cystein 200 mg",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Kombinasi Kaolin + Pektin (Neo Diaform)",
		Type:     "Tablet",
		Unit:     "Strip",
		Category: "Obat Dalam",
		IsActive: true,
	},
	{
		Name:     "Kloramfenikol Tetes Telinga 1% (Reco TT 1% 10 ml)",
		Type:     "Tetes",
		Unit:     "Botol",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Acyclovir Cream",
		Type:     "Cream",
		Unit:     "Tube",
		Category: "Obat Luar",
		IsActive: true,
	},
	{
		Name:     "Gentamisin Salep Kulit",
		Type:     "Cream",
		Unit:     "Tube",
		Category: "Obat Luar",
		IsActive: true,
	},
}

var doctors = []model.Doctor{
	{
		HealthcenterID:       1,
		Username:             "dr-ernita",
		Name:                 "dr. Ernita Rosyanti Dewi",
		Photo:                "dr-ernita.png",
		Type:                 "Dokter Umum",
		Number:               "STR 3121100220145544",
		Experience:           5,
		GraduatedFrom:        "Universitas Yarsi, 2013",
		ConsultationSchedule: "06.00 - 23.00 WIB",
		PlacePractice:        "Jakarta Timur, DKI Jakarta",
		Phone:                "6282115353192",
		Start:                time.Date(2020, time.April, 11, 00, 01, 01, 0, time.Local),
		End:                  time.Date(2020, time.April, 11, 23, 59, 01, 0, time.Local),
		IsActive:             true,
	},
	{
		HealthcenterID:       1,
		Username:             "dr-ayu",
		Name:                 "dr. Ayu A. Istiana",
		Photo:                "dr-ayu.png",
		Type:                 "Dokter Umum",
		Number:               "STR 3121100220145699",
		Experience:           7,
		GraduatedFrom:        "Universitas Yarsi, 2013",
		ConsultationSchedule: "06.00 - 23.00 WIB",
		PlacePractice:        "Bogor, Jawa Barat",
		Phone:                "6281288068122",
		Start:                time.Date(2020, time.April, 11, 00, 01, 01, 0, time.Local),
		End:                  time.Date(2020, time.April, 11, 23, 59, 01, 0, time.Local),
		IsActive:             true,
	},
	{
		HealthcenterID:       1,
		Username:             "dr-peter",
		Name:                 "dr. Peter Fernando",
		Photo:                "dr-peter.png",
		Type:                 "Dokter Umum",
		Number:               "STR 6111100120221435",
		Experience:           3,
		GraduatedFrom:        "Universitas Tanjungpura, 2019",
		ConsultationSchedule: "06.00 - 23.00 WIB",
		PlacePractice:        "Ngabang, Kalimantan Timur",
		Phone:                "6281299708787",
		Start:                time.Date(2020, time.April, 11, 00, 01, 01, 0, time.Local),
		End:                  time.Date(2020, time.April, 11, 23, 59, 01, 0, time.Local),
		IsActive:             true,
	},
}

var specialists = []model.Specialist{
	{
		Username:             "dr-men",
		Name:                 "dr. Indra Sandinirwan, Sp.A",
		Photo:                "dr-men.png",
		Type:                 "Spesialis Anak",
		Number:               "STR 61111233456677",
		Experience:           3,
		GraduatedFrom:        "Universitas Tanjungpura, 2019",
		ConsultationSchedule: "06.00 - 23.00 WIB",
		PlacePractice:        "Ngabang, Kalimantan Timur",
		Phone:                "6281299708788",
		Start:                time.Date(2020, time.April, 11, 00, 01, 01, 0, time.Local),
		End:                  time.Date(2020, time.April, 11, 23, 59, 01, 0, time.Local),
		IsActive:             true,
	},
	{
		Username:             "dr-hari",
		Name:                 "dr. Hari Prasetyo Rahardjo, Sp.OG",
		Photo:                "dr-men.png",
		Type:                 "Spesialis Ortopedi",
		Number:               "STR 61111233456677",
		Experience:           3,
		GraduatedFrom:        "Universitas Tanjungpura, 2019",
		ConsultationSchedule: "06.00 - 23.00 WIB",
		PlacePractice:        "Ngabang, Kalimantan Timur",
		Phone:                "6281299708787",
		Start:                time.Date(2020, time.April, 11, 00, 01, 01, 0, time.Local),
		End:                  time.Date(2020, time.April, 11, 23, 59, 01, 0, time.Local),
		IsActive:             true,
	},
}

var officers = []model.Officer{
	{
		HealthcenterID: 1,
		Name:           "Andi",
		Photo:          "officer.png",
		Phone:          "6282115353192",
		IsActive:       true,
	},
	{
		HealthcenterID: 2,
		Name:           "Juli",
		Photo:          "officer.png",
		Phone:          "6281288068122",
		IsActive:       true,
	},
}

var drivers = []model.Driver{
	{
		HealthcenterID: 1,
		Name:           "Alex",
		Photo:          "driver.png",
		Phone:          "6282115353192",
		IsActive:       true,
	},
	{
		HealthcenterID: 2,
		Name:           "Doni",
		Photo:          "driver.png",
		Phone:          "6281288068122",
		IsActive:       true,
	},
}

var apothecaries = []model.Courier{
	{
		HealthcenterID: 1,
		Name:           "Dian",
		Photo:          "pharmacy.png",
		Phone:          "6282115353192",
		IsActive:       true,
	},
	{
		HealthcenterID: 2,
		Name:           "Dony",
		Photo:          "pharmacy.png",
		Phone:          "6281288068122",
		IsActive:       true,
	},
}

var couriers = []model.Courier{
	{
		HealthcenterID: 1,
		Name:           "Dian",
		Photo:          "courier.png",
		Phone:          "6282115353192",
		IsActive:       true,
	},
	{
		HealthcenterID: 2,
		Name:           "Dini",
		Photo:          "courier.png",
		Phone:          "6281288068122",
		IsActive:       true,
	},
}

var admins = []model.Admin{
	{
		HealthcenterID: 1,
		Phone:          "6281299708787",
		Password:       "XXXXXX",
		IsActive:       true,
	},
	{
		HealthcenterID: 2,
		Phone:          "62812997087812",
		Password:       "XXXXXX",
		IsActive:       true,
	},
}

var superadmins = []model.SuperAdmin{
	{
		Phone:    "6281299708787",
		Password: "XXXXXX",
		IsActive: true,
	},
}
