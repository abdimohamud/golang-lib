package hazardhub

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var required = map[string]string{
	"ENVIRONMENT":   env,
	"PARSED_BUCKET": parsedBucket,
	"RAW_BUCKET":    rawBucket,
}

func init() {
	for k, v := range required {
		if len(v) == 0 {
			panic(fmt.Sprintf("%s env var is required", k))
		}
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	s3Up = s3manager.NewUploader(sess)
	s3Down = s3manager.NewDownloader(sess)
}
