package main

import (
	"fmt"
	"github.com/labstack/echo"
	// "github.com/labstack/gommon/log"
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
Isupider(-image)?\+`
	///// (bot|crawler|spider)(?:[-_ .\/;@()]|$)/i // これできてない。
	botRegExpStringArr := strings.Split(st, "\n")
	maxSize := len(botRegExpStringArr)
	botRegExpHeader	= make([]*regexp.Regexp, maxSize)
	for i, botRSt := range botRegExpStringArr {
		botRegExpHeader[i] = regexp.MustCompile(botRSt)
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

			// log.Debugf("[HandleBot] header: %v\n", header)
			// log.Debugf("[HandleBot] user agent: %v \n", userAgentValue)

			for i, compiled := range botRegExpHeader {
				if compiled.MatchString(userAgentValue) {
					// println("useragent: ", userAgentValue, ", compiled num: ", i)
					errorString := fmt.Sprintf("error user is bot. the strings %s is matched with rule %d",
						userAgentValue,
						i)
					c.Error(errorImpl{errorString})
					break
				}
			}
			return next(c)
		}
	}
}
