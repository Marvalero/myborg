package secret

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
)

func List() {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	bucket := client.Bucket(viper.Get("secrets-bucket").(string))

	it := bucket.Objects(ctx, nil)
	fmt.Println("Files:")
	for {
		fileAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fileAttrs.Name)
	}
}
