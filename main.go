package urlquery

var defaultKey string
var defaultServer string

func init() {
	defaultServer = "api.urlquery.net"
}

func SetDefaultKey(apikey string) {
	defaultKey = apikey
}
