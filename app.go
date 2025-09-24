package main

import (
	"context"
	"fmt"

	"github.com/ProtonMail/gopenpgp/v3/crypto"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Encrypt returns a encrypted string for the given message and private key
func (a *App) Encrypt(message string, pubkey string) string {
	pgp := crypto.PGP()
	publicKey, err := crypto.NewKeyFromArmored(pubkey)
	if err != nil {
		return ""
	}

	encHandle, err := pgp.Encryption().Recipient(publicKey).New()
	if err != nil {
		return ""
	}

	pgpMessage, err := encHandle.Encrypt([]byte(message))
	if err != nil {
		return ""
	}

	armored, err := pgpMessage.ArmorBytes()
	if err != nil {
		return ""
	}

	// fmt.Printf("%s", armored)
	return fmt.Sprintf("%s", armored)
}

// Sign returns a signed string for the given message, private key and passphrase
func (a *App) Sign(message string, privkey string, password string) string {
	pgp := crypto.PGP()

	passphrase := []byte(password)
	privateKey, err := crypto.NewPrivateKeyFromArmored(privkey, passphrase)
	if err != nil {
		return ""
	}

	signer, err := pgp.Sign().SigningKey(privateKey).New()
	if err != nil {
		return ""
	}

	cleartextArmored, err := signer.SignCleartext([]byte(message))
	if err != nil {
		return ""
	}

	signer.ClearPrivateParams()

	// fmt.Printf("%s", cleartextArmored)
	return fmt.Sprintf("%s", cleartextArmored)
}

// Decrypt returns a encrypted string for the given message and private key
func (a *App) Decrypt(message string, privkey string, password string) string {
	pgp := crypto.PGP()

	passphrase := []byte(password)
	privateKey, err := crypto.NewPrivateKeyFromArmored(privkey, passphrase)
	if err != nil {
		return ""
	}

	decHandle, err := pgp.Decryption().DecryptionKey(privateKey).New()
	if err != nil {
		return ""
	}

	decrypted, err := decHandle.Decrypt([]byte(message), crypto.Armor)
	if err != nil {
		return ""
	}

	decHandle.ClearPrivateParams()

	// fmt.Printf("%s", decrypted)
	return fmt.Sprintf("%s", decrypted)
}
