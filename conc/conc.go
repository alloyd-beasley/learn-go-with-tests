package conc

type WebsiteChecker func(string) bool
type result struct {
	string 
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		//pass current value of url to go routine
		go func(url string) {

			//send a result to the channel
			resultChannel <- result{url, wc(url)}
		}(url)
	}

	for i:=0; i < len(urls); i++ {

		//receive a result to the result variable
		result := <- resultChannel

		//assign result from the channel to results
		results[result.string] = result.bool
	}

	return results
}
