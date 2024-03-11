package twitch

import (
	"errors"
	"log"
	"net/http"
	"time"
)

func ValidateAccessToken(accessToken string) error {
	c := http.Client{Timeout: 1 * time.Second}

	req, err := http.NewRequest("GET", "https://id.twitch.tv/oauth2/validate", nil)
	if err != nil {
		log.Fatalf("Unable to fetch oauth validate endpoint %s\n", err)
		return err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	res, err := c.Do(req)
	if err != nil {
		log.Fatalf("Unable to make request to validate endpoint %s\n", err)
		return err
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Accesstoken is not valid")
		return errors.New("accesstoken invalid")
	}

	return nil
}
