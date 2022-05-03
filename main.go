package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/aarondl/tpl"
	"github.com/go-chi/chi/v5"

	// "github.com/gorilla/sessions"

	"github.com/volatiletech/authboss/v3"

	_ "github.com/volatiletech/authboss/v3/auth"
	"github.com/volatiletech/authboss/v3/confirm"
	"github.com/volatiletech/authboss/v3/lock"

	abclientstate "github.com/volatiletech/authboss-clientstate"
	_ "github.com/volatiletech/authboss/v3/logout"
	_ "github.com/volatiletech/authboss/v3/recover"
	_ "github.com/volatiletech/authboss/v3/register"
	// "github.com/volatiletech/authboss/v3/remember"
)

var funcs = template.FuncMap{
	"yield": func() string { return "" },
}

var (
	// authboss
	ab       = authboss.New()
	config   = readConfig()
	database = NewStorer()

	sessionStore abclientstate.SessionStorer
	cookieStore  abclientstate.CookieStorer

	templates tpl.Templates
)

func main() {
	// Load templates
	templates = tpl.Must(tpl.Load("views", "views/partials", "layout.tpl.html", funcs))

	// Init stores
	cookieStore = abclientstate.NewCookieStorer(cookieStoreKey, cookieStoreEncryption)
	sessionStore = abclientstate.NewSessionStorer("ab_session", sessionStoreKey, sessionStoreEncryption)
	// cstore := sessionStore.Store.(*sessions.CookieStore)

	// Init Authboss
	setupAuthboss()

	// Webserver
	mux := chi.NewRouter()

	// The middlewares we're using:
	// - nosurfing is a more verbose wrapper around csrf handling
	// - LoadClientStateMiddleware is required for session/cookie stuff
	// - dataInjector is for putting data into the request context we need for our template layout
	mux.Use(nosurfing, ab.LoadClientStateMiddleware, dataInjector)

	// Routes
	mux.Group(func(mux chi.Router) {
		mux.Use(authboss.ModuleListMiddleware(ab))
		mux.Mount("/auth", http.StripPrefix("/auth", ab.Config.Core.Router))
	})

	// Server static files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", neuter(fs)))

	// Server rotues (requires auth)
	mux.Group(func(mux chi.Router) {
		mux.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondRedirect), lock.Middleware(ab), confirm.Middleware(ab))
		mux.MethodFunc("GET", "/account", renderAccount)
	})

	// Server routes
	mux.Get("/", renderIndex)
	mux.Get("/privacy", renderPrivacy)
	mux.Get("/terms", renderTerms)

	// Start server
	log.Printf("App listening on: %s", config.Domain+":"+config.Port)
	log.Println(http.ListenAndServe(config.Domain+":"+config.Port, mux))
}
