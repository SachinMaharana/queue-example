package twitter

import (
	"SachinMaharana/twitbot/logger"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required env variable" + name)
	}
	return v
}

//GetHashTagStream ..
func GetHashTagStream(twitterStream chan<- string) {
	log := logger.NewLogger()
	var (
		consumerKey       = getenv("CONSUMER_KEY")
		consumerSecret    = getenv("CONSUMER_SECRET")
		accessToken       = getenv("ACCESS_TOKEN")
		accessTokenSecret = getenv("ACCESS_SECRET")
	)

	api := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)
	api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"#MasterThirdLook"},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Warningf("Recevied unexpected value of type %T", v)
			continue
		}
		twitterStream <- t.User.Name
	}
	close(twitterStream)
}
