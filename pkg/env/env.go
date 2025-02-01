package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		User     string
		Password string
		Name     string
		Port     string
		Timezone string
	}
	Secret struct {
		JwtSecretKey string
	}
	Google struct {
		ClientID     string
		ClientSecret string
		RedirectURL  string
	}
}

func LoadEnv() (*Env, error) {
	// GODOTENV for development only
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return nil, err
	}
	env := &Env{}

	// Load Server Config
	env.Server.Port = os.Getenv("SERVER_PORT")

	// Load Database Config
	env.Database.Host = os.Getenv("DB_HOST")
	env.Database.User = os.Getenv("DB_USER")
	env.Database.Port = os.Getenv("DB_PORT")
	env.Database.Password = os.Getenv("DB_PASSWORD")
	env.Database.Name = os.Getenv("DB_NAME")
	env.Database.Timezone = os.Getenv("DB_TIMEZONE")

	// Secret
	env.Secret.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")

	// Google
	env.Google.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	env.Google.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	env.Google.RedirectURL = os.Getenv("GOOGLE_REDIRECT_URL")

	// Debug print
	log.Printf("Database Config: Host=%s, User=%s, Port=%s, DBName=%s",
		env.Database.Host,
		env.Database.User,
		env.Database.Port,
		env.Database.Name)

	return env, nil
}
