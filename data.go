package main

// store global values here

import (
	"github.com/gokyle/adn"
	"os"
	"path/filepath"
)

const (
	AppName     = "The Social Gopher"
	AppUnixName = "socialgopher"
	AppVersion  = "0.1"
)

var (
	homeDir          = os.ExpandEnv("${HOME}/." + AppUnixName)
	authDatabaseFile = filepath.Join(homeDir, "profiles.db")
)

// The desktop experience should mimic in functionality the web experience.
// Accordingly, we need to be able to do everything.
var Scopes = []string{
	adn.ScopeBasic,
	adn.ScopeStream,
	adn.ScopeEmail,
	adn.ScopeWritePost,
	adn.ScopeFollow,
	adn.ScopeMessages,
	adn.ScopeExport,
}

// App represents the Social Gopher, and stores our client information.
var App = &adn.Application{
	"STUdQHPA8EC9zeKpqd3hWGUC2VzJASxG", // client ID
	"placeholder_secret",               // client secret
	"",                                 // redirect URI
	Scopes,                             // scopes
	"hgq6MjvtePat6ZZwDsJxAMX4Vvq5PvCE", // password secret
}

// Profiles is a list of all the users the app knows about.
var Profiles []*Profile

var (
	BodyTypeJSON = "application/json"
)
