package totp

import (
	"net/url"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type TotpBasic struct{}

// GenerateKey Generate a new TOTP key
func (t *TotpBasic) GenerateKey(issuer, accountName string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
	})
	if err != nil {
		return nil, err
	}
	return key, nil
}

// GenerateCode Generate a TOTP code for the given secret
func (t *TotpBasic) GenerateCode(secret string) (string, error) {
	return t.GenerateCodeAt(secret, time.Now())
}

// GenerateCodeAt Generate a TOTP code for the given secret at a specific time
func (t *TotpBasic) GenerateCodeAt(secret string, at time.Time) (string, error) {
	return totp.GenerateCode(secret, at)
}

// Validate Validate a TOTP code for the given secret
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
	key, err := otp.NewKeyFromURL(otpURL.String())
	if err != nil {
		return nil, err
	}
	return key, nil
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

func New() *TotpBasic {
	return new(TotpBasic)
}
