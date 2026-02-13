package totp

import (
	"net/url"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const defaultPeriod = 30

type TotpBasic struct{}

func New() *TotpBasic {
	return &TotpBasic{}
}

// GenerateKey Generate a new TOTP key with optional custom period (default 30s)
func (t *TotpBasic) GenerateKey(issuer, accountName string, period ...uint) (*otp.Key, error) {
	p := uint(defaultPeriod)
	if len(period) > 0 && period[0] > 0 {
		p = period[0]
	}
	return totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
		Period:      p,
	})
}

func (t *TotpBasic) validateOpts(period, skew uint) totp.ValidateOpts {
	if period == 0 {
		period = defaultPeriod
	}
	return totp.ValidateOpts{
		Period:    period,
		Skew:      skew,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	}
}

// GenerateCode Generate a TOTP code for current time with optional period (default 30s)
func (t *TotpBasic) GenerateCode(secret string, period ...uint) (string, error) {
	return t.GenerateCodeAt(secret, time.Now(), period...)
}

// GenerateCodeAt Generate a TOTP code at a given time with optional period (default 30s)
func (t *TotpBasic) GenerateCodeAt(secret string, at time.Time, period ...uint) (string, error) {
	p := uint(defaultPeriod)
	if len(period) > 0 && period[0] > 0 {
		p = period[0]
	}
	return totp.GenerateCodeCustom(secret, at, t.validateOpts(p, 0))
}

// GenerateCodeWithRemaining Generate code and remaining time at the same moment
func (t *TotpBasic) GenerateCodeWithRemaining(secret string, period ...uint) (string, time.Duration, error) {
	now := time.Now()
	p := uint(defaultPeriod)
	if len(period) > 0 && period[0] > 0 {
		p = period[0]
	}

	code, err := totp.GenerateCodeCustom(secret, now, t.validateOpts(p, 0))
	if err != nil {
		return "", 0, err
	}

	return code, t.remaining(now, p), nil
}

// Remaining Get remaining time in current TOTP window with optional period (default 30s)
func (t *TotpBasic) Remaining(period ...uint) time.Duration {
	p := uint(defaultPeriod)
	if len(period) > 0 && period[0] > 0 {
		p = period[0]
	}
	return t.remaining(time.Now(), p)
}

// RemainingFromKey Get remaining time in current TOTP window from a key
func (t *TotpBasic) RemainingFromKey(key *otp.Key) time.Duration {
	p := uint(defaultPeriod)
	if key != nil {
		p = uint(key.Period())
	}
	return t.remaining(time.Now(), p)
}

func (t *TotpBasic) remaining(at time.Time, period uint) time.Duration {
	if period == 0 {
		return 0
	}
	remain := int64(period) - (at.Unix() % int64(period))
	return time.Duration(remain) * time.Second
}

// Validate Validate a TOTP code (default: now, period 30s, skew 1)
func (t *TotpBasic) Validate(code, secret string, at ...time.Time) (bool, error) {
	now := time.Now()
	if len(at) > 0 {
		now = at[0]
	}
	return totp.ValidateCustom(code, secret, now, t.validateOpts(defaultPeriod, 1))
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

	q := url.Values{}
	q.Set("secret", secret)
	if issuer != "" {
		q.Set("issuer", issuer)
	}

	u := url.URL{
		Scheme:   "otpauth",
		Host:     "totp",
		Path:     "/" + url.PathEscape(label),
		RawQuery: q.Encode(),
	}

	return u.String()
}
