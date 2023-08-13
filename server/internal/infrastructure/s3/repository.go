package repository

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math"
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
		region: region,
		bucket: bucket,
	}
}

func (r *S3Repository) UploadToBucket(file bytes.Buffer, fileExt string) (*string, error) {
	buff := make([]byte, int(math.Ceil(float64(16)/float64(1.33333333333))))
	rand.Read(buff)
	key := base64.RawURLEncoding.EncodeToString(buff)
	var timeout time.Duration

	ctx := context.Background()

	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}

	if cancelFn != nil {
		defer cancelFn()
	}

	_, err := r.svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(fmt.Sprintf("%s.%s", key, fileExt)),
		Body:   bytes.NewReader(file.Bytes()),
	})
	if err != nil {
		if err, ok := err.(awserr.Error); ok && err.Code() == request.CanceledErrorCode {
			return nil, fmt.Errorf("upload canceled due to timeout, %v", err)
		} else {
			return nil, fmt.Errorf("failed to upload object, %v", err)
		}
	}
	url := fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s.%s", r.bucket, r.region, key, fileExt)
	return &url, nil
}
