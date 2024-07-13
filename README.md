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

### 二、post request（Content-Type：application/x-www-form-urlencoded）
```golang
var requestUrl = "https://api-article.huxiu.com/channel/channelList"
var header = map[string]string{
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
}
client := http.PostRequest{
    RequestUrl: requestUrl,
    Header:     header,
    FormData: map[string]string{
        "platform": "m",
    },
    BodyContent:        "",
    InsecureSkipVerify: false, //skip TLS/SSL  true or false default:false
    TimeOut:            2,     //timeout s default:2
}
response := client.Post()

if response.Error == nil && response.HttpCode == 200 {
    type RStruct struct {
        Data []struct {
            Name string `json:"name"`
        } `json:"data"`
    }
    rStruct := RStruct{}
    err := json.Unmarshal([]byte(response.Content), &rStruct)
    if err == nil {
        for _, v := range rStruct.Data {
            sprintf := fmt.Sprintf("name: %s", v.Name)
            fmt.Println(sprintf)
        }
    }
}
```

### 三、post request（Content-Type：application/json）
```golang
var requestUrl = "https://api.juejin.cn/recommend_api/v1/article/detail_rela_rec"
var header = map[string]string{
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
}
client := http.PostRequest{
    RequestUrl:         requestUrl,
    Header:             header,
    BodyContent:        `{"id_type":2,"cursor":"0","item_id":"7016742808560074783","sort_type":200,"limit":5,"referer":""}`,
}
response := client.Post()
if response.Error == nil && response.HttpCode == 200 {
    type RStruct struct {
        Data []struct {
            ArticleId string `json:"article_id"`
        } `json:"data"`
    }
    rStruct := RStruct{}
    err := json.Unmarshal([]byte(response.Content), &rStruct)
    if err == nil {
        for _, v := range rStruct.Data {
            sprintf := fmt.Sprintf("articleId: %s", v.ArticleId)
            fmt.Println(sprintf)
        }
    }
}
```