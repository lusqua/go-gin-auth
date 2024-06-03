package jwt

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/MicahParks/jwkset"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

const (
	KeyID = "kid"
)

var (
	PrivateKey  *rsa.PrivateKey
	PublicKey   *rsa.PublicKey
	ServerStore jwkset.Storage
)

func JWKSetup() {

	var err error

	// Generate RSA private key.
	PrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate RSA private key.\nError: %s", err)
	}

	// Extract the public key from the private key.
	PublicKey = &PrivateKey.PublicKey

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Turn the RSA public key into a JWK.
	marshalOptions := jwkset.JWKMarshalOptions{
		Private: false, // Only public key is shared for validation
	}
	metadata := jwkset.JWKMetadataOptions{
		KID: KeyID,
	}
	options := jwkset.JWKOptions{
		Marshal:  marshalOptions,
		Metadata: metadata,
	}
	jwk, err := jwkset.NewJWKFromKey(PublicKey, options)
	if err != nil {
		log.Fatalf("Failed to create a JWK from the RSA public key.\nError: %s", err)
	}

	// Write the JWK to the server's storage.
	ServerStore = jwkset.NewMemoryStorage()
	err = ServerStore.KeyWrite(ctx, jwk)
	if err != nil {
		log.Fatalf("Failed to write the JWK to the server's storage.\nError: %s", err)
	}
}

func SignClaim(claim jwt.Claims) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	token.Header[KeyID] = KeyID
	token.Claims = claim

	signed, err := token.SignedString(PrivateKey)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func GetClaim(token string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(
		token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return PublicKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	return t.Claims.(jwt.MapClaims), nil
}
