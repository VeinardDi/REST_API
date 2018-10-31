package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"net/http"
)

type M map[string]interface{}
type response struct {
	Rates []M `json:"Rates"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	bodyMap := map[string]interface{}{}

	json.Unmarshal(body, &bodyMap)

	valutesMap, ok := bodyMap["Valute"].(map[string]interface{})
	if !ok {
		fmt.Println("Error while parsing valutes")
		return
	}
	fmt.Println(valutesMap["USD"])
	fmt.Println(valutesMap["EUR"])

	USDMap, ok := valutesMap["USD"].(map[string]interface{})
	if !ok {
		fmt.Println("Error while parsing USD")
		return
	}

	EURMap, ok := valutesMap["EUR"].(map[string]interface{})
	if !ok {
		fmt.Println("Error while parsing EUR")
		return
	}


	fmt.Println(USDMap["Value"])
	fmt.Println(EURMap["Value"])

	//usdValue := USDMap["Value"].(float64)
	//eurValue := EURMap["Value"].(float64)

	var MapSlice []M

	m1 := M{"From": "Рубль", "To": "Доллар США", "Value": USDMap["Value"]}
	m2 := M{"From": "Рубль", "To": "Евро", "Value": EURMap["Value"]}

	MapSlice = append(MapSlice, m1, m2)

	res := &response{
		Rates: MapSlice}
	resed, _ := json.MarshalIndent(res, "", "    ")
	w.Write(resed)

}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)

}
