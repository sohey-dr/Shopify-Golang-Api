package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
    "time"

		"github.com/joho/godotenv"
    goshopify "github.com/bold-commerce/go-shopify"
)

// Create an app somewhere.
app := goshopify.App{
	ApiKey: "abcd",
	ApiSecret: "efgh",
	RedirectUrl: "https://example.com/shopify/callback",
	Scope: "read_products,read_orders",
}

b, err := ioutil.ReadAll(r.Body)//リクエストのbodyをパース
defer r.Body.Close()
if err != nil {
    panic(err)
}
var jsonData Webhook//jsonの形式に合わせたstructを初期化
if err := json.Unmarshal(b, &jsonData); err != nil {//jsonをパースして構造体に値を入れる。
    panic(err)
}
//Webgook Webhookの構造体
type Webhook = struct {
    ID              json.Number `json:id`//これがID。受注番号とは違う。
    Gateway         string      `json:gateway`
    FinancialStatus string      `json:financial_status`//支払いステータス 読み取り専用
}