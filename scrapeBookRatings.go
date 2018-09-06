package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	for _, blob := range getBlobs() {
		bleb := formatBlobQuery(blob)
		html := getHTML(bleb)
		ratings := getRating(html)

		fmt.Println(bleb)
		for _, rating := range ratings {
			fmt.Println(rating)
		}
		fmt.Println("")
	}
}

func formatBlobQuery(blob string) string {
	return strings.Split(blob, "(")[0]
}

func getHTML(query string) string {
	client := &http.Client{}
	cartURL := "https://www.goodreads.com/search?utf8=%E2%9C%93&search_type=books"

	req, _ := http.NewRequest("GET", cartURL, nil)
	q := req.URL.Query()
	q.Add("q", query)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return strings.TrimSpace(buf.String())
}

func getRating(rawHTML string) []string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHTML))

	if err != nil {
		panic(err)
	}

	items := doc.Find(".minirating")

	var ret []string
	items.Each(func(i int, selection *goquery.Selection) {
		ret = append(ret, selection.Text())
	})

	return ret
}

func getBlobs() []string {
	return []string{
		"Bailey, Beth, Sex in the Heartland (Cambridge, Mass.: Harvard University Press, 2002) [available as electronic resource to UG students].",

		"Gao, Mobo, The Battle for China’s Past: Mao and the Cultural Revolution (London: Pluto Press, 2008) [available as electronic resource to UG students].",

		"Grubbs, Larry, Secular Missionaries: Americans and African Development in the 1960s (Boston: University of Massachusetts Press, 2010). [Course reserve].",

		"Hartmann, Douglas, Race, Culture, and the Revolt of the Black Athlete: The 1968 Olympic Protests and Their Aftermath (Chicago: University of Chicago Press, 2003). [Course reserve].",

		"Joseph, Peniel, Waiting ‘Til the Midnight Hour: A Narrative History of Black Power in America (New York: Henry Holt, 2006). [Course reserve].",

		"Klatch, Rebecca, A Generation Divided: The New Left, the New Right, and the 1960s (Berkeley: University of California Press, 1999) [available as electronic resource to UG students].",

		"Lawrence, Mark Atwood, The Vietnam War: A Concise International History (Oxford: Oxford University Press, 2008). [Course reserve].",

		"MacDonald, Ian, Revolution in the Head: The Beatles’ Records and the Sixties (London: Fourth Estate, 1994) [available as electronic resource to UG students].",

		"Marqusee, Mike, Wicked Messenger: Bob Dylan and the 1960s (New York: Seven Stories, 2005). [Course reserve].",

		"Marqusee, Mike, Redemption Song: Muhammad Ali and the Spirit of the Sixties (New York: Verso, 1999). [Course reserve].",

		"Schildt, Axel and Detlef Siegfried (eds.), Between Marx and Coca Cola: Youth Cultures in Changing European Societies, 1960-1980 (New York & Oxford, Berghahn, 2006). [Course reserve].",

		"Seidman, Michael, The Imaginary Revolution: Parisian Students and Workers in 1968 (New York/Oxford: Berghahn, 2004). [Course reserve].",

		"Witherspoon, Kevin, Before the Eyes of the World: Mexico and the 1968 Olympics (De Kalb, IL: Northern Illinois Press, 2014). [Course reserve].",
	}
}
