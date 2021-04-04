package hazardhub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	parsedBucket = os.Getenv("PARSED_BUCKET")
	rawBucket    = os.Getenv("RAW_BUCKET")
	s3Up         *s3manager.Uploader
	s3Down       *s3manager.Downloader
)

func getParsed(key string) (map[int]*RiskProfile, bool) {
	buf := aws.NewWriteAtBuffer([]byte{})

	_, err := s3Down.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(parsedBucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NoSuchKey" {
			return nil, false
		}
		tmplt := "unable to get parsed hazard hub profile(%s): %s\n"
		fmt.Printf(tmplt, key, err)
		return nil, false
	}

	var p map[int]*RiskProfile
	err = json.Unmarshal(buf.Bytes(), &p)
	if err != nil {
		tmplt := "unable to Unmarshal hazard hub profile(%s): %s\n"
		fmt.Printf(tmplt, key, err)
		return nil, false
	}

	return p, true
}

func getRaw(key string) (*RisksResponse, bool) {
	buf := aws.NewWriteAtBuffer([]byte{})

	_, err := s3Down.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(rawBucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NoSuchKey" {
			return nil, false
		}
		tmplt := "unable to get raw hazard hub profile(%s): %s\n"
		fmt.Printf(tmplt, key, err)
		return nil, false
	}

	var r RisksResponse
	err = json.Unmarshal(buf.Bytes(), &r)
	if err != nil {
		tmplt := "unable to parse hazard hub raw response(%s): %s\n"
		fmt.Printf(tmplt, key, err)
		return nil, false
	}

	return &r, true
}

func uploadParsed(key string, doc map[int]*RiskProfile) {
	b, _ := json.Marshal(doc)
	_, err := s3Up.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(parsedBucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(b),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		tmplt := "unable to upload parsed hazard hub doc(%s): %s, err: %s"
		fmt.Printf(tmplt, key, string(b), err)
	}
}

func uploadRaw(key string, b []byte) {
	_, err := s3Up.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(rawBucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(b),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		tmplt := "unable to upload raw hazard hub response(%s): %s, err: %s"
		fmt.Printf(tmplt, key, string(b), err)
	}
}
