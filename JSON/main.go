package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"courseprice"`
	Platform string   `json:"website"`
	password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to JSON video")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{
			Name:     "Android Dev",
			Price:    1000,
			Platform: "LCO",
			password: "1221",
			Tags:     []string{"Android", "mobile dev"},
		},
		{
			Name:     "IOS Dev",
			Price:    1000,
			Platform: "LCO",
			password: "1221",
			Tags:     []string{"IOS", "mobile dev"},
		},
		{
			Name:     "WEB Dev",
			Price:    1000,
			Platform: "LCO",
			password: "1221",
			Tags:     nil,
		},
	}
	// finalJSON, err := json.Marshal(lcoCourses)
	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")

	// var jsonObject map[string]any
	// json.Unmarshal(finalJSON, &jsonObject)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJSON))

}
