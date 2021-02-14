package main

import (
	"github.com/joho/godotenv"
	"github.com/maxisme/appserver"
	"os"
)

func main() {
	// load .env if there
	_ = godotenv.Load()

	conf := appserver.ProjectConfig{
		Name:        "Pass For It",
		Host:        "passfor.it",
		DmgPath:     "Pass For It.dmg",
		KeyWords:    "mac, osx, password, generator, key, macos, password generator",
		Description: "Easily generate passwords from your menubar.",
		Recaptcha: appserver.Recaptcha{
			Pub:  os.Getenv("captchpub"),
			Priv: os.Getenv("captchpriv"),
		},
		Sparkle: appserver.Sparkle{
			Version:     "0.01",
			Description: "foo",
		},
		Email: appserver.Email{
			To:       "max@max.me.uk",
			Host:     os.Getenv("emailhost"),
			Port:     587,
			Username: os.Getenv("emailusername"),
			Password: os.Getenv("emailpassword"),
		},
		WebPort: 8080,
	}

	if err := appserver.Serve(conf); err != nil {
		panic(err)
	}
}
