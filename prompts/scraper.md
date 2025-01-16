- create a go library function that can scrape data from the web and return it as a string.
- place it in the pkg/util/scraper.go file.
- use the colly library to perform the scraping.
- create one functions, ScrapeURL, that takes a url as a parameter and returns the scraped data as a string.
- create one function, ScrapeAll that takes a list of urls and returns the scraped data as a string. it should use the ScrapeUrl function to scrape each url in the list.
- make sure the file has the following documentation:

  - a package comment
  - a function comment for each function
  - a comment for each parameter
  - a comment for each return value
  - a comment for each error condition
  - an example usage of each function
  - relevant comments inside the functions

- add tests to pkg/util for scraper.go. include tests that download valid urls, invalid urls and non-existent urls. test both ScapeUrl and ScrapeAll
