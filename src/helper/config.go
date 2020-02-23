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
  "os"
)

// --- General config

type Config struct {
  Json JsonConfig
  AppPath string
}

// --- Json Config

type JsonConfig struct {
  Main Main `json:"main"`
  Index Index `json:"index"`
  Pubs Pubs `json:"pubs"`
}


// --- Main

type Main struct {
  AppUrl string `json:"app_url"`
  Dirs Dirs `json:"dirs"`
}

type Dirs struct {
  Bin string
  Drafts string
  Templates string
}


// --- Index

type Index struct {
  Default IndexDefault `json:"default"`  
}

type IndexDefault struct {
  PerPage int `json:"per_page"`
}


// --- Pubs

type Pubs struct {
  Default PubsDefault `json:"default"`
}

type PubsDefault struct {
  PerPage int `json:"per_page"`
  Langs []string `json:"langs"`
}

func LoadJsonConfig() (JsonConfig) {

  var config JsonConfig

  path, _ := os.Getwd()

  jsonFile, err := os.Open(path + "/etc/main.json")

  if err != nil {
    return config
  }

  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  json.Unmarshal(byteValue, &config)

  return config
}


func LoadConfig() (Config) {

  return Config{Json: LoadJsonConfig(), AppPath: os.Getenv("GOPATH")}
}