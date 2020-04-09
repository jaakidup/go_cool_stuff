package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	jsonData, err := GetData("./data.json")
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	json.Unmarshal(jsonData, &result)

	// all users
	fmt.Println(result["users"])

	// Select users, [1] in array
	fmt.Println(result["users"].([]interface{})[1])

	// select users, 2 result, select social links
	fmt.Println(result["users"].([]interface{})[1].(map[string]interface{})["social"])
	fmt.Println(result["users"].([]interface{})[1].(map[string]interface{})["social"].(map[string]interface{})["facebook"])

	// Using MapData ...
	var otherway map[string]interface{}
	jsonObject, err := MapData(jsonData, otherway)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(jsonObject)
	fmt.Println(jsonObject.(map[string]interface{})["users"].([]interface{})[1].(map[string]interface{})["name"]) // Amy
}

// MapData takes a []byte and encodes it into dest interface{}
func MapData(data []byte, dest interface{}) (interface{}, error) {
	dataObject := dest
	err := json.Unmarshal(data, &dataObject)
	if err != nil {
		return nil, err
	}
	return dataObject, nil
}

// GetData reads the contents of fileName in memory
// returns []byte and error
func GetData(fileName string) ([]byte, error) {
	jsonData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read json data from %s", fileName)
	}
	return jsonData, nil
}
