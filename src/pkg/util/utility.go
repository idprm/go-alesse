package util

import (
	"crypto/rand"
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

func GenerateOTP(length int) (string, error) {
	const otpChars = "123456789"

	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func ContentOTPToUser(content string, otp string, link string) string {
	// Berikut adalah kode OTP kamu : *@otp* untuk mulai konsultasi dokter di @link
	replacer := strings.NewReplacer("@otp", otp, "@link", link)
	content = replacer.Replace(content)
	return content
}

func StatusOTPToUser(content string, otp string, user model.User) string {
	// Mengirim kode OTP @otp kepada pasien @patient
	replacer := strings.NewReplacer("@otp", otp, "@patient", user.Name)
	content = replacer.Replace(content)
	return content
}

func ContentMessageToUser(content string, homecare model.Homecare, officer model.Officer) string {
	// Hello pasien *@patient*, Apabila ada pertanyaan silakan hubungi nomor ini @phone
	replacer := strings.NewReplacer("@patient", homecare.Chat.User.Name, "@phone", officer.Phone)
	content = replacer.Replace(content)
	return content
}

func StatusMessageToUser(content string, homecare model.Homecare, officer model.Officer) string {
	// Mengirim pesan kontak @phone kepada pasien @patient
	replacer := strings.NewReplacer("@phone", officer.Phone, "@patient", homecare.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func ContentMessageToDoctor(content string, user model.User, doctor model.Doctor, url string) string {
	// Hi *@doctor*, User *@patient* menunggu konfirmasi untuk konsultasi online. Klik disini untuk memulai chat *@link*
	urlWeb := config.ViperEnv("APP_HOST") + "/chat/" + url
	replacer := strings.NewReplacer("@doctor", doctor.Name, "@patient", user.Name, "@link", urlWeb)
	content = replacer.Replace(content)
	return content
}

func StatusMessageToDoctor(content string, user model.User, doctor model.Doctor) string {
	// Pasien @patient mengajukan Konsultasi Dokter @doctor
	replacer := strings.NewReplacer("@patient", user.Name, "@doctor", doctor.Name)
	content = replacer.Replace(content)
	return content
}

func PushMessageToDoctor(content string, user model.User, doctor model.Doctor) string {
	// Hi @doctor, User @patient menunggu konfirmasi untuk konsultasi online.
	replacer := strings.NewReplacer("@doctor", doctor.Name, "@patient", user.Name)
	content = replacer.Replace(content)
	return content
}

func ContentMessageToSpecialist(content string, specialist model.Specialist, doctor model.Doctor, url string) string {
	// Hi *@specialist*, Dokter Umum *@doctor* menunggu konfirmasi untuk konsultasi online. Klik disini untuk memulai chat @link
	urlWeb := config.ViperEnv("APP_HOST") + "/specialist/chat/" + url
	replacer := strings.NewReplacer("@specialist", specialist.Name, "@doctor", doctor.Name, "@link", urlWeb)
	content = replacer.Replace(content)
	return content
}

func StatusMessageToSpecialist(content string, specialist model.Specialist, doctor model.Doctor) string {
	// Dokter umum @doctor mengajukan Konsultasi dengan Dokter spesialis @specialist
	replacer := strings.NewReplacer("@doctor", doctor.Name, "@specialist", specialist.Name)
	content = replacer.Replace(content)
	return content
}

func PushMessageToSpecialist(content string, specialist model.Specialist, doctor model.Doctor) string {
	// Hi *@specialist*, Dokter Umum *@doctor* menunggu konfirmasi untuk konsultasi online.
	replacer := strings.NewReplacer("@specialist", specialist.Name, "@doctor", doctor.Name)
	content = replacer.Replace(content)
	return content
}

func ContentDoctorToPharmacy(content string, pharmacy model.Pharmacy) string {
	// Hello Admin Farmasi @health_center terdapat pengajuan resep obat dari @doctor untuk pasien @patient Cek disini @link
	urlWeb := config.ViperEnv("APP_HOST") + "/pharmacy/process/" + pharmacy.Chat.ChannelUrl
	replacer := strings.NewReplacer("@health_center", pharmacy.Chat.Healthcenter.Name, "@doctor", pharmacy.Chat.Doctor.Name, "@patient", pharmacy.Chat.User.Name, "@link", urlWeb)
	content = replacer.Replace(content)
	return content
}

func StatusDoctorToPharmacy(content string, pharmacy model.Pharmacy) string {
	// Dokter @doctor meresepkan e-Resep @number
	replacer := strings.NewReplacer("@doctor", pharmacy.Chat.Doctor.Name, "@number", pharmacy.Number)
	content = replacer.Replace(content)
	return content
}

func PushDoctorToPharmacy(content string, pharmacy model.Pharmacy) string {
	// Hello Admin Farmasi @health_center terdapat pengajuan resep obat dari @doctor untuk pasien @patient
	replacer := strings.NewReplacer("@health_center", pharmacy.Chat.Healthcenter.Name, "@doctor", pharmacy.Chat.Doctor.Name, "@patient", pharmacy.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func ContentPharmacyToCourier(content string, pharmacy model.Pharmacy, courier model.Courier) string {
	// Hello Kurir @courier, terdapat pemintaan pengantaran obat dari Farmasi @pharmacy untuk pasien @patient. Cek disini @link
	urlWeb := config.ViperEnv("APP_HOST") + "/pharmacy/take/" + pharmacy.Chat.ChannelUrl
	replacer := strings.NewReplacer("@courier", courier.Name, "@pharmacy", pharmacy.Chat.Doctor.Name, "@patient", pharmacy.Chat.User.Name, "@link", urlWeb)
	content = replacer.Replace(content)
	return content
}

func StatusPharmacyToCourier(content string, pharmacy model.Pharmacy, courier model.Courier) string {
	// e-Resep @number telah dibuat. Menunggu Kurir @courier untuk mengambil obat.
	replacer := strings.NewReplacer("@number", pharmacy.Number, "@courier", courier.Name)
	content = replacer.Replace(content)
	return content
}

func PushPharmacyToCourier(content string, pharmacy model.Pharmacy, courier model.Courier) string {
	// Hello Kurir @courier, terdapat pemintaan pengantaran obat dari Farmasi @pharmacy untuk pasien @patient.
	replacer := strings.NewReplacer("@courier", courier.Name, "@pharmacy", pharmacy.Chat.Doctor.Name, "@patient", pharmacy.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func ContentCourierToPharmacy(content string, pharmacy model.Pharmacy, courier model.Courier) string {
	// Hello Admin Farmasi @health_center, Kurir @courier sudah menyelesaikan pengantaran obat ke pasien @patient
	replacer := strings.NewReplacer("@health_center", pharmacy.Chat.Healthcenter.Name, "@courier", courier.Name, "@patient", pharmacy.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func StatusCourierToPharmacy(content string, pharmacy model.Pharmacy, courier model.Courier) string {
	// Pasien @patient telah menerima obat dari Kurir @courier.
	replacer := strings.NewReplacer("@patient", pharmacy.Chat.User.Name, "@courier", courier.Name)
	content = replacer.Replace(content)
	return content
}

func PushCourierToPharmacy(content string, pharmacy model.Pharmacy, courier model.Courier) string {
	// Hello Admin Farmasi @health_center, Kurir @courier sudah menyelesaikan pengantaran obat ke pasien @patient
	replacer := strings.NewReplacer("@health_center", pharmacy.Chat.Healthcenter.Name, "@courier", courier.Name, "@patient", pharmacy.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func ContentPharmacyToPatient(content string, pharmacy model.Pharmacy) string {
	// Hello pasien @patient obat Anda sedang disiapkan oleh Farmasi @health_center
	replacer := strings.NewReplacer("@patient", pharmacy.Chat.User.Name, "@health_center", pharmacy.Chat.Healthcenter.Name)
	content = replacer.Replace(content)
	return content
}

func StatusPharmacyToPatient(content string, pharmacy model.Pharmacy) string {
	// Farmasi Puskesmas @health_center telah memverifikasi dan sedang menyiapkan e-Resep @number.
	replacer := strings.NewReplacer("@health_center", pharmacy.Chat.Healthcenter.Name, "@number", pharmacy.Number)
	content = replacer.Replace(content)
	return content
}

func PushPharmacyToPatient(content string, pharmacy model.Pharmacy) string {
	// Hello pasien @patient obat Anda sedang disiapkan oleh Farmasi @health_center
	replacer := strings.NewReplacer("@patient", pharmacy.Chat.User.Name, "@health_center", pharmacy.Chat.Healthcenter.Name)
	content = replacer.Replace(content)
	return content
}

func ContentCourierToPatient(content string, pharmacy model.Pharmacy) string {
	// Hello pasien @patient obat Anda sedang diantarkan oleh Kurir @health_center
	replacer := strings.NewReplacer("@patient", pharmacy.Chat.User.Name, "@health_center", pharmacy.Chat.Healthcenter.Name)
	content = replacer.Replace(content)
	return content
}

func StatusCourierToPatient(content string, pharmacy model.Pharmacy) string {
	// Kurir telah mengambil obat di Farmasi Puskesmas @health_center dan menuju alamat Pasien @patient.
	replacer := strings.NewReplacer("@health_center", pharmacy.Chat.Healthcenter.Name, "@patient", pharmacy.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func ContentDoctorToHomecare(content string, homecare model.Homecare) string {
	// Hello Admin Homecare @health_center, terdapat permintaan layanan homecare dari @doctor untuk pasien @patient Cek disini @link
	urlWeb := config.ViperEnv("APP_HOST") + "/homecare/visit/" + homecare.Chat.ChannelUrl
	replacer := strings.NewReplacer("@health_center", homecare.Chat.Healthcenter.Name, "@doctor", homecare.Chat.Doctor.Name, "@patient", homecare.Chat.User.Name, "@link", urlWeb)
	content = replacer.Replace(content)
	return content
}

func PushCourierToPatient(content string, pharmacy model.Pharmacy) string {
	// Kurir telah mengambil obat di Farmasi Puskesmas @health_center dan menuju alamat Pasien @patient.
	replacer := strings.NewReplacer("@health_center", pharmacy.Chat.Healthcenter.Name, "@patient", pharmacy.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func StatusDoctorToHomecare(content string, homecare model.Homecare) string {
	// Dokter @doctor mengajukan Layanan Homecare untuk Pasien @patient
	replacer := strings.NewReplacer("@doctor", homecare.Chat.Doctor.Name, "@patient", homecare.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func PushDoctorToHomecare(content string, homecare model.Homecare) string {
	// Hello Admin Homecare @health_center, terdapat permintaan layanan homecare dari @doctor untuk pasien @patient
	replacer := strings.NewReplacer("@health_center", homecare.Chat.Healthcenter.Name, "@doctor", homecare.Chat.Doctor.Name, "@patient", homecare.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func ContentDoctorToPatientHomecare(content string, homecare model.Homecare) string {
	// Dokter @doctor telah menjadwalkan kunjungan Homecare pada tanggal @visit_at pukul @hour untuk Pasien @patient
	replacer := strings.NewReplacer("@doctor", homecare.Chat.Doctor.Name, "@visit_at", homecare.VisitAt.Format("02-01-2006"), "@hour", homecare.VisitAt.Format("15:04"))
	content = replacer.Replace(content)
	return content
}

func ContentHomecareToPatientProgress(content string, homecare model.Homecare, officer model.Officer) string {
	//Hello pasien *@patient*, tim Homecare *@health_center* akan datang kerumah Anda dalam waktu 15 menit - 1 jam. Apabila ada pertanyaan silakan hubungi nomor ini *@phone*
	replacer := strings.NewReplacer("@patient", homecare.Chat.User.Name, "@health_center", homecare.Chat.Healthcenter.Name, "@phone", officer.Phone)
	content = replacer.Replace(content)
	return content
}

func StatusHomecareToPatientProgress(content string, homecare model.Homecare) string {
	//Tim Homecare Puskesmas @health_center telah memverifikasi Permintaan Homecare Dokter @doctor dan sedang menuju alamat Pasien @patient.
	replacer := strings.NewReplacer("@health_center", homecare.Chat.Healthcenter.Name, "@doctor", homecare.Chat.Doctor.Name, "@patient", homecare.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func PushHomecareToPatientProgress(content string, homecare model.Homecare, officer model.Officer) string {
	//Hello pasien *@patient*, tim Homecare *@health_center* akan datang kerumah Anda dalam waktu 15 menit - 1 jam. Apabila ada pertanyaan silakan hubungi nomor ini *@phone*
	replacer := strings.NewReplacer("@patient", homecare.Chat.User.Name, "@health_center", homecare.Chat.Healthcenter.Name, "@phone", officer.Phone)
	content = replacer.Replace(content)
	return content
}

func ContentHomecareToPatientDone(content string, homecare model.Homecare) string {
	// Hello pasien @patient, layanan homecare dari tim Homecare @health_center sudah selesai dilakukan. Semoga Anda lekas sembuh. @link
	urlWeb := config.ViperEnv("APP_HOST") + "/feedback/" + homecare.Chat.ChannelUrl
	replacer := strings.NewReplacer("@patient", homecare.Chat.User.Name, "@health_center", homecare.Chat.Doctor.Name, "@link", urlWeb)
	content = replacer.Replace(content)
	return content
}

func StatusHomecareToPatientDone(content string, homecare model.Homecare) string {
	// Tim Homecare Puskesmas @health_center telah mengunjungi dan memeriksa Pasien @patient.
	replacer := strings.NewReplacer("@health_center", homecare.Chat.Doctor.Name, "@patient", homecare.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func PushHomecareToPatientDone(content string, homecare model.Homecare) string {
	// Hello pasien @patient, layanan homecare dari tim Homecare @health_center sudah selesai dilakukan. Semoga Anda lekas sembuh.
	replacer := strings.NewReplacer("@patient", homecare.Chat.User.Name, "@health_center", homecare.Chat.Doctor.Name)
	content = replacer.Replace(content)
	return content
}

func ContentHomecareToHealthoffice(content string, homecare model.Homecare) string {
	// Hello Admin Dinkes, tim homecare @health_center sudah menyelesaikan layanan homecare untuk pasien @patient
	replacer := strings.NewReplacer("@health_center", homecare.Chat.Healthcenter.Name, "@patient", homecare.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func StatusHomecareToHealthoffice(content string, homecare model.Homecare) string {
	// Aktivitas Homecare Puskesmas @health_center untuk pasien @patient telah selesai
	replacer := strings.NewReplacer("@health_center", homecare.Chat.Healthcenter.Name, "@patient", homecare.Chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func PushHomecareToHealthoffice(content string, homecare model.Homecare) string {
	// Hello Admin Dinkes, tim homecare @health_center sudah menyelesaikan layanan homecare untuk pasien @patient
	replacer := strings.NewReplacer("@health_center", homecare.Chat.Healthcenter.Name, "@patient", homecare.Chat.User.Name)
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

func StatusFeedbackToPatient(content string, chat model.Chat) string {
	//Mengirim link feedback kepada pasien @patient
	replacer := strings.NewReplacer("@patient", chat.User.Name)
	content = replacer.Replace(content)
	return content
}

func PushFeedbackToPatient(content string, chat model.Chat) string {
	//Hello pasien *@patient*, Semoga Anda lekas sembuh. Seberapa puaskah Anda dengan layanan puskesmas *@health_center* ?
	replacer := strings.NewReplacer("@patient", chat.User.Name, "@health_center", chat.Healthcenter.Name)
	content = replacer.Replace(content)
	return content
}
