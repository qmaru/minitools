package totp

import (
	"net/url"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type TotpBasic struct{}

func New() *TotpBasic {
	return &TotpBasic{}
}

// GenerateKey Generate a new TOTP key
func (t *TotpBasic) GenerateKey(issuer, accountName string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
	})
}

// GenerateCode Generate a TOTP code for the given secret (now)
func (t *TotpBasic) GenerateCode(secret string) (string, error) {
	return t.GenerateCodeAt(secret, time.Now())
}

// GenerateCodeAt Generate a TOTP code for the given secret at a specific time
func (t *TotpBasic) GenerateCodeAt(secret string, at time.Time) (string, error) {
	return totp.GenerateCode(secret, at)
}

// GenerateCodeWithRemaining Generate code and remaining time at the same moment
func (t *TotpBasic) GenerateCodeWithRemaining(secret string) (string, time.Duration, error) {
	now := time.Now()

	code, err := t.GenerateCodeAt(secret, now)
	if err != nil {
		return "", 0, err
	}

	period := int64(30)
	remain := period - (now.Unix() % period)

	return code, time.Duration(remain) * time.Second, nil
}

// RemainingWithPeriod Get remaining time in current TOTP window with given period (e.g. 30)
func (t *TotpBasic) RemainingWithPeriod(period int64) time.Duration {
	if period <= 0 {
		return 0
	}
	now := time.Now().Unix()
	remain := period - (now % period)
	return time.Duration(remain) * time.Second
}

// RemainingWithKey Get remaining time based on otp.Key (respects key.Period())
func (t *TotpBasic) RemainingWithKey(key *otp.Key) time.Duration {
	if key == nil {
		return 0
	}
	period := int64(key.Period())
	return t.RemainingWithPeriod(period)
}

// Validate Validate a TOTP code for the given secret (now)
func (t *TotpBasic) Validate(code, secret string) bool {
	return totp.Validate(code, secret)
}

// SecretToString Convert the secret key to a Base32 encoded secret
func (t *TotpBasic) SecretToString(key *otp.Key) string {
	return key.Secret()
}

// URLToString Convert the secret key to an otpauth URL
func (t *TotpBasic) URLToString(key *otp.Key) string {
	return key.URL()
}

// ParseURL Parse a key from an otpauth URL
func (t *TotpBasic) ParseURL(urlStr string) (*otp.Key, error) {
	otpURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	return otp.NewKeyFromURL(otpURL.String())
}

// URLMinimalToString Convert the secret key to an otpauth URL with minimal parameters
func (t *TotpBasic) URLMinimalToString(key *otp.Key) string {
	issuer := key.Issuer()
	account := key.AccountName()
	secret := key.Secret()

	label := account
	if issuer != "" {
		label = issuer + ":" + account
	}
	label = url.PathEscape(label)

	q := url.Values{}
	q.Set("secret", secret)
	if issuer != "" {
		q.Set("issuer", issuer)
	}

	u := url.URL{
		Scheme:   "otpauth",
		Host:     "totp",
		Path:     "/" + label,
		RawQuery: q.Encode(),
	}

	return u.String()
}
