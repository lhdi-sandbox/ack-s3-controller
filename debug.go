package main

import (
	"context"
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
		log.Fatalf("ListBuckets error: %s", err.Error())
	}

	fmt.Println("BUCKETS:")
	fmt.Println(response.GoString())
	fmt.Println("---------")

	context := context.TODO()
	for i := 0; i < len(response.Buckets); i++ {
		fmt.Println(fmt.Sprintf("\nBucket: %s", *response.Buckets[i].Name))
		input := &s3.GetBucketAccelerateConfigurationInput{
			Bucket: response.Buckets[i].Name,
		}

		resp, err := s3Client.GetBucketAccelerateConfigurationWithContext(context, input)
		if err != nil {
			fmt.Println(fmt.Sprintf("GetBucketAccelerateConfigurationWithContext ERROR: %s", err.Error()))
		} else {
			fmt.Println("GetBucketAccelerateConfigurationWithContext RESULT", resp.GoString())
		}

		resp, err = s3Client.GetBucketAccelerateConfiguration(input)
		if err != nil {
			fmt.Println(fmt.Sprintf("GetBucketAccelerateConfiguration ERROR: %s", err.Error()))
		} else {
			fmt.Println("GetBucketAccelerateConfiguration RESULT", resp.GoString())
		}

		fmt.Println("\n---------")
	}
}
