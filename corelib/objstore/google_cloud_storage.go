package objstore

import (
	"bytes"
	"context"
	"io"

	"cloud.google.com/go/storage"
)

type GoogleCloudStorage struct {
	storageClient *storage.Client
}

func NewGoogleCloudStorage(storageClient *storage.Client) *GoogleCloudStorage {
	return &GoogleCloudStorage{
		storageClient: storageClient,
	}
}

func (o GoogleCloudStorage) Upload(ctx context.Context, src []byte, bucket, remoteObjectName, srcContentType string) error {

	writer := o.storageClient.Bucket(bucket).Object(remoteObjectName).NewWriter(ctx)

	writer.ContentType = "application/octet-stream"
	if srcContentType != "" {
		writer.ContentType = srcContentType
	}

	byteReader := bytes.NewReader(src)
	if _, err := io.Copy(writer, byteReader); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	return nil
}

func (o GoogleCloudStorage) DownloadIntoMemory(ctx context.Context, bucket, remoteObjectName string) ([]byte, error) {

	rc, err := o.storageClient.Bucket(bucket).Object(remoteObjectName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	return data, nil
}
