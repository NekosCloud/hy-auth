package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/otp/twofactor/sms2fa"
	"github.com/volatiletech/authboss/v3/otp/twofactor/totp2fa"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func NewStorer() Storer {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "authboss", options.Client().ApplyURI(readConfig().Database), options.Client().SetRetryWrites(true), options.Client().SetWriteConcern(writeconcern.New(writeconcern.WMajority())))

	if err != nil {
		panic(err)
	}

	return Storer{}
}

type User struct {
	// Default mongo
	mgm.DefaultModel `bson:",inline"`

	// Custom
	Name string `json:"name" bson:"name"`

	// PID
	ID string `json:"id" bson:"id"`

	// Auth
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`

	// Confirm
	ConfirmSelector string `json:"confirm_selector" bson:"confirm_selector"`
	ConfirmVerifier string `json:"confirm_verifier" bson:"confirm_verifier"`
	Confirmed       bool   `json:"confirmed" bson:"confirmed"`

	// Lock
	AttemptCount int       `json:"attempt_count" bson:"attempt_count"`
	LastAttempt  time.Time `json:"last_attempt" bson:"last_attempt"`
	Locked       time.Time `json:"locked" bson:"locked"`

	// Recover
	RecoverSelector    string    `json:"recover_selector" bson:"recover_selector"`
	RecoverVerifier    string    `json:"recover_verifier" bson:"recover_verifier"`
	RecoverTokenExpiry time.Time `json:"recover_token_expiry" bson:"recover_token_expiry"`

	// 2fa
	TOTPSecretKey      string `json:"totp_secret_key" bson:"totp_secret_key"`
	SMSPhoneNumber     string `json:"sms_phone_number" bson:"sms_phone_number"`
	SMSSeedPhoneNumber string `json:"sms_seed_phone_number" bson:"sms_seed_phone_number"`
	RecoveryCodes      string `json:"recovery_codes" bson:"recovery_codes"`
}

// type RememberTokenObject struct {
// 	// Default mongo
// 	mgm.DefaultModel `bson:",inline"`

// 	// PID
// 	ID string `json:"id" bson:"id"`

// 	// Token
// 	Token string `json:"token" bson:"token"`
// }

// func NewRememberTokenObject(id string, token string) *RememberTokenObject {
// 	return &RememberTokenObject{
// 		ID:    id,
// 		Token: token,
// 	}
// }

var (
	assertUser   = &User{}
	assertStorer = &Storer{}

	_ authboss.User            = assertUser
	_ authboss.AuthableUser    = assertUser
	_ authboss.ConfirmableUser = assertUser
	_ authboss.LockableUser    = assertUser
	_ authboss.RecoverableUser = assertUser
	_ authboss.ArbitraryUser   = assertUser

	_ totp2fa.User = assertUser
	_ sms2fa.User  = assertUser

	_ authboss.CreatingServerStorer   = assertStorer
	_ authboss.ConfirmingServerStorer = assertStorer
	_ authboss.RecoveringServerStorer = assertStorer
	// _ authboss.RememberingServerStorer = assertStorer
)

// Write
// PutPID into user
func (u *User) PutPID(pid string) {
	u.Email = pid
}

// PutPassword into user
func (u *User) PutPassword(password string) {
	u.Password = password
}

// PutEmail into user
func (u *User) PutEmail(email string) {
	u.Email = email
}

// PutConfirmed into user
func (u *User) PutConfirmed(confirmed bool) {
	u.Confirmed = confirmed
}

// PutConfirmSelector into user
func (u *User) PutConfirmSelector(confirmSelector string) {
	u.ConfirmSelector = confirmSelector
}

// PutConfirmVerifier into user
func (u *User) PutConfirmVerifier(confirmVerifier string) {
	u.ConfirmVerifier = confirmVerifier
}

// PutLocked into user
func (u *User) PutLocked(locked time.Time) {
	u.Locked = locked
}

// PutAttemptCount into user
func (u *User) PutAttemptCount(attempts int) {
	u.AttemptCount = attempts
}

// PutLastAttempt into user
func (u *User) PutLastAttempt(last time.Time) {
	u.LastAttempt = last
}

// PutRecoverSelector into user
func (u *User) PutRecoverSelector(token string) {
	u.RecoverSelector = token
}

// PutRecoverVerifier into user
func (u *User) PutRecoverVerifier(token string) {
	u.RecoverVerifier = token
}

// PutRecoverExpiry into user
func (u *User) PutRecoverExpiry(expiry time.Time) {
	u.RecoverTokenExpiry = expiry
}

// PutTOTPSecretKey into user
func (u *User) PutTOTPSecretKey(key string) {
	u.TOTPSecretKey = key
}

// PutSMSPhoneNumber into user
func (u *User) PutSMSPhoneNumber(key string) {
	u.SMSPhoneNumber = key
}

// PutRecoveryCodes into user
func (u *User) PutRecoveryCodes(key string) {
	u.RecoveryCodes = key
}

// PutArbitrary into user
func (u *User) PutArbitrary(values map[string]string) {
	if n, ok := values["name"]; ok {
		u.Name = n
	}
}

// Read
// GetPID from user
func (u User) GetPID() string {
	return u.Email
}

// GetPassword from user
func (u User) GetPassword() string {
	return u.Password
}

// GetEmail from user
func (u User) GetEmail() string {
	return u.Email
}

// GetConfirmed from user
func (u User) GetConfirmed() bool {
	return u.Confirmed
}

// GetConfirmSelector from user
func (u User) GetConfirmSelector() string {
	return u.ConfirmSelector
}

// GetConfirmVerifier from user
func (u User) GetConfirmVerifier() string {
	return u.ConfirmVerifier
}

// GetLocked from user
func (u User) GetLocked() time.Time {
	return u.Locked
}

// GetAttemptCount from user
func (u User) GetAttemptCount() int {
	return u.AttemptCount
}

// GetLastAttempt from user
func (u User) GetLastAttempt() time.Time {
	return u.LastAttempt
}

// GetRecoverSelector from user
func (u User) GetRecoverSelector() string {
	return u.RecoverSelector
}

// GetRecoverVerifier from user
func (u User) GetRecoverVerifier() string {
	return u.RecoverVerifier
}

// GetRecoverExpiry from user
func (u User) GetRecoverExpiry() time.Time {
	return u.RecoverTokenExpiry
}

// GetTOTPSecretKey from user
func (u User) GetTOTPSecretKey() string {
	return u.TOTPSecretKey
}

// GetSMSPhoneNumber from user
func (u User) GetSMSPhoneNumber() string {
	return u.SMSPhoneNumber
}

// GetSMSPhoneNumberSeed from user
func (u User) GetSMSPhoneNumberSeed() string {
	return u.SMSSeedPhoneNumber
}

// GetRecoveryCodes from user
func (u User) GetRecoveryCodes() string {
	return u.RecoveryCodes
}

// GetArbitrary from user
func (u User) GetArbitrary() map[string]string {
	return map[string]string{
		"name": u.Name,
	}
}

// Storer
type Storer struct{}

// Save user to DB
func (s Storer) Save(_ context.Context, user authboss.User) error {
	u := user.(*User)

	var result User

	query := mgm.Coll(u).FindOne(mgm.Ctx(), &bson.M{"email": u.Email})

	if query.Err() != nil {
		log.Println(query.Err().Error())
		return authboss.ErrUserNotFound
	}

	query.Decode(&result)

	if err := mgm.Coll(u).Update(u); err != nil {
		return errors.New("DB save failed")
	}

	return nil
}

// Load user from DB
func (s Storer) Load(_ context.Context, key string) (user authboss.User, err error) {
	var result User

	query := mgm.Coll(&User{}).FindOne(mgm.Ctx(), &bson.M{"email": key})

	if query.Err() != nil {
		return nil, authboss.ErrUserNotFound
	}

	query.Decode(&result)

	u := &result

	return u, nil
}

// Get struct for new user
func (s Storer) New(_ context.Context) authboss.User {
	return &User{}
}

// Create new user in DB
func (s Storer) Create(_ context.Context, user authboss.User) error {
	u := user.(*User)

	check := mgm.Coll(&User{}).FindOne(mgm.Ctx(), bson.M{"email": u.Email})

	var v mgm.Model
	check.Decode(v)
	if v != nil {
		return authboss.ErrUserFound
	}

	mgm.Coll(u).Create(u)

	return nil
}

// Get user by confirm selector
func (s Storer) LoadByConfirmSelector(_ context.Context, selector string) (user authboss.ConfirmableUser, err error) {
	var result User

	query := mgm.Coll(&User{}).FindOne(mgm.Ctx(), &bson.M{"confirm_selector": selector})

	if query.Err() != nil {
		return nil, authboss.ErrUserNotFound
	}

	query.Decode(&result)

	u := &result

	return u, nil
}

// Get user by recover selector
func (s Storer) LoadByRecoverSelector(_ context.Context, selector string) (user authboss.RecoverableUser, err error) {
	var result User

	query := mgm.Coll(&User{}).FindOne(mgm.Ctx(), &bson.M{"recover_selector": selector})

	if query.Err() != nil {
		return nil, authboss.ErrUserNotFound
	}

	query.Decode(&result)

	u := &result

	return u, nil
}

// // Add remember token to account in MemStore
// func (s Storer) AddRememberToken(_ context.Context, pid, token string) error {
// 	s.Mem[pid] = append(s.Mem[pid], token)
// 	return nil
// }

// // Delete all remember tokens from account in MemStore
// func (s Storer) DelRememberTokens(_ context.Context, pid string) error {
// 	delete(s.Mem, pid)
// 	return nil
// }

// // Find account + token pair and delete from MemStore
// func (s Storer) UseRememberToken(_ context.Context, pid, token string) error {
// 	tokens, ok := s.Mem[pid]
// 	if !ok {
// 		return authboss.ErrTokenNotFound
// 	}

// 	for i, tok := range tokens {
// 		if tok == token {
// 			tokens[len(tokens)-1] = tokens[i]
// 			s.Mem[pid] = tokens[:len(tokens)-1]
// 			return nil
// 		}
// 	}

// 	return authboss.ErrTokenNotFound
// }
