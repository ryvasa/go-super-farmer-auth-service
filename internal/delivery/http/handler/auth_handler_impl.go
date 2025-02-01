package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthHandlerImpl struct {
	oauth *oauth2.Config
}

func NewAuthHandlerImpl(oauth *oauth2.Config) AuthHandler {
	return &AuthHandlerImpl{oauth}
}

func (a *AuthHandlerImpl) Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home Page",
	})
}

var oauthStateString = "random-state-string"

func (a *AuthHandlerImpl) GoogleLogin(ctx *gin.Context) {
	url := a.oauth.AuthCodeURL(oauthStateString)
	ctx.Redirect(http.StatusFound, url)
}

func (a *AuthHandlerImpl) Redirect(ctx *gin.Context) {
	// Memeriksa state
	if ctx.Query("state") != oauthStateString {
		ctx.String(http.StatusBadRequest, "State tidak valid")
		return
	}

	// Mendapatkan authorization code dari URL
	code := ctx.Query("code")
	if code == "" {
		ctx.String(http.StatusBadRequest, "Kode otorisasi tidak ditemukan")
		return
	}

	// Menukar authorization code dengan access token
	token, err := a.oauth.Exchange(context.Background(), code)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Gagal menukar kode dengan token")
		return
	}

	// Menggunakan token untuk mendapatkan informasi pengguna
	client := a.oauth.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Gagal mendapatkan data pengguna")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Gagal membaca data pengguna")
		return
	}

	// Menampilkan informasi pengguna
	responseText := fmt.Sprintf("Login berhasil! Informasi pengguna:\n\n%s", string(body))
	ctx.String(http.StatusOK, responseText)

	fmt.Println("Data User dari Google API:")
	fmt.Println(string(body))
}
