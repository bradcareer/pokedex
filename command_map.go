package main
import (
    "fmt"
    "log"
    "net/http"
	"io"
	"encoding/json"
)



func commandMap(cfg *config) error {
	var url string
	
	if cfg.next == ""{
		url = "https://pokeapi.co/api/v2/location-area"
	} else{
		url = cfg.next
	}
		
	
	res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code: %d", res.StatusCode)
		}
		if err != nil {
			return err
		}
		location := locations{}
		err = json.Unmarshal(body, &location)
		if err != nil{
			fmt.Println(err)
			return err
		}
		for _, loc := range location.Results {
			fmt.Println(loc.Name)

		}
	cfg.next = location.Next
	//fmt.Println("DEBUG new cfg.next:", cfg.next)
	if location.Previous !=nil {
		cfg.previous = location.Previous
	}
    return nil
}