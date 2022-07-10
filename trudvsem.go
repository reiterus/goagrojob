package goagrojob

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Specialisation what we need
const spec = "Сельское хозяйство, экология, ветеринария"

// URL to the domain
const site = "http://opendata.trudvsem.ru"

// Zero suffix
const zeros = "00000000000"

// GetResponse Get response from Opendata REST API
func GetResponse(url string) string {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	result := string(data)

	if !strings.Contains(result, spec) {
		result = ""
	}

	return result
}

// GetRegionUrls Get all opendata region URLs
func GetRegionUrls() []string {
	var result []string
	format := "%s/api/v1/vacancies/region/%s"
	items := GetRegionCodes()

	for _, item := range items {
		url := fmt.Sprintf(format, site, item)
		result = append(result, url)
	}

	return result
}

// GetRegionCodes Get all RF region codes
func GetRegionCodes() []string {
	var items []string
	var ids []string
	var code string

	for i := 1; i < 80; i++ {
		idx := strconv.Itoa(i)
		if i < 10 {
			code = "0" + idx
		} else {
			code = idx
		}

		rgn := code + zeros
		items = append(items, rgn)
	}

	nums := []int{83, 86, 87, 89, 91, 92, 99}

	for _, num := range nums {
		ix := strconv.Itoa(num)
		rn := ix + zeros
		ids = append(ids, rn)
	}

	return append(items, ids...)
}
