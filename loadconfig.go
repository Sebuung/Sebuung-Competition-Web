package main

import (
	"os"
	"encoding/json"
	"fmt"
)

type s3_config struct {
	BUCKET_NAME string `json: "BUCKET_NAME"`
	REGION string `json: "REGION"`
}

func LoadConfig() (s3_config, error){
	var config s3_config
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
	return config, err
}
