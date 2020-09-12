package main

import (
	"fmt"
	"github.com/labstack/echo"
	"regexp"
	"strings"
)

var botRegExpHeader = make(map[string]*regexp.Regexp)

func initRegExp() {
	//var botRegExpString =
	var st = `/ISUCONbot(-Mobile)?/
/ISUCONbot-Image\//
/Mediapartners-ISUCON/
/ISUCONCoffee/
/ISUCONFeedSeeker(Beta)?/
/crawler \(https:\/\/isucon\.invalid\/(support\/faq\/|help\/jp\/)/
/isubot/
/Isupider/
/Isupider(-image)?\+/
/(bot|crawler|spider)(?:[-_ .\/;@()]|$)/i`
	botRegExpStringArr := strings.Split(st, "\n")

	//botRegExpStringArr := [...]string{
	//
	//} // こうかける
	for _, botRSt := range botRegExpStringArr {
		println("botRSt: ", botRSt)
	}

	//  *regexp.Regexp
	//var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

}

func HandleBot() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header
			fmt.Printf("[HandleBot] header: %v\n", header)
			userAgentValue := header.Get("User-Agent")
			fmt.Printf("[HandleBot] user agent: %v \n", userAgentValue)
			return next(c)
		}
	}
}
