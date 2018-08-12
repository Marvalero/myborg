package secret

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/google"
	cloudkms "google.golang.org/api/cloudkms/v1"
)

func Encrypt(name string, credentials string) {
	fmt.Println("Create file:", name, "with content:", credentials)
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	client.Bucket(viper.Get("secrets-bucket").(string))

	oauthClient, err := google.DefaultClient(ctx, cloudkms.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	_, err = cloudkms.New(oauthClient)
	if err != nil {
		log.Fatal(err)
	}

}
