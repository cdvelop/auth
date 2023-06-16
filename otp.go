package auth

import (
	"time"
)

// One-Time Passwords (OTP)
type otp struct {
	key     string
	created time.Time
}

// newOTP crea y agrega un nuevo OTP al mapa
func (a *auth) newOTP() otp {
	o := otp{
		key:     buildUniqueKey(32),
		created: time.Now(),
	}

	a.rm[o.key] = o
	return o
}

// verifyOTP se asegurará de que exista un OTP
// y devolverá true en caso afirmativo
// También eliminará la clave para que no se pueda reutilizar
func (a *auth) verifyOTP(otp string) bool {
	// Verify OTP is existing
	if _, ok := a.rm[otp]; !ok {
		// otp does not exist
		return false
	}
	delete(a.rm, otp)
	return true
}

// esta función se asegurará periódicamente que se eliminen los "OTP" antiguos
// que han expirado según el período de retención especificado.
// La función utiliza un ticker para realizar la retención en intervalos regulares
// y se detendrá si se recibe una señal de cancelación desde el contexto.
// es bloqueante, así que ejecútelo como una goroutine
func (a *auth) retention(retentionPeriod time.Duration) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			for _, otp := range a.rm {
				// comprobar si ha caducado
				if otp.created.Add(retentionPeriod).Before(time.Now()) {
					delete(a.rm, otp.key)
				}
			}
		case <-a.ctx.Done():
			return

		}
	}
}
