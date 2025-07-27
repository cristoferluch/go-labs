package games

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Game struct {
	Name          string
	URL           string
	OriginalPrice string
	FinalPrice    string
}

func GetGames() ([]Game, error) {

	res, err := http.Get("https://store.steampowered.com/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	games := make([]Game, 0)

	divDiscounts := doc.Find(`[aria-labelledby="tab_discounts"]`)
	divDiscounts.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {

		if i == 10 {
			return false
		}

		name := s.Find(".tab_item_name").Text()
		url, _ := s.Attr("href")
		originalPrice := s.Find(".discount_original_price").Text()
		finalPrice := s.Find(".discount_final_price").Text()

		var game Game

		game.Name = name
		game.URL = url
		game.OriginalPrice = originalPrice
		game.FinalPrice = finalPrice

		games = append(games, game)

		return true
	})

	return games, nil
}
