package main

import (
    "fmt"
    "io/ioutil"
    "log"
    // "os"
    "net/http"
    "net/url"
    "strings"
    // "strconv"
)

func main() {
    apiUrl := "{BASEURL}"
    resource := "/api/views/create"
    username := "{USERNAME}"
    password := "{PASSWORD}"
    // name := "tes handling"

    list, err := ioutil.ReadFile("list.txt")
    if err != nil {
        log.Fatal(err)
    }

    lines := strings.Split(string(list), "\n")
    for _, name := range lines {

        data := url.Values{}
        key := name

        //if the name of the application contains "("
        if strings.Contains(name, "("){
            before, _, _ := strings.Cut(name, "(")
            key = before
        }
        data.Set("key", strings.ReplaceAll(key, " ", "_")+":key")
        data.Add("name", name)
        data.Add("qualifier", "VM")
    
        u, _ := url.ParseRequestURI(apiUrl)
        u.Path = resource
        urlStr := u.String()
    
        client := &http.Client{}
        r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
        r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
        r.SetBasicAuth(username, password)
    
        resp, err := client.Do(r)    
        if err != nil {
            fmt.Println(name)
            log.Fatal(err)
            break
        }
        if strings.Contains(resp.Status, "200"){
            fmt.Println("Success "+name)
        }else{
            fmt.Println("Error "+name+" ")
            fmt.Println(resp)
        }
    }
}