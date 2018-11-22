package firebase

import "firebase.google.com/go"

type FireStore struct {
	fireApp *firebase.App
}

func NewFireStore(app *firebase.App) *FireStore {
	return &FireStore{fireApp: app}
}
