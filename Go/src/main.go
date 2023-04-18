package main 


import (
    "fmt"
    "net/http"
    "errors"
    "io/ioutil"
    )


func main () {
    var url string
    fmt.Println("Type the PixelBattle API URL")
    fmt.Scanln(&url)
    data, err :=  GET(url)
    if err != nil {
        errors.New("FUCKING FATAL ERROR")
    } else {
        fmt.Println(data)
    }
}

func GET(link string) (string, error) {
	result, err := http.Get(link)
	if err != nil {
	    return " ", errors.New("FATAL ERROR")
	}
	defer result.Body.Close()
	
	body, err := ioutil.ReadAll(result.Body)
	if err != err {
	    return " ", err
	}
	return string(body), err
}
