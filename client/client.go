package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"io/ioutil"
	//"strconv"
)

const (
	config_file = "dump/dump.json"
)

type Request struct {
	Name string `json:"name"`
	Nr   int    `json:"nr"`
}

type ReliefWeb struct {
	AppName     string    `json:"appname"`
	MaxRequests int       `json:"limit"`
	Profile     string    `json:"profile"`
	Preset      string    `json:"preset"`
	Requests    []Request `json:"requests"`
	Sort        string    `json:"sort"`
}

type RequestResponse struct {
	TotalCount int `json:"totalCount"`
	Count      int `json:"count"`
}

func main() {
	/*
	   init reliefweb params
	*/
	reliefweb := &ReliefWeb{
		AppName:     "sokrates.ai", // "apidoc" is the default
		MaxRequests: 2,             // maximum of relief web is 1000
		Profile:     "full",        // full , minimal(default), list
		Preset:      "analysis",    // minimal, analysis, latest
		Requests: []Request{
			{"reports", 0},
			{"disasters", 0},
			{"countries", 0},
			{"jobs", 0},
			{"training", 0},
			{"sources", 0},
		},
		Sort: "asc", // asc , desc

	}

	fmt.Printf("hallo %s\n", "world")
	logger := log.New(os.Stdout, "sk_reliefweb_client: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Print("hallo")

	client := &http.Client{}

	request_url := fmt.Sprintf("http://api.reliefweb.int/v1/%s?sort[]=id:%s&appname=%s&profile=%s&preset=%s&offset=0&limit=%d", "reports", reliefweb.Sort, reliefweb.AppName, reliefweb.Profile, reliefweb.Preset, reliefweb.MaxRequests)
	logger.Printf("request url => %s", request_url)

	resp, err := client.Get(request_url)
	if err != nil {
		logger.Printf("ERROR get : %s", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("response status : %s\n", resp.Status)
	fmt.Printf("response code : %d\n", resp.StatusCode)

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		logger.Printf("ERROR reading body : %s", err)
		return
	}

	fmt.Println("*******************************************************")
	fmt.Println(buf.String())
	fmt.Println("*******************************************************")
	fmt.Println(buf.Bytes())
	fmt.Printf("\n\n\n")

	var f interface{}
	//err = json.Unmarshal(buf.Bytes(), &f)
	//t, _ := strconv.Unquote(buf.String())
	b := []byte(buf.String())
	err = json.Unmarshal(b, &f)
	if err != nil {
		logger.Printf("ERROR : ", err)
		os.Exit(1)
	}

	m := f.(map[string]interface{})
	for k, v := range m {

		switch vv := v.(type) {
		//switch v.(type) {
		case string:
			//fmt.Printf(k, " is string ", vv)
			fmt.Printf("%s is a string: %s\n\n", k, vv)
		case int:
			//fmt.Println(k, " is int ", vv)
			fmt.Printf("%s is a int: %d\n\n", k, vv)
		case float64:
			fmt.Printf("%s is a float64 : %f\n\n", k, vv)
		case bool:
			fmt.Printf("%s is a bool: %t\n\n", k)
		case nil:
			fmt.Printf("%s is a nil\n\n", k)
		case []interface{}:
			fmt.Println(k, " is an array: ")
			for i, u := range vv {
				fmt.Printf("\t%d : %v\n\n", i, u)
			}
		case map[string]interface{}:
			fmt.Println(k, " is an map")
			for x , y := range vv {
				fmt.Printf("\t%s: %v\n\n", x, y)
			}
		default:
			fmt.Println(k, " is of a type i don't know how to handle")
			fmt.Printf("default formate : %v, Go-syntax representation : %#v, Go-syntax representation of the type of the value : %T\n", k, k, k)

		}
	}
	
	
	// format the json in nice readable format
	jsonString , err := json.MarshalIndent(m, "", "   ")
	if err != nil {
		fmt.Errorf("Error marshall configuration: %s\n\n", err)
		return 
	}
	
	fmt.Printf("\n\n*********\n%s\n********\n\n", jsonString)
	
	//write to disk
	if err := os.MkdirAll(filepath.Dir(config_file), 0755); err != nil {
		fmt.Errorf("Error creating directory %s : \n%s\n\n", filepath.Dir(config_file), err)
		return 
	}
	
	if err := ioutil.WriteFile(config_file, jsonString, 0755); err != nil {
		fmt.Errorf("error writing config file : %s\n\n", err)
		return 
	}

}
