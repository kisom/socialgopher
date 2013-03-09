package main

import (
	"database/sql"
	"fmt"
	"github.com/gokyle/adn"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// A profile is the coupling of the App value with an identity. It is provided
// so that API calls can act on a single value with all the information
// needed to make API calls.
type Profile struct {
	App      *adn.Application
	Identity *Identity
}

// An identity contains the minimum infomration required to identify a user.
// An empty secret indicates the user is not authentication.
type Identity struct {
	User   string
	secret string
}

func (i *Identity) Secret() string {
	return i.secret
}

// AuthPassword attempts to retrieve an access token using the identity's
// password.
func (i *Identity) AuthPassword(pass string) (err error) {
	t, err := App.PasswordToken(i.User, pass)
	if err != nil {
		return
	}

	i.secret = t
	return
}

// Authenticated checks whether the identity has an access token. It does not
// attempt to access the ADN API using that token.
func (i *Identity) Authenticated() bool {
	if i.Secret() == "" {
		return false
	}
	return true
}

// Create profile attempts to log into the service as the given user. If
// successful, it returns a new profile. Otherwise, it returns an error.
// This should only be called after initialisation, as it uses the Profiles
// global value.
func CreateProfile(user, pass string) (p *Profile, err error) {
	i := new(Identity)
	i.User = user
	t, err := App.PasswordToken(user, pass)
	if err != nil {
		return
	}
	i.secret = t
	p = new(Profile)
	p.App = App
	p.Identity = i
	Profiles = append(Profiles, p)
	return
}

// Store saves the profile to the database.
func (p *Profile) Store() (err error) {
	db, err := sql.Open("sqlite3", authDatabaseFile)
	if err != nil {
		return
	}
	defer db.Close()

	_, err = db.Exec(`delete from sg_profiles where user=?`,
		p.Identity.User)
	query := `insert into sg_profiles values (?, ?)`
	_, err = db.Exec(query, p.Identity.User, p.Identity.Secret())
	return

}

// LoadProfile loads the profile (if one exists) for the specified user from
// the database.
func LoadProfile(user string) (p *Profile, err error) {
	db, err := sql.Open("sqlite3", authDatabaseFile)
	if err != nil {
		return
	}
	defer db.Close()
	row := db.QueryRow(`select secret from sg_profiles where user=?`, user)

	var secret string
	err = row.Scan(&secret)
	if err == nil {
		p = new(Profile)
		p.Identity = new(Identity)
		p.App = App
		p.Identity.User = user
		p.Identity.secret = secret
	}
	return
}

// getUsersInDB retrieves a list of all users in the database.
func getUsersInDB() (users []string, err error) {
	users = make([]string, 0)
	db, err := sql.Open("sqlite3", authDatabaseFile)
	if err != nil {
		return
	}
	defer db.Close()

	query := `select user from sg_profiles`
	rows, err := db.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var user string
		err = rows.Scan(&user)
		if err != nil {
			break
		}
		users = append(users, user)
	}
	return
}

// Load every user stored in the database.
func LoadProfiles() (pl []*Profile, err error) {
	pl = make([]*Profile, 0)
	users, err := getUsersInDB()
	if err != nil {
		return
	}

	for _, u := range users {
		var p *Profile
		p, err = LoadProfile(u)
		if err != nil {
			break
		}
		pl = append(pl, p)
	}
	return
}

// SelectProfile returns the Profile for the user matching the passed-in
// string argument.
func SelectProfile(user string) (p *Profile) {
	for _, profile := range Profiles {
		if profile.Identity.User == user {
			p = profile
			break
		}
	}
	return
}

// AppDirectory checks for the presence of ~/.socialgopher, and creates it
// if it doesn't exist.
func checkAppDirectory() {
	_, err := os.Stat(homeDir)
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir(homeDir, 0700)
	}
	if err != nil {
		panic(err.Error())
	}
}

func checkDatabase() {
	db, err := sql.Open("sqlite3", authDatabaseFile)
	if err != nil {
		return
	}
	defer db.Close()

	var missingTable = fmt.Errorf("no such table: sg_profiles")
	_, err = LoadProfiles()
	if err != nil && err.Error() == missingTable.Error() {
		fmt.Println("creating table")
		err = createDatabase()
	}
	if err != nil {
		panic("[!] socialgopher: opening profile database: " +
			err.Error())
	}
}

func createDatabase() (err error) {
	db, err := sql.Open("sqlite3", authDatabaseFile)
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(`create table sg_profiles
                                (user unique not null primary key,
                                 secret not null)`)
	return
}

func init() {
	var err error
	checkAppDirectory()
	checkDatabase()
	Profiles, err = LoadProfiles()
	if err != nil {
		fmt.Println("[!] couldn't load profiles:", err.Error())
		os.Exit(1)
	}
}
