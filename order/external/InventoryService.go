package external 

import (
	"github.com/go-resty/resty/v2"
	"fmt"
	"strconv"
	"order/config"
)

var client = resty.New()

func Order(entity *Inventory) *resty.Response{
	config.ConfigMode()
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s", options["api_url_inventory"], "inventories" )
	resp, _ := client.R().SetBody(entity).Post(target)

	return resp
}

