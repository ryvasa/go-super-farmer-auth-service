package main

import (
	"log"

	"github.com/gin-gonic/gin"
	wire_auth "github.com/ryvasa/go-super-farmer-auth-service/pkg/wire"
)

// var data, err = env.LoadEnv()

// // Konfigurasi OAuth 2.0
// var googleOauthConfig = &oauth2.Config{
// 	ClientID:     data.Google.ClientID,
// 	ClientSecret: data.Google.ClientSecret,
// 	RedirectURL:  data.Google.RedirectURL,
// 	Scopes: []string{
// 		"https://www.googleapis.com/auth/userinfo.profile",
// 		"https://www.googleapis.com/auth/userinfo.email",
// 	},
// 	Endpoint: google.Endpoint,
// }

// var oauthStateString = "random-state-string"

// func main() {
// 	fmt.Println(data.Google.RedirectURL)
// 	// Route untuk halaman login
// 	http.HandleFunc("/", handleHome)
// 	http.HandleFunc("/login", handleGoogleLogin)
// 	http.HandleFunc("/redirect", handleGoogleCallback)

// 	// Menjalankan server di port 8080
// 	fmt.Println("Server berjalan di http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(data.Server.Port, nil))
// }

// // Halaman utama
// func handleHome(w http.ResponseWriter, r *http.Request) {
// 	html := `<html><body><a href="/login">Login dengan Google</a></body></html>`
// 	fmt.Fprint(w, html)
// }

// // Redirect pengguna ke halaman login Google
// func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
// 	url := googleOauthConfig.AuthCodeURL(oauthStateString)
// 	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
// }

// // Callback setelah login berhasil
// func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
// 	// Validasi state
// 	if r.URL.Query().Get("state") != oauthStateString {
// 		http.Error(w, "State tidak valid", http.StatusBadRequest)
// 		return
// 	}

// 	// Mendapatkan authorization code dari URL
// 	code := r.URL.Query().Get("code")
// 	if code == "" {
// 		http.Error(w, "Kode otorisasi tidak ditemukan", http.StatusBadRequest)
// 		return
// 	}

// 	// Menukar authorization code dengan access token
// 	token, err := googleOauthConfig.Exchange(context.Background(), code)
// 	if err != nil {
// 		http.Error(w, "Gagal menukar kode dengan token", http.StatusInternalServerError)
// 		return
// 	}

// 	// Menggunakan token untuk mendapatkan informasi pengguna
// 	client := googleOauthConfig.Client(context.Background(), token)
// 	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
// 	if err != nil {
// 		http.Error(w, "Gagal mendapatkan data pengguna", http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		http.Error(w, "Gagal membaca data pengguna", http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprint(w, "Login berhasil! Informasi pengguna:\n\n")
// 	fmt.Fprint(w, string(body))

// 	fmt.Println("Data User dari Google API:")
// 	fmt.Println(string(body))
// }

func main() {
	app, err := wire_auth.InitializeAuthApp()
	if err != nil {
		log.Fatal(err)
	}

	app.Router.Use(gin.Recovery())
	app.Router.Use(gin.Logger())
	app.Router.Run(":" + app.Env.Server.Port)
}
