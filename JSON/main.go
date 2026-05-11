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
	// EncodeJson()
	DecodeJSON()
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

func DecodeJSON() {
	jsonData := []byte(`
		{
                "coursename": "Android Dev",
                "courseprice": 1000,
                "website": "LCO",
                "tags": [
                        "Android",
                        "mobile dev"
				]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("json is valid")

		json.Unmarshal(jsonData, &lcoCourse)

		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Invalid JSON")
	}

	var myOnlineData map[string]interface{}

	if checkValid {
		fmt.Println("json is valid")

		json.Unmarshal(jsonData, &myOnlineData)

		fmt.Printf("%#v\n", myOnlineData["coursename"])
	} else {
		fmt.Println("Invalid JSON")
	}

}
