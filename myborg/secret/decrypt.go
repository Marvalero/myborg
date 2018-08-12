package secret

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
	"github.com/Marvalero/myborg/myborg/gcp"
	"github.com/spf13/viper"
)

func Decrypt(fileName string) {
	fmt.Println("Decrypt file:", fileName)
	ctx := context.Background()
	text := fileContent(fileName, ctx)
	decryptedContent, err := gcp.DecryptText(text, ctx)
	check(err, "Error decrypting text")
	fmt.Println("Secret:", decryptedContent)

}

func fetchObject(bucket *storage.BucketHandle, filename string, ctx context.Context) *storage.ObjectHandle {
	obj := bucket.Object(filename)
	if _, err := obj.Attrs(ctx); err != nil {
		log.Fatal("Already existent file: ", filename)
	}
	return obj
}

func fileContent(fileName string, ctx context.Context) []byte {
	bucket := findBucket(viper.Get("secrets-bucket").(string), ctx)
	obj := fetchObject(bucket, fileName, ctx)
	rc, err := obj.NewReader(ctx)
	check(err, "Error opening file")
	defer rc.Close()
	slurp, err := ioutil.ReadAll(rc)
	check(err, "Error reading file")

	return slurp
}

func check(err error, message string) {
	if err != nil {
		log.Fatalf("%v: %v", message, err)
	}
}
