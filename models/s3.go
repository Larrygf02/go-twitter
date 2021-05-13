package models

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var S3 *S3Client

func init() {
	S3 = new(S3Client)
}

type S3Client struct {
	Region  string
	Session *session.Session
	Svc     *s3.S3
}

func (t *S3Client) NewSession(region string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Session = sess
	t.Svc = s3.New(t.Session)
}

func (t *S3Client) Upload(file multipart.File, myBucket string, keyName string) (bool, error) {
	uploader := s3manager.NewUploader(t.Session)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(keyName),
		Body:   file,
	})

	if err != nil {
		return false, err
	}
	fmt.Println(result)
	return true, nil
}
