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

    goshopify "github.com/tao-s/go-shopify"
)

// Create an app somewhere.
app := goshopify.App{
	ApiKey: "abcd",
	ApiSecret: "efgh",
	RedirectUrl: "https://example.com/shopify/callback",
	Scope: "read_products,read_orders",
}