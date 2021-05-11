# Akamai OPEN EdgeGrid for GoLang v1

[![Build Status](https://travis-ci.org/akamai/AkamaiOPEN-edgegrid-golang.svg?branch=master)](https://travis-ci.org/akamai/AkamaiOPEN-edgegrid-golang)
[![GoDoc](https://godoc.org/github.com/akamai/AkamaiOPEN-edgegrid-golang?status.svg)](https://godoc.org/github.com/akamai/AkamaiOPEN-edgegrid-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/akamai/AkamaiOPEN-edgegrid-golang)](https://goreportcard.com/report/github.com/akamai/AkamaiOPEN-edgegrid-golang)
[![License](http://img.shields.io/:license-apache-blue.svg)](https://github.com/akamai/AkamaiOPEN-edgegrid-golang/blob/master/LICENSE)

This library implements an Authentication handler for [net/http](https://golang.org/pkg/net/http/)
that provides the [Akamai OPEN Edgegrid Authentication](https://developer.akamai.com/introduction/Client_Auth.html) 
scheme. For more information visit the [Akamai OPEN Developer Community](https://developer.akamai.com).  This library 
has been released as a v1 library though future development will be on the v2 branch

## Announcing Akamai OPEN EdgeGrid for GoLang v2 (release v2.0.0)

The v2 branch of this module is under active development and provides a subset of Akamai APIs for use in the 
Akamai Terraform Provider. The v2 branch **does not yet** implement the full set of Akamai endpoints supported by the 
0.x and 1.x releases.  

New users are encouraged to adopt v2 branch it is a simpler API wrapper with little to no business logic.

Current direct users of this v0.9 library are recommended to continue to use the the v1 version as initialization 
and package structure has significantly changed in v2 and will require substantial work to migrate existing 
applications. Non-backwards compatible changes were made to improve the code quality and make the project more 
maintainable. 

## Usage of the v1 library

GET Example:

```go
  package main

import (
	"fmt"
	"io/ioutil"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

func main() {
	config, _ := edgegrid.Init("~/.edgerc", "default")

	// Retrieve all locations for diagnostic tools
	req, _ := client.NewRequest(config, "GET", "/diagnostic-tools/v2/ghost-locations/available", nil)
	resp, _ := client.Do(config, req)

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byt))
}
```

Parameter Example:

```go
  package main

import (
	"fmt"
	"io/ioutil"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

func main() {
	config, _ := edgegrid.Init("~/.edgerc", "default")

	// Retrieve dig information for specified location
	req, _ := client.NewRequest(config, "GET", "/diagnostic-tools/v2/ghost-locations/zurich-switzerland/dig-info", nil)

	q := req.URL.Query()
	q.Add("hostName", "developer.akamai.com")
	q.Add("queryType", "A")

	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(config, req)

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byt))
}
```

POST Example:

```go
  package main

  import (
    "fmt"
    "io/ioutil"
    "net/http"
    
    "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
    "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
  )

  func main() {
    config, _ := edgegrid.Init("~/.edgerc", "default")
    
    // Acknowledge a map
    req, _ := client.NewRequest(config, "POST", "/siteshield/v1/maps/1/acknowledge", nil)
    resp, _ := client.Do(config, req)

    defer resp.Body.Close()
    byt, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(byt))
  }
```
POST Example with body:

```go
  package main

  import (
    "fmt"
    "io/ioutil"
    "net/http"
    
    "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
    "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
  )

  //Create a stream
  func main() {
    config, _ := edgegrid.Init("~/.edgerc", "default")
    
    strm := &stream{
    			Name:              "TestName123",
    			ContractId:        "P-1234567",
    			Format:            "HLS",
    			Cpcode:            987654,
    			IngestAccelerated: false,
    			AllowedIps:        []string{"1.2.3.4"},
    			EncoderZone:       "US_EAST",
    			BackupEncoderZone: "AUSTRALIA",
    			Origin: origin{
    				HostName: "testhost9124",
    			},
    			AdditionalEmailIds:          []string{"xyz@akamai1.com"},
    			IsDedicatedOrigin:           false,
    			ActiveArchiveDurationInDays: 1,
    		}
    
    		req, _ := client.NewJSONRequest(config, "POST", "/config-media-live/v2/msl-origin/streams", *strm)
    	
    		resp, _ := client.Do(config, req)
    	
    		defer resp.Body.Close()

    		for k, v := range resp.Header {
    			fmt.Print(k)
    			fmt.Print(" : ")
    			fmt.Println(v)
    		}
  }


type stream struct {
	Id                          int      `json:"id,omitempty"`
	Name                        string   `json:"name,omitempty"`
	ContractId                  string   `json:"contractId,omitempty"`
	Format                      string   `json:"format,omitempty"`
	Type                        string   `json:"type,omitempty"`
	Cpcode                      int      `json:"cpcode,omitempty"`
	IngestAccelerated           bool     `json:"ingestAccelerated"`
	AllowedIps                  []string `json:"allowedIps,omitempty"`
	EncoderZone                 string   `json:"encoderZone,omitempty"`
	BackupEncoderZone           string   `json:"backupEncoderZone,omitempty"`
	Origin                      origin   `json:"origin,omitempty"`
	PrimaryPublishingUrl        string   `json:"primaryPublishingUrl,omitempty"`
	BackupPublishingUrl         string   `json:"backupPublishingUrl,omitempty"`
	AdditionalEmailIds          []string `json:"additionalEmailIds,omitempty"`
	IsDedicatedOrigin           bool     `json:"isDedicatedOrigin"`
	ActiveArchiveDurationInDays int      `json:"activeArchiveDurationInDays,omitempty"`
}


type origin struct {
	Id          int    `json:"id,omitempty"`
	Cpcode      int    `json:"cpcode,omitempty"`
	HostName    string `json:"hostName,omitempty"`
	EncoderZone string `json:"encoderZone,omitempty"`
	ContractId  string `json:"contractId,omitempty"`
}
```

PUT Example:

```go
  package main

  import (
    "fmt"
    "io/ioutil"
    "net/http"
    
    "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
    "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
  )

  func main() {

    config, _ := edgegrid.Init("~/.edgerc", "default")
    body := []byte("{\n  \"name\": \"Simple List\",\n  \"type\": \"IP\",\n  \"unique-id\": \"345_BOTLIST\",\n  \"list\": [\n    \"192.168.0.1\",\n    \"192.168.0.2\",\n  ],\n  \"sync-point\": 0\n}")
    
    // Update a Network List
    req, _ := client.NewJSONRequest(config, "PUT", "/network-list/v1/network_lists/unique-id?extended=extended", body)
    resp, _ := client.Do(config, req)

    defer resp.Body.Close()
    byt, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(byt))
  }
```

Alternatively, your program can read it from config struct.

```go
  package main

  import (
    "fmt"
    "io/ioutil"
    "net/http"
    
    "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
    "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
  )

  func main() {
    config := edgegrid.Config{
      Host : "xxxxxx.luna.akamaiapis.net",
      ClientToken:  "xxxx-xxxxxxxxxxx-xxxxxxxxxxx",
      ClientSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      AccessToken:  "xxxx-xxxxxxxxxxx-xxxxxxxxxxx",
      MaxBody:      1024,
      HeaderToSign: []string{
        "X-Test1",
        "X-Test2",
        "X-Test3",
      },
      Debug:        false,
    }
    
   // Retrieve all locations for diagnostic tools
	req, _ := client.NewRequest(config, "GET", fmt.Sprintf("https://%s/diagnostic-tools/v2/ghost-locations/available",config.Host), nil)
	resp, _ := client.Do(config, req)

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byt))
}
```

## Contribute

1. Fork [the repository](https://github.com/akamai/AkamaiOPEN-edgegrid-golang) to start making your changes to the **master** branch
2. Send a pull request.

## Author

[Davey Shafik](mailto:dshafik@akamai.com) - Developer Evangelist @ [Akamai Technologies](https://developer.akamai.com)
[Nick Juettner](mailto:hello@juni.io) - Software Engineer @ [Zalando SE](https://tech.zalando.com/)  

