package customerusecase

import "errors"

var ErrUserNotFound = errors.New("User tidak ditemukan")
var ErrNikAlreadyUsed = errors.New("Nik sudah digunakan")
var ErrInvalidPassword = errors.New("Password salah")
var ErrInvalidToken = errors.New("Token tidak valid")
var ErrExpiredToken = errors.New("Token kadaluwarsa")
var ErrTokenNotFound = errors.New("Token tidak ditemukan")
var ErrInvalidUser = errors.New("Email atau password salah")
var ErrPhoneNumberAlreadyUsed = errors.New("Phone Number sudah digunakan")
var ErrApps = errors.New("Loan Application Error")
