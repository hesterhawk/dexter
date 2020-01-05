package helper

/*
  JSON Configuration

  ---
  1. Sample (Go):

  type Config struct {
    PubsPerPage    int `json:"pubs_per_page"`
    Pub            Pub `json:"pub"`
  }

  type Pub struct {
    Resources []Resource `json:"resources"`
  }

  type Resource struct {
    Type string `json:"type"`
    Url  string `json:"url"`
  }


  ---
  2. Sample (json)

  {
    "pubs_per_page": 10,
    "pub": {
      "resources": [
        {
          "type": "js",
          "url": "http:/google.com"
        },
        {
          "type": "css",
          "url": "http:/google.com"
        }
      ]
    }
  }
*/

import (
  "encoding/json"
  "io/ioutil"
  "fmt"
  "os"
)

// --- General config

type Config struct {
  Pub Pub `json:"pub"`
  Pubs Pubs `json:"pubs"`
}


// --- Pub

type Pub struct {
  Resources []Resource `json:"resources"`
}

type Resource struct {
  Type string `json:"type"`
  Url  string `json:"url"`
}


// --- Pubs

type Pubs struct {
  Default Default `json:"default"`
}

type Default struct {
  PerPage int `json:"per_page"`
  Langs []string `json:"langs"`
}


func LoadConfig() (Config) {

  var config Config

  path, _ := os.Getwd()

  jsonFile, err := os.Open(path + "/config.json")

  if err != nil {
    fmt.Println(err)

    return config
  }

  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  json.Unmarshal(byteValue, &config)

  return config
}
