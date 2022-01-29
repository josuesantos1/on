package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func request(file string, url string) {

}

func Upload(title string, file string) {

	if file == "" {
		fmt.Println("No file specified")
		return
	}

	if title == "" {
		var spliteded = strings.Split(file, "/")
		title = spliteded[len(spliteded)-1]
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-2"),
		},
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	svc := s3.New(sess)

	resp, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(),
		Key:    aws.String(title),
	})

	url, err := resp.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("error presigning request", err)
		return
	}

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("PUT", url, strings.NewReader(string(content)))
	if err != nil {
		fmt.Println("error creating request", url)
		return
	}

	result, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("error sending request", err)
		return
	}

	if result.StatusCode != 200 {
		data, _ := ioutil.ReadAll(result.Body)
		fmt.Println("[{+}] error sending request", result.Status)
		fmt.Println("[{+}] Body: \n", string(data))
		return
	}

	reqGet, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(),
		Key:    aws.String(title),
	})

	urlGet, err := reqGet.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	fmt.Println("[+] Status:", result.Status)
	fmt.Println("[+] look:", urlGet)
}
