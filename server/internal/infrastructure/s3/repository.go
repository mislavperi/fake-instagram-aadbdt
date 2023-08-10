package repository

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Repository struct {
	svc    *s3.S3
	bucket string
	region string
}

func NewS3Repository(bucket string, region string, id string, secret string, token string) *S3Repository {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(id, secret, token),
	})
	if err != nil {
		panic(err)
	}

	return &S3Repository{
		svc:    s3.New(sess),
		bucket: bucket,
	}
}

func (r *S3Repository) UploadToBucket(file multipart.File) (*string, error) {
	key := "title.jpg"
	var timeout time.Duration

	ctx := context.Background()

	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}

	if cancelFn != nil {
		defer cancelFn()
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}

	_, err := r.svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			return nil, fmt.Errorf("upload canceled due to timeout, %v", err)
		} else {
			return nil, fmt.Errorf("failed to upload object, %v", err)
		}
	}
	url := fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", r.bucket, r.region, key)
	return &url, nil
}
