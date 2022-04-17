package pkg

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func GetObject() {
	bucket, err := GetBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	err = bucket.GetObjectToFile("my-object", "LocalFile")
	if err != nil {
		// HandleError(err)
	}
	fmt.Println("OSS Go SDK Version: ", oss.Version)

	err = bucket.DownloadFile(objectKey, "mynewfile-3.jpg", 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
	if err != nil {
		HandleError(err)
	}

}
