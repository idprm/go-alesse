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
		Value: "Berikut adalah kode OTP kamu : *@v1* untuk mulai konsultasi dokter di a-lesse.com",
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
		Value: "Hello Admin Farmasi @health_center terdapat pengajuan resep obat dari @doctor untuk pasien @patient Cek disini @link",
	},
	{
		Name:  "NOTIF_PHARMACY_TO_COURIER",
		Value: "Hello Kurir @courier, terdapat pemintaan pengantaran obat dari Farmasi @pharmacy untuk pasien @patient. Cek disini @link ",
	},
	{
		Name:  "NOTIF_COURIER_TO_PHARMACY",
		Value: "Hello Admin Farmasi @health_center, Kurir @courier sudah menyelesaikan pengantaran obat ke pasien @patient",
	},
	{
		Name:  "NOTIF_PHARMACY_TO_PATIENT",
		Value: "Hello pasien @patient obat Anda sedang disiapkan oleh Farmasi @health_center",
	},
	{
		Name:  "NOTIF_COURIER_TO_PATIENT",
		Value: "Hello pasien @patient obat Anda sedang diantarkan oleh Kurir @health_center",
	},
	{
		Name:  "NOTIF_HOMECARE_TO_PATIENT_PROGRESS",
		Value: "Hello pasien @patient, tim Homecare @health_center akan datang kerumah Anda dalam waktu 1 jam.",
	},
	{
		Name:  "NOTIF_HOMECARE_TO_PATIENT_DONE",
		Value: "Hello pasien @patient, layanan homecare dari tim Homecare @health_center sudah selesai dilakukan. Semoga Anda lekas sembuh.",
	},
	{
		Name:  "NOTIF_DOCTOR_TO_HOMECARE",
		Value: "Hello Admin Homecare @health_center, terdapat permintaan layanan homecare dari @doctor untuk pasien @patient Cek disini @link",
	},
	{
		Name:  "NOTIF_HOMECARE_TO_HEALTHOFFICE",
		Value: "Hello Admin Dinkes @admin, tim homecare @health_center sudah menyelesaikan layanan homecare untuk pasien @patient",
	},
}

var healthcenters = []model.Healthcenter{
	{
		Name:    "ANDALAS",
		Code:    "1070781",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Sangir Lr.209 No.6, Kel. Melayu, Kec. Wajo",
	},
	{
		Name:    "ANTANG",
		Code:    "1070795",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Antang Raya No.43, Kel. Antang, Kec. Manggala",
	},
	{
		Name:    "ANTANG PERUMNAS",
		Code:    "1070796",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Jl. Lasuloro Kel. Manggalo, Kec. Manggala",
	},
	{
		Name:    "ANTARA",
		Code:    "1070802",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Komp.BTN Antara Blok B No.6, Kec. Tamalanrea",
	},
	{
		Name:    "BALLAPARANG",
		Code:    "1071285",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Nikel III No 1 Makassar",
	},
	{
		Name:    "BANGKALA",
		Code:    "1070798",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Kec. Manggala",
	},
	{
		Name:    "BARA-BARAYA",
		Code:    "1070776",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Jl. Abubakar Lambogo, Kel. Bara-Baraya, Kec. Makassar",
	},
	{
		Name:    "BARANG LOMPO",
		Code:    "1070785",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Jl. Pulau Barraang Lompo Rw 1, Kec. Ujung Tanah",
	},
	{
		Name:    "BAROMBONG",
		Code:    "1070769",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Perjanjian Buangaya No.13, Kel. Barombong, Kec. Tamalate",
	},
	{
		Name:    "BATUA",
		Code:    "1070791",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Jl. Abd Daeng Sirua No.338 Kel. Tello Baru, Kec. Panakkukang",
	},
	{
		Name:    "BIRA",
		Code:    "1070803",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Prof. Ir. Sutami No.128, Kec. Tamalanrea",
	},
	{
		Name:    "BULUROKENG",
		Code:    "1070805",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Kec. Biringkanaya",
	},
	{
		Name:    "CENDRAWASIH",
		Code:    "1070768",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Cendrawasih No.404 Kel. Sambung Jawa, Kec. Mamajang",
	},
	{
		Name:    "DAHLIA",
		Code:    "1070764",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Seroja No.3, Kel. Kampung Buyang, Kec. Mariso",
	},
	{
		Name:    "JONGAYA",
		Code:    "1070770",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Jl. Andi Tonro No.70A, Kel. Jongaya, Kec. Tamalate",
	},
	{
		Name:    "JUMPANDANG BARU",
		Code:    "1070788",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Jl. Ade Irma Nasution Blok C No.1, Kel. Rappo Jawa, Kec. Tallo",
	},
	{
		Name:    "KALUKU BODOA",
		Code:    "1070789",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Butta Caddi No.15 Kel. Kaluku Bodoa, Kec. Tallo",
	},
	{
		Name:    "KAPASA",
		Code:    "1070801",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Komp. BTN Angkatan Laut Kec. Biring Kanaya",
	},
	{
		Name:    "KARUWISI",
		Code:    "1070792",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Urip Sumiharjo Lr.2, Kel. Karuwisi, Kec. Panakkukang",
	},
	{
		Name:    "KASSI-KASSI",
		Code:    "1070773",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Jl. Tamalate I No.43, Kel. Kassi-kassi, Kec. Rappocini",
	},
	{
		Name:    "LAYANG",
		Code:    "1070782",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Tinumbu Lr.148 No.1-2, Kel. Layang, Kec. Bontoala",
	},
	{
		Name:    "MACCINI SAWAH",
		Code:    "1070778",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Maccini Raya Kel. Maccini, Kec. Makassar",
	},
	{
		Name:    "MACCINI SOMBALA",
		Code:    "1070772",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Kec. Tamalate",
	},
	{
		Name:    "MAKKASAU",
		Code:    "1070779",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Ratulangi, Kec. Ujung Pandang",
	},
	{
		Name:    "MALIMONGAN BARU",
		Code:    "1070783",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Pontiku Lr.5 No.1, Kel. Malimongan Baru, Kec. Bontoala",
	},
	{
		Name:    "MAMAJANG",
		Code:    "1070767",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Jl. Baji Minasa No.10 Kel. Mamajang Luar, Kec. Mamajang",
	},
	{
		Name:    "MANGASA",
		Code:    "1070775",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Monumen Emmy Saelan Komp BTN M 11, Kec. Rappocini",
	},
	{
		Name:    "MARADEKAYA",
		Code:    "1070777",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Sungai Saddang Baru Lr.5 No.27 Kel. Mandekaya, Kec. Makassar",
	},
	{
		Name:    "MINASA UPA",
		Code:    "1070774",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Jl. Minasa Upa Raya No.18, Kel. Gunung Sari, Kec. Rappocini",
	},
	{
		Name:    "PACCERAKKANG",
		Code:    "1070806",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Kec. Tamalanrea",
	},
	{
		Name:    "PAMPANG",
		Code:    "1070793",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Pampang III No. 28 A, Kec. Panakkukang",
	},
	{
		Name:    "PANAMBUNGAN",
		Code:    "1070765",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Rajawali Lr.13B No.13, Kel. Panambungan, Kec. Mariso",
	},
	{
		Name:    "PATTINGALLOANG",
		Code:    "1070784",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Jl. Barukang VI No.15, Kel. Pattingalloang Lama, Kec. Ujung Tanah",
	},
	{
		Name:    "PERTIWI",
		Code:    "1070766",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Cendrawasih III No. 2, Kel. Mariso, Kec. Mariso",
	},
	{
		Name:    "PULAU KODINGARENG",
		Code:    "1070787",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Pulau Kodingareng Kec. Ujung Tanah",
	},
	{
		Name:    "RAPPOKALLING",
		Code:    "1070790",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Daeng Ragge Kel. Rappokalling, Kec. Tallo",
	},
	{
		Name:    "SUDIANG",
		Code:    "1070800",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Gowa Raya, Kec. Biring Kanaya",
	},
	{
		Name:    "SUDIANG RAYA",
		Code:    "1070799",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Makasar Raya, Kec. Biring Kanaya",
	},
	{
		Name:    "TABARINGAN",
		Code:    "1070786",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Tinumbu Lr.154 Kel. Ujung Tanah, Kec. Ujung Tanah",
	},
	{
		Name:    "TAMALANREA",
		Code:    "1070804",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Kesejahteraan Timur I BTP Blok B. Kec. Tamalanrea",
	},
	{
		Name:    "TAMALANREA JAYA",
		Code:    "1071287",
		Photo:   "healthcenter.png",
		Type:    "Rawat Inap",
		Address: "Jl. Perintis Kemerdekaan IV No 9, Tamalanrea, Makassar",
	},
	{
		Name:    "TAMALATE",
		Code:    "1070771",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Dg.Tata I (Komp.Tabaria Blok G 8 Kel. Pr Tambung, Kec. Tamalate",
	},
	{
		Name:    "TAMAMAUNG",
		Code:    "1070794",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Abdul Daeng Sirua No.158, Kec. Panakkukang",
	},
	{
		Name:    "TAMANGAPA",
		Code:    "1070797",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Tamangapa Raya No.26 H, Kec. Manggala",
	},
	{
		Name:    "TARAKANG",
		Code:    "1070780",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Kondingareng Lr.181 No.5, Kel. Malimongan Tua, Kec. Wajo",
	},
	{
		Name:    "TODDOPULI",
		Code:    "1071286",
		Photo:   "healthcenter.png",
		Type:    "Non Rawat Inap",
		Address: "Jl. Toddopuli Raya No 96 Makassar",
	},
}

var diseases = []model.Disease{
	{
		Name:     "Demam berdarah akibat virus",
		IsActive: true,
	},
	{
		Name:     "Kejang demam",
		IsActive: true,
	},
	{
		Name:     "Keracunan makanan",
		IsActive: true,
	},
	{
		Name:     "Penyakit jantung koroner",
		IsActive: true,
	},
	{
		Name:     "Penyakit gaya hidup",
		IsActive: true,
	},
}

var medicines = []model.Medicine{
	{
		Type:     "Tablet",
		Name:     "Paracetamol",
		Unit:     "mg",
		IsActive: true,
	},
	{
		Type:     "Tablet",
		Name:     "Cetrizine",
		Unit:     "mg",
		IsActive: true,
	},
	{
		Type:     "Kapsul",
		Name:     "Lansoprazole",
		Unit:     "mg",
		IsActive: true,
	},
	{
		Type:     "Tablet",
		Name:     "Amoxicillin",
		Unit:     "mg",
		IsActive: true,
	},
	{
		Type:     "Kapsul",
		Name:     "Omeprazole",
		Unit:     "mg",
		IsActive: true,
	},
	{
		Type:     "Tablet",
		Name:     "Ranitidine",
		Unit:     "mg",
		IsActive: true,
	},
	{
		Type:     "Tablet",
		Name:     "Cefixime",
		Unit:     "mg",
		IsActive: true,
	},
	{
		Type:     "Tablet",
		Name:     "Loperamide",
		Unit:     "mg",
		IsActive: true,
	},
	{
		Type:     "Tablet",
		Name:     "Aminophylline",
		Unit:     "mg",
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
		Phone:                "6281299708787",
		Start:                time.Now(),
		End:                  time.Now(),
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
		Phone:                "6281299708787",
		Start:                time.Now(),
		End:                  time.Now(),
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
		Start:                time.Now(),
		End:                  time.Now(),
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
		Phone:                "6281299708787",
		Start:                time.Now(),
		End:                  time.Now(),
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
		Start:                time.Now(),
		End:                  time.Now(),
		IsActive:             true,
	},
}

var officers = []model.Officer{
	{
		HealthcenterID: 1,
		Name:           "Andi",
		Photo:          "officer.png",
		Phone:          "081299708787",
		IsActive:       true,
	},
	{
		HealthcenterID: 1,
		Name:           "Juli",
		Photo:          "officer.png",
		Phone:          "081299708787",
		IsActive:       true,
	},
}

var drivers = []model.Driver{
	{
		HealthcenterID: 1,
		Name:           "Alex",
		Photo:          "driver.png",
		Phone:          "081299708787",
		IsActive:       true,
	},
	{
		HealthcenterID: 1,
		Name:           "Doni",
		Photo:          "driver.png",
		Phone:          "081299708787",
		IsActive:       true,
	},
}

var apothecaries = []model.Courier{
	{
		HealthcenterID: 1,
		Name:           "Dian",
		Photo:          "pharmacy.png",
		Phone:          "081299708787",
		IsActive:       true,
	},
	{
		HealthcenterID: 1,
		Name:           "Dony",
		Photo:          "pharmacy.png",
		Phone:          "081299708787",
		IsActive:       true,
	},
}

var couriers = []model.Courier{
	{
		HealthcenterID: 1,
		Name:           "Dian",
		Photo:          "courier.png",
		Phone:          "081299708787",
		IsActive:       true,
	},
	{
		HealthcenterID: 1,
		Name:           "Dian",
		Photo:          "courier.png",
		Phone:          "081299708787",
		IsActive:       true,
	},
}
