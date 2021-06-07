package client

import (
	"base/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Consumer() {
	for k := 0; k < 10; k++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("adsklfjladskjflasdjlaf")
		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://localhost:8080/userall", nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Print(err.Error())
		}
		defer resp.Body.Close()
		// efer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err.Error())
		}
		var responseObject []model.Users
		json.Unmarshal(bodyBytes, &responseObject)
		fmt.Printf("API Response as struct %+v\n", responseObject)
	}
}
