package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-playground/validator/v10"
	"github.com/phial3/go-notes/cloudstation/store"
)

var (
	validate = validator.New()
)

func NewUploader(endpt, accessID, accessKey string) (store.Uploader, error) {
	p := &aliyun{
		Endpoint:  endpt,
		AccessID:  accessID,
		AccessKey: accessKey,

		listener: NewOssProgressListener(),
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return p, nil
}

type aliyun struct {
	Endpoint  string `validate:"required,url"`
	AccessID  string `validate:"required"`
	AccessKey string `validate:"required"`

	listener oss.ProgressListener
}

func (p *aliyun) validate() error {
	return validate.Struct(p)
}

func (p *aliyun) UploadFile(bucketName, objectKey, localFilePath string) error {
	bucket, err := p.GetBucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(objectKey, localFilePath, oss.Progress(p.listener))
	if err != nil {
		return err
	}

	signedURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return fmt.Errorf("SignURL error, %s", err)
	}
	fmt.Printf("下载链接: %s\n", signedURL)
	fmt.Println("\n注意: 文件下载有效期为1天, 中转站保存时间为3天, 请及时下载")

	return nil
}

func (p *aliyun) GetBucket(bucketName string) (*oss.Bucket, error) {
	if bucketName == "" {
		return nil, fmt.Errorf("upload bucket name required")
	}

	// New client
	client, err := oss.New(p.Endpoint, p.AccessID, p.AccessKey)
	if err != nil {
		return nil, err
	}

	// Get bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
