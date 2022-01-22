package parser

// Create gorutine and start parses
func CreateTasks() {
	page := getCountPage("/zhenshchinam/odezhda/bryuki-i-shorty")
	countGor := checkCountPage(page)
	countObjInPage := page / countGor

	for i := 1; i < countGor+1; i++ {
		page = page - countObjInPage
		start := page
		end := page + countObjInPage
		// fmt.Println(start, end)
		go saveProduct(start, end)
		// fmt.Scanln()
	}
}
