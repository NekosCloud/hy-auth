package main

import (
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"
	"github.com/volatiletech/authboss/v3/otp/twofactor"
	// "github.com/volatiletech/authboss/v3/otp/twofactor/totp2fa"
)

func setupAuthboss() {
	// Init authboss configs
	ab.Config.Storage.Server = database
	ab.Config.Storage.CookieState = cookieStore
	ab.Config.Storage.SessionState = sessionStore

	// Paths
	ab.Config.Paths.Mount = "/auth"
	ab.Config.Paths.RootURL = config.Domain

	// Auth paths
	ab.Config.Paths.ConfirmNotOK = "/action/confirm"
	// ab.Config.Paths.NotAuthorized = "/login"
	// ab.Config.Paths.AuthLoginOK = "/account"
	// ab.Config.Paths.LogoutOK = "/logout/ok"

	// Auth methods
	ab.Config.Modules.LogoutMethod = "DELETE"

	// Set renderer
	ab.Config.Core.ViewRenderer = NewHTML("/auth", "ab_views")

	// Set mail renderer
	ab.Config.Core.MailRenderer = NewEmail("/auth", "email_views")

	// Don't loose data when registering users
	ab.Config.Modules.RegisterPreserveFields = []string{"email", "name"}

	// Issuer name for TOTP2FA
	ab.Config.Modules.TOTP2FAIssuer = "N23Account"

	// Redirect user when unauthed
	ab.Config.Modules.ResponseOnUnauthed = authboss.RespondRedirect

	// Require e-mail 2fa
	ab.Config.Modules.TwoFactorEmailAuthRequired = true

	// Init default values for other configs
	defaults.SetCore(&ab.Config, false, false)

	// Setup 2FA recovery
	twofaRecovery := &twofactor.Recovery{Authboss: ab}
	if err := twofaRecovery.Setup(); err != nil {
		panic(err)
	}

	// Setup TOTP 2FA
	// totp := &totp2fa.TOTP{Authboss: ab}
	// if err := totp.Setup(); err != nil {
	// 	panic(err)
	// }

	// Load
	if err := ab.Init(); err != nil {
		panic(err)
	}
}
