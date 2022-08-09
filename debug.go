package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	awsSession := session.Must(session.NewSession())
	s3Client := s3.New(awsSession)

	bucketQuery := &s3.ListBucketsInput{}
	response, err := s3Client.ListBuckets(bucketQuery)

	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	fmt.Println(response.GoString())
}
