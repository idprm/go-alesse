package util

import (
	"strings"
	"time"

	"github.com/idprm/go-alesse/src/pkg/model"
)

func TimeStamp() string {
	now := time.Now()
	return now.Format("20060102150405")
}

func TrimByteToString(b []byte) string {
	str := string(b)
	return strings.Join(strings.Fields(str), " ")
}

func ContentDoctorToPharmacy(content string, chat model.Chat) string {
	// Hello Admin Farmasi @health_center terdapat pengajuan resep obat dari @doctor untuk pasien @patient Cek disini @link
	replacer := strings.NewReplacer("@health_center", chat.Doctor.Healthcenter.Name, "@doctor", chat.Doctor.Name, "@patient", chat.User.Name, "@link", chat.ChannelUrl)
	content = replacer.Replace(content)
	return content
}

func ContentPharmacyToCourier(content string) string {
	// Hello Kurir @courier, terdapat pemintaan pengantaran obat dari Farmasi @pharmacy untuk pasien @patient. Cek disini @link
	replacer := strings.NewReplacer("@courier", "", "@pharmacy", "", "@patient", "", "@link", "")
	content = replacer.Replace(content)
	return content
}

func ContentCourierToPharmacy(content string) string {
	// Hello Admin Farmasi @health_center, Kurir @courier sudah menyelesaikan pengantaran obat ke pasien @patient
	replacer := strings.NewReplacer("@health_center", "", "@courier", "", "@patient", "")
	content = replacer.Replace(content)
	return content
}

func ContentPharmacyToPatient(content string) string {
	// Hello pasien @patient obat Anda sedang disiapkan oleh Farmasi @health_center
	replacer := strings.NewReplacer("@patient", "", "@health_center", "")
	content = replacer.Replace(content)
	return content
}

func ContentCourierToPatient(content string) string {
	// Hello pasien @patient obat Anda sedang diantarkan oleh Kurir @health_center
	replacer := strings.NewReplacer("@patient", "", "@health_center", "")
	content = replacer.Replace(content)
	return content
}

func ContentHomecareToPatientProgress(content string) string {
	// Hello pasien @patient, tim Homecare @health_center akan datang kerumah Anda dalam waktu 1 jam.
	replacer := strings.NewReplacer("@patient", "", "@health_center", "")
	content = replacer.Replace(content)
	return content
}

func ContentHomecareToPatientDone(content string) string {
	// Hello pasien @patient, layanan homecare dari tim Homecare @health_center sudah selesai dilakukan. Semoga Anda lekas sembuh.
	replacer := strings.NewReplacer("@patient", "", "@health_center", "")
	content = replacer.Replace(content)
	return content
}

func ContentDoctorToHomecare(content string) string {
	// Hello Admin Homecare @health_center, terdapat permintaan layanan homecare dari @doctor untuk pasien @patient Cek disini @link
	replacer := strings.NewReplacer("@health_center", "", "@doctor", "")
	content = replacer.Replace(content)
	return content
}

func ContentHomecareToHealthoffice(content string) string {
	// Hello Admin Dinkes @admin, tim homecare @health_center sudah menyelesaikan layanan homecare untuk pasien @patient
	replacer := strings.NewReplacer("@admin", "", "@health_center", "", "@patient", "")
	content = replacer.Replace(content)
	return content
}
