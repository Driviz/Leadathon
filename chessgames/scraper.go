package chessgames

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Data struct {
	Name  string
	Moves map[int][]string
}

type DataMap map[string]Data

func GetFile(ctx context.Context, url string) (DataMap, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln("error fetching from url", url, "error", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalln("error unexpected response code", url, "responsecode", res.StatusCode, "error", errors.New("response non 200"))
		return nil, err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	bodyString := string(bodyBytes)

	data, err := MapTableToStruct(ctx, bodyString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}

func MapTableToStruct(ctx context.Context, data string) (DataMap, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}

	result := make(DataMap, 0)
	doc.Find("tr").Each(func(idxtr int, rowhtml *goquery.Selection) {
		var data Data
		var code string
		rowhtml.Find("td").Each(func(idxrd int, rowdata *goquery.Selection) {
			if idxrd == 0 {
				code = rowdata.Find("font").Text()
			}
			if idxrd == 1 {
				rowdata.Find("font").Each(func(idxd int, ddata *goquery.Selection) {
					if idxd == 0 {
						name := ddata.Find("b").Text()
						data.Name = name
					}
					if idxd == 1 {
						moves := strings.Split(ddata.Text(), " ")
						m := make(map[int][]string, 0)
						var idx int
						for _, v := range moves {
							if i, err := strconv.Atoi(v); err == nil {
								idx = i
								continue
							}
							cm := m[idx]
							cm = append(cm, v)
							m[idx] = cm
						}
						data.Moves = m
					}
				})
			}
		})
		result[code] = data
	})

	return result, nil
}
