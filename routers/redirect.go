package routers

import (
	"log"
	"net/http"

	"github.com/berabulut/kapsule/helpers"
	"github.com/berabulut/kapsule/models"
	db "github.com/berabulut/kapsule/mongo"
	"github.com/gin-gonic/gin"
)

func RedirectURL(records map[string]*models.ShortURL) func(c *gin.Context) {

	return func(c *gin.Context) {
		key := c.Request.URL.Path[1:]

		record, found := records[key]

		if found {
			userAgent := helpers.ParseUserAgent(c.Request.UserAgent())
			language := helpers.ParseLanguage(c.Request.Header.Get("Accept-Language"))
			remoteAddr := c.Request.RemoteAddr
			xForwardedFor := c.Request.Header.Get("X-FORWARDED-FOR")
			countryCode, _ := helpers.GetCountryCode(xForwardedFor)

			helpers.HandleClick(record, userAgent, language, remoteAddr, xForwardedFor, countryCode)

			err := db.HandleClick(key, record.Clicks, record.LastTimeVisited, record.Visits)
			if err != nil {
				log.Fatal(err)
			}

			c.Redirect(http.StatusFound, records[key].Value)
			return
		}

		r, err := db.GetRecord(key)

		if err != nil {
			c.Redirect(http.StatusSeeOther, "http://localhost:3000/")
			return
		}

		if r.Key != "" {
			c.Redirect(http.StatusFound, r.Value)
			return
		}

		c.Redirect(http.StatusSeeOther, "http://localhost:3000/")
	}
}
