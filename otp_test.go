package auth

import (
	"testing"
)

func TestRetentionMap_VerifyOTP(t *testing.T) {
	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)

	// rm := newRetentionOtpMap(ctx, 1*time.Second)

	// otp := rm.newOTP()

	// if ok := rm.verifyOTP(otp.key); !ok {
	// 	t.Error("falló al verificar la clave OTP que existe")
	// }
	// if ok := rm.verifyOTP(otp.key); ok {
	// 	t.Error("No se debe permitir reutilizar un OTP")
	// }

	// cancel()
}

func TestOTP_Retention(t *testing.T) {

	// // Crear contexto con cancelación para detener la goroutine
	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)

	// // Crear retention_otp y agregar algunos OTP con algunos segundos de diferencia
	// rm := newRetentionOtpMap(ctx, 1*time.Second)

	// rm.newOTP()
	// rm.newOTP()

	// time.Sleep(2 * time.Second)

	// otp := rm.newOTP()

	// // Asegurarse de que solo quede una contraseña y coincida con la última
	// if len(rm) != 1 {
	// 	t.Error("No se pudo limpiar")
	// }

	// if rm[otp.key] != otp {
	// 	t.Error("La clave debería estar en su lugar")
	// }
	// cancel()
}
