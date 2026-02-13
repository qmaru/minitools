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

// GenerateKey Generate a new TOTP key (default period = 30s)
func (t *TotpBasic) GenerateKey(issuer, accountName string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
	})
}

// GenerateKeyWithPeriod Generate a new TOTP key with custom period
func (t *TotpBasic) GenerateKeyWithPeriod(issuer, accountName string, period uint) (*otp.Key, error) {
	if period == 0 {
		period = 30
	}
	return totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
		Period:      period,
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

// internal: calculate remaining time at specific moment with given period
func (t *TotpBasic) remainingAtWithPeriod(at time.Time, period int64) time.Duration {
	if period <= 0 {
		return 0
	}
	remain := period - (at.Unix() % period)
	return time.Duration(remain) * time.Second
}

// GenerateCodeWithRemaining Generate code and remaining time at the same moment (default period = 30s)
func (t *TotpBasic) GenerateCodeWithRemaining(secret string) (string, time.Duration, error) {
	now := time.Now()

	code, err := t.GenerateCodeAt(secret, now)
	if err != nil {
		return "", 0, err
	}

	period := int64(30)
	return code, t.remainingAtWithPeriod(now, period), nil
}

// RemainingWithPeriod Get remaining time in current TOTP window with given period (e.g. 30)
func (t *TotpBasic) RemainingWithPeriod(period int64) time.Duration {
	return t.remainingAtWithPeriod(time.Now(), period)
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
func (t *TotpBasic) Validate(code, secret string) (bool, error) {
	return t.ValidateAt(code, secret, time.Now())
}

// ValidateAt Validate a TOTP code for the given secret at a specific time
func (t *TotpBasic) ValidateAt(code, secret string, at time.Time) (bool, error) {
	return totp.ValidateCustom(code, secret, at, totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
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
