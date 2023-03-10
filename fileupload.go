package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"context"
	"bytes"
)

func fileUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	f, err := os.OpenFile("./files/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	
	client := s3.NewFromConfig(cfg)
	fileByte, err := ioutil.ReadFile("./files/" + handler.Filename)
	filebody := bytes.NewReader(fileByte)
	uploader := manager.NewUploader(client)
	config, err := LoadConfig()
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(config.BUCKET_NAME),
		Key:    aws.String(handler.Filename),
		Body:   filebody,
	})
	if err != nil{
		fmt.Println(err)
		fmt.Println(result)
	}

	fmt.Fprintf(w, `<script>
	alert('%v이 성공적으로 제출되셨습니다, 수고하셨습니다!')
	location.href = '/'
	</script>`, handler.Filename)
}
