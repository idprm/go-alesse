package util

import (
	"strings"
	"time"

	"github.com/idprm/go-alesse/src/pkg/config"
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

func ContentNotifToUser(content string, homecare model.Homecare) string {
	// Hello pasien *@user*, Apabila ada pertanyaan silakan hubungi nomor ini 08126853852
	replacer := strings.NewReplacer("@user", homecare.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func ContentDoctorToPharmacy(content string, pharmacy model.Pharmacy) string {
	// Hello Admin Farmasi @health_center terdapat pengajuan resep obat dari @doctor untuk pasien @patient Cek disini @link
	urlWeb := config.ViperEnv("APP_HOST") + "/pharmacy/process/" + pharmacy.Chat.ChannelUrl
	replacer := strings.NewReplacer("@health_center", pharmacy.Chat.Doctor.Healthcenter.Name, "@doctor", pharmacy.Chat.Doctor.Name, "@patient", pharmacy.Chat.User.Name, "@link", urlWeb)
	content = replacer.Replace(content)
	return content
}

func ContentPharmacyToCourier(content string, pharmacy model.Pharmacy, courier model.Courier) string {
	urlWeb := config.ViperEnv("APP_HOST") + "/pharmacy/take/" + pharmacy.Chat.ChannelUrl
	// Hello Kurir @courier, terdapat pemintaan pengantaran obat dari Farmasi @pharmacy untuk pasien @patient. Cek disini @link
	replacer := strings.NewReplacer("@courier", courier.Name, "@pharmacy", pharmacy.Chat.Doctor.Name, "@patient", pharmacy.Chat.User.Name, "@link", urlWeb)
	content = replacer.Replace(content)
	return content
}

func ContentCourierToPharmacy(content string, pharmacy model.Pharmacy, courier model.Courier) string {
	// Hello Admin Farmasi @health_center, Kurir @courier sudah menyelesaikan pengantaran obat ke pasien @patient
	replacer := strings.NewReplacer("@health_center", pharmacy.Chat.Doctor.Healthcenter.Name, "@courier", courier.Name, "@patient", pharmacy.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func ContentPharmacyToPatient(content string, pharmacy model.Pharmacy) string {
	// Hello pasien @patient obat Anda sedang disiapkan oleh Farmasi @health_center
	replacer := strings.NewReplacer("@patient", pharmacy.Chat.User.Name, "@health_center", pharmacy.Chat.Doctor.Healthcenter.Name)
	content = replacer.Replace(content)
	return content
}

func ContentCourierToPatient(content string, pharmacy model.Pharmacy) string {
	// Hello pasien @patient obat Anda sedang diantarkan oleh Kurir @health_center
	replacer := strings.NewReplacer("@patient", pharmacy.Chat.User.Name, "@health_center", pharmacy.Chat.Doctor.Healthcenter.Name)
	content = replacer.Replace(content)
	return content
}

func ContentDoctorToHomecare(content string, homecare model.Homecare) string {
	// Hello Admin Homecare @health_center, terdapat permintaan layanan homecare dari @doctor untuk pasien @patient Cek disini @link
	urlWeb := config.ViperEnv("APP_HOST") + "/pharmacy/visit/" + homecare.Chat.ChannelUrl
	replacer := strings.NewReplacer("@health_center", homecare.Chat.Healthcenter.Name, "@doctor", homecare.Chat.Doctor.Name, "@patient", homecare.Chat.User.Name, "@link", urlWeb)
	content = replacer.Replace(content)
	return content
}

func ContentHomecareToPatientProgress(content string, homecare model.Homecare) string {
	// Hello pasien @patient, tim Homecare @health_center akan datang kerumah Anda dalam waktu 1 jam.
	replacer := strings.NewReplacer("@patient", homecare.Chat.User.Name, "@health_center", homecare.Chat.Doctor.Healthcenter.Name)
	content = replacer.Replace(content)
	return content
}

func ContentHomecareToPatientDone(content string, homecare model.Homecare) string {
	// Hello pasien @patient, layanan homecare dari tim Homecare @health_center sudah selesai dilakukan. Semoga Anda lekas sembuh. @link
	urlWeb := config.ViperEnv("APP_HOST") + "/feedback/" + homecare.Chat.ChannelUrl
	replacer := strings.NewReplacer("@patient", homecare.Chat.User.Name, "@health_center", homecare.Chat.Doctor.Healthcenter.Name, "@link", urlWeb)
	content = replacer.Replace(content)
	return content
}

func ContentHomecareToHealthoffice(content string, homecare model.Homecare) string {
	// Hello Admin Dinkes, tim homecare @health_center sudah menyelesaikan layanan homecare untuk pasien @patient
	replacer := strings.NewReplacer("@health_center", homecare.Chat.Doctor.Healthcenter.Name, "@patient", homecare.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func ContentFeedbackToPatient(content string, chat model.Chat) string {
	//Hello pasien *@patient*, Semoga Anda lekas sembuh. Seberapa puaskah Anda dengan layanan puskesmas *@health_center* ? *@link*
	urlWeb := config.ViperEnv("APP_HOST") + "/feedback/" + chat.ChannelUrl
	replacer := strings.NewReplacer("@patient", chat.User.Name, "@health_center", chat.Healthcenter.Name, "@link", urlWeb)
	content = replacer.Replace(content)
	return content
}
