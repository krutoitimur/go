package routes

import (
	"beastmode/packages/utils"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var (
	store *sessions.CookieStore
)

func init() {
	store = sessions.NewCookieStore([]byte("dp%sco2%sa[2mni12zpmy%%vqf_w!enk"))
}

type User struct {
	UserID   int    `db:"UserID"`
	Username string `db:"Username"`
	Hash     []byte `db:"Hash"`
}

func Auth(db *sqlx.DB) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		userID, ok := session.Values["userID"].(int)
		if !ok || userID == 0 {
			utils.RenderTemplate(w, "auth.html", nil)
			return
		}

		var user User
		err := db.Get(&user, "SELECT UserID, Username FROM bcrypt WHERE UserID = ?", userID)
		if err != nil {
			session.Options = &sessions.Options{MaxAge: -1}
			session.Save(r, w)
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
	}).Methods("GET")

	r.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		username := r.FormValue("username")
		password := r.FormValue("password")

		var user User
		// get user by username from db

		err := db.Get(&user, "SELECT UserID, Hash FROM bcrypt WHERE Username = ?", username)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword(user.Hash, []byte(password))
		if err != nil {
			fmt.Print(" ", err)
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		session, _ := store.Get(r, "session")
		session.Values["userID"] = user.UserID
		err = session.Save(r, w)
		if err != nil {
			fmt.Print(" ", err)
			http.Error(w, "Session error", http.StatusUnauthorized)
			return
		}
	}).Methods("POST")

	return r
}

func Reg(db *sqlx.DB) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {

		utils.RenderTemplate(w, "registration.html", nil)

	}).Methods("GET")

	r.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("X-Frame-Options", "DENY")
		username := r.FormValue("username")
		password := r.FormValue("password")
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		//check if username is exists
		var count int
		err = db.Get(&count, "SELECT COUNT(*) FROM bcrypt WHERE Username = ?", username)

		if err != nil || count > 0 {
			http.Error(w, "This username is already exists", http.StatusInternalServerError)
			return
		}
		//add new user
		_, err = db.Exec("INSERT INTO bcrypt (Username, Hash) VALUES (?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(w, "Sign up went wrong..", http.StatusInternalServerError)
			return
		}

	}).Methods("POST")

	return r
}
