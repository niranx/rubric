package main

func main() {
	url := parseArgs()
	headers := fetchHeaders(url)
	printHeaders(headers)
}
