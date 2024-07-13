# request demo

### 一、get request 
```golang
var requestUrl = "https://movie.douban.com/j/search_tags?type=movie&source=index"
var header = map[string]string{
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
}
client := http.GetRequest{
    RequestUrl:         requestUrl,
    Header:             header,
    InsecureSkipVerify: false, //skip TLS/SSL  true or false default:false
    TimeOut:            2,     //timeout s default:2
}
response := client.Get()
if response.Error == nil && response.HttpCode == 200 {
    type RStruct struct {
        Tags []string `json:"tags"`
    }
    rStruct := RStruct{}
    err := json.Unmarshal([]byte(response.Content), &rStruct)
    if err == nil {
        for k, v := range rStruct.Tags {
            sprintf := fmt.Sprintf("tag_id: %d, tag_name: %s", k, v)
            fmt.Println(sprintf)
        }
    }
}
```