# AgroJob

Thanks to this package, 
you will be able to conveniently work with site 
[opendata.trudvsem.ru](http://opendata.trudvsem.ru/api/v1/vacancies).

# Usage

Get response from REST API
```go
package main

import (
	"fmt"
	trudvsem "github.com/reiterus/goagrojob"
)

func main() {
	url := "http://opendata.trudvsem.ru/api/v1/vacancies/region/8300000000000"
	response := trudvsem.GetResponse(url)
	// If the string contains this specialization
	// Сельское хозяйство, экология, ветеринария
	if response != "" {
		// do something with string like this
		// "{"status":"200","results":{"vacancies":[{...}]}"
		fmt.Println(response)
	}
}
```

Get region codes `fmt.Println(trudvsem.GetRegionCodes())`
```
[0100000000000 0200000000000 ... etc]
```

Get region URLs `fmt.Println(trudvsem.GetRegionUrls())`
```
[http://opendata.trudvsem.ru/api/v1/vacancies/region/0100000000000 ... etc]

```
