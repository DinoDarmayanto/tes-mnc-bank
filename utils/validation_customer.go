package utils

import "regexp"

func IsValidPassword(password string) bool {
	// Contoh validasi password:
	// Minimal 8 karakter
	// Mengandung setidaknya satu huruf besar, satu huruf kecil, dan satu angka

	// Minimal 8 karakter
	if len(password) < 8 {
		return false
	}

	// Mengandung setidaknya satu huruf besar, satu huruf kecil, dan satu angka
	hasUppercase := false
	hasLowercase := false
	hasDigit := false

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUppercase = true
		case 'a' <= char && char <= 'z':
			hasLowercase = true
		case '0' <= char && char <= '9':
			hasDigit = true
		}

		if hasUppercase && hasLowercase && hasDigit {
			return true
		}
	}

	return false
}

func IsValidEmail(email string) bool {
	// Pattern regex untuk validasi email
	// Contoh: menggunakan regex yang sederhana untuk mengecek format email
	// Sesuaikan dengan kebijakan validasi email yang Anda tetapkan
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}
