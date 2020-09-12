package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"regexp"
	"strings"
)

var botRegExpHeader []*regexp.Regexp

func initRegExp() {
	var st = `ISUCONbot(-Mobile)?
ISUCONbot-Image\/
^ISUCONbot-Image.*
Mediapartners-ISUCON
ISUCONCoffee
ISUCONFeedSeeker(Beta)?
crawler \(https:\/\/isucon\.invalid\/(support\/faq\/|help\/jp\/)
isubot
Isupider
Isupider(-image)?\+
.*BOT.*
.*spider.*
.*Spider.*
.*SPIDER.*
.*Crawler.*
.*crawler.*
.*bot.*`
	// (bot|crawler|spider)(?:[-_ .\/;@()]|$)/i // これできてない。
	botRegExpStringArr := strings.Split(st, "\n")
	maxSize := len(botRegExpStringArr)
	botRegExpHeader	= make([]*regexp.Regexp, maxSize)

	for i, botRSt := range botRegExpStringArr {
		println("botRSt: ", botRSt)
		botRegExpHeader[i] = regexp.MustCompile(botRSt)
		pp := botRegExpHeader[i]
		b := pp.MatchString("ISUCONbot-Image/")
		log.Debugf("found  botRst : ", botRSt, ", is ISUCONbot-Image? : ", b)
	}

}

type errorImpl struct  {
	val string
}
func (e errorImpl) Error() string {
	return e.val
}

func HandleBot() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header

			userAgentValue := header.Get("User-Agent")

			log.Debugf("[HandleBot] header: %v\n", header)
			log.Debugf("[HandleBot] user agent: %v \n", userAgentValue)

			for _, compiled := range botRegExpHeader {
				if compiled.MatchString(userAgentValue) {
					c.Error(errorImpl{"error user is bot."})
				}
			}
			return next(c)
		}
	}
}
