package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

const usage = `
USAGE:
	hmac sign signing-key text-to-sign
	hmac verify signing-key signature text-to-verify
`

// Exit codes
// iota resets to 0 every time you start
// a const block, and increments by 1
// each time you use it.
const (
	_                        = iota //ignore first usage (which is zero)
	exitCodeUsage                   // = 1
	exitCodeProcessing              // = 2
	exitCodeInvalidSignature        // = 3
)

// showUsage shows the usage string and exits
// with the code exitCodeUsage.
func showUsage() {
	fmt.Println(usage)
	os.Exit(exitCodeUsage)
}

// sign returns a base64-encoded HMAC signature (essentially an encrypted hash) given a
// signingKey and a read stream. It returns an error if
// there was an error reading from the stream.
func sign(signingKey string, text string) (string, error) {

	// Convert the string to byte
	// because HMAC function operates on byte level.
	key := []byte(signingKey)
	txt := []byte(text)

	// Create a new HMAC hasher.
	h := hmac.New(sha256.New, key)

	h.Write(txt)

	// Calculate the HMAC signature.
	signature := h.Sum(nil)

	return base64.URLEncoding.EncodeToString(signature), nil
}

// verify returns true if the base64-encoded HMAC `signature`
// matches , or false if otherwise.
// If there is an error decoding the base64 signature,
// this will return false and the error.
func verify(signingKey string, signature string, text string) (bool, error) {
	sig1, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return false, fmt.Errorf("error base64-decoding: %v", err)
	}

	h := hmac.New(sha256.New, []byte(signingKey))

	h.Write([]byte(text))

	sig2 := h.Sum(nil)

	// Prevent timing attacks.
	return subtle.ConstantTimeCompare(sig1, sig2) == 1, err
}

func main() {
	if len(os.Args) < 3 {
		showUsage()
	}

	command := strings.ToLower(os.Args[1])
	signingKey := os.Args[2]
	text := os.Args[3]

	switch command {
	case "sign":
		sig, err := sign(signingKey, text)
		if err != nil {
			fmt.Printf("error signing: %v", err)
			os.Exit(exitCodeProcessing)
		}
		fmt.Printf(sig)

	case "verify":
		sig64 := os.Args[3]
		if len(sig64) == 0 {
			showUsage()
		}

		valid, err := verify(signingKey, sig64, text)
		if err != nil {
			fmt.Printf("error validating %v", err)
			os.Exit(exitCodeProcessing)
		}

		if valid {
			fmt.Println("Valid Signature")
		} else {
			fmt.Println("Invalid Signature")
		}

	default:
		showUsage()

	}
}
