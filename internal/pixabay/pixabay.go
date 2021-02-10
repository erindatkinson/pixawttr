package pixabay

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/hashicorp/errwrap"
)

const endpoint = "https://pixabay.com/api/"

type pixabayResponse struct {
	Total     int          `json:"total"`
	TotalHits int          `json:"totalHits"`
	Hits      []pixabayHit `json:"hits"`
}

type pixabayHit struct {
	ID              int    `json:"ID"`
	PageURL         string `json:"PageURL"`
	Type            string `json:"Type"`
	Tags            string `json:"Tags"`
	PreviewURL      string `json:"PreviewURL"`
	PreviewWidth    int    `json:"PreviewWidth"`
	PreviewHeight   int    `json:"PreviewHeight"`
	WebformatURL    string `json:"WebformatURL"`
	WebformatWidth  int    `json:"WebformatWidth"`
	WebformatHeight int    `json:"WebformatHeight"`
	LargeImageURL   string `json:"LargeImageURL"`
	FullHDURL       string `json:"FullHDURL"`
	ImageURL        string `json:"ImageURL"`
	ImageWidth      int    `json:"ImageWidth"`
	ImageHeight     int    `json:"ImageHeight"`
	Views           int    `json:"Views"`
	Downloads       int    `json:"Downloads"`
	Favorites       int    `json:"Favorites"`
	Likes           int    `json:"Likes"`
	Comments        int    `json:"Comments"`
	UserID          int    `json:"UserID"`
	User            string `json:"User"`
	UserImageURL    string `json:"UserImageURL"`
}

//GetImage gets an image
func GetImage(apiKey string, query string) (string, error) {

	hits, err := getList(apiKey, query)
	if err != nil {
		return "", errwrap.Wrapf("error listing background images: {{err}}", err)
	}

	img, err := downloadImage(hits)
	if err != nil {
		return "", errwrap.Wrapf("error downloading background image: {{err}}", err)
	}

	return img, nil

}

func getList(apiKey string, query string) (*pixabayResponse, error) {
	apiURL, _ := url.Parse(endpoint)
	q := apiURL.Query()
	q.Set("key", apiKey)
	q.Set("q", query)
	apiURL.RawQuery = q.Encode()

	resp, err := http.Get(apiURL.String())
	if err != nil {
		return &pixabayResponse{}, err
	}

	var hits pixabayResponse

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&hits)
	if err != nil {
		return &pixabayResponse{}, err
	}

	return &hits, nil
}

func downloadImage(hits *pixabayResponse) (string, error) {
	rand.Seed(time.Now().Unix())
	choice := rand.Intn(len(hits.Hits))

	resp, err := http.Get(hits.Hits[choice].LargeImageURL)
	if err != nil {
		return "", errwrap.Wrapf("error requesting image data: {{err}}", err)
	}

	filetype := filepath.Ext(hits.Hits[choice].LargeImageURL)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errwrap.Wrapf("error reading image data: {{err}}", err)
	}

	err = ioutil.WriteFile("bgImage"+filetype, data, 0644)
	if err != nil {
		return "", errwrap.Wrapf("error writing image data: {{err}}", err)
	}

	return "bgImage" + filetype, nil

}
