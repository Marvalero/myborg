package gcp

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/oauth2/google"
	cloudkms "google.golang.org/api/cloudkms/v1"
)

func EncryptText(plaintext string, ctx context.Context) (string, error) {
	oauthClient, err := google.DefaultClient(ctx, cloudkms.CloudPlatformScope)
	if err != nil {
		return "", err
	}

	cloudkmsService, err := cloudkms.New(oauthClient)
	if err != nil {
		return "", err
	}

	parentName := fmt.Sprintf(
		"projects/%s/locations/global/keyRings/%s/cryptoKeys/%s",
		viper.Get("project-id").(string),
		viper.Get("secrets-key-ring").(string),
		viper.Get("secrets-key").(string))

	req := &cloudkms.EncryptRequest{
		Plaintext: base64.StdEncoding.EncodeToString([]byte(plaintext)),
	}
	resp, err := cloudkmsService.Projects.Locations.KeyRings.CryptoKeys.Encrypt(parentName, req).Do()
	if err != nil {
		return "", err
	}

	encrypted, err := base64.StdEncoding.DecodeString(resp.Ciphertext)
	return string(encrypted), err

}

func DecryptText(text []byte, ctx context.Context) (string, error) {
	oauthClient, err := google.DefaultClient(ctx, cloudkms.CloudPlatformScope)
	if err != nil {
		return "", err
	}

	cloudkmsService, err := cloudkms.New(oauthClient)
	if err != nil {
		return "", err
	}

	parentName := fmt.Sprintf(
		"projects/%s/locations/global/keyRings/%s/cryptoKeys/%s",
		viper.Get("project-id").(string),
		viper.Get("secrets-key-ring").(string),
		viper.Get("secrets-key").(string))

	req := &cloudkms.DecryptRequest{
		Ciphertext: base64.StdEncoding.EncodeToString(text),
	}
	resp, err := cloudkmsService.Projects.Locations.KeyRings.CryptoKeys.Decrypt(parentName, req).Do()
	if err != nil {
		return "", err
	}

	decrypted, err := base64.StdEncoding.DecodeString(resp.Plaintext)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil

}
