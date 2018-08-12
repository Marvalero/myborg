package secret

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"github.com/Marvalero/myborg/myborg/gcp"
	"github.com/spf13/viper"
)

func Encrypt(fileName string, text string) {
	fmt.Println("Create file:", fileName, "with content:", text)
	ctx := context.Background()
	encryptedContent, err := gcp.EncryptText(text, ctx)
	if err != nil {
		log.Fatalf("Failed to encrypt content: %v", err)
	}
	createFile(fileName, encryptedContent)

}

func findBucket(bucketName string, ctx context.Context) *storage.BucketHandle {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	bucket := client.Bucket(bucketName)
	if err != nil {
		log.Fatalf("Failed to fetch bucket: %v", err)
	}
	return bucket

}

func createObject(bucket *storage.BucketHandle, filename string, ctx context.Context) *storage.ObjectHandle {
	obj := bucket.Object(filename)
	if _, err := obj.Attrs(ctx); err == nil {
		log.Fatal("Already existent file: ", filename)
	}
	return obj
}

func createFile(fileName string, content string) {
	ctx := context.Background()
	bucket := findBucket(viper.Get("secrets-bucket").(string), ctx)
	obj := createObject(bucket, fileName, ctx)
	wc := obj.NewWriter(ctx)
	wc.ContentType = "text/plain"
	if _, err := wc.Write([]byte(content)); err != nil {
		log.Fatal("createFile: unable to write file %q: %v", fileName, err)
	}
	if err := wc.Close(); err != nil {
		log.Fatal("createFile: unable to close file %q: %v", fileName, err)
	}
}
