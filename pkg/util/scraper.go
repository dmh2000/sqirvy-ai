// Package util provides utility functions for web scraping and data processing.
//
// This file implements web scraping functionality using the colly library,
// allowing for both single URL and batch URL scraping operations.
package util

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
)

const MaxScraperDepth = 2

// ScrapeURL scrapes the content from a single URL and returns it as a string.
//
// Parameters:
//   - url: The URL to scrape (must be a valid HTTP/HTTPS URL)
//
// Returns:
//   - string: The scraped content with preserved structure
//   - error: Error if scraping fails, URL is invalid, or site is unreachable
//
// Example usage:
//
//	content, err := ScrapeURL("https://example.com")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(content)
func ScrapeURL(link string) (string, error) {
	// Validate URL is not empty
	if link == "" {
		return "", fmt.Errorf("URL cannot be empty")
	}

	// validate the url
	_, err := url.ParseRequestURI(link)
	if err != nil {
		return "", fmt.Errorf("failed to scrape URL %s: %w", link, err)
	}

	// Initialize collector
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.MaxDepth(MaxScraperDepth),
	)

	// Store scraped content
	var content strings.Builder

	// Collect text content
	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Get text content while preserving some structure
		content.WriteString(e.Text)
	})

	// Handle errors
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error scraping %s: %v\n", r.Request.URL, err)
	})

	// Start scraping
	err = c.Visit(link)
	if err != nil {
		return "", fmt.Errorf("failed to scrape URL %s: %w", link, err)
	}

	text := fmt.Sprintf("```%s\n%s```\n", link, content.String())
	return text, nil
}

// ScrapeAll scrapes content from multiple URLs and concatenates the results.
//
// Parameters:
//   - urls: A slice of URLs to scrape (must be valid HTTP/HTTPS URLs)
//
// Returns:
//   - string: Concatenated content from all successfully scraped URLs
//   - error: Error if all URLs fail to scrape or if urls slice is empty
//
// Example usage:
//
//	urls := []string{
//	    "https://example.com",
//	    "https://example.org",
//	}
//	content, err := ScrapeAll(urls)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(content)
func ScrapeAll(urls []string) (string, error) {
	// Validate URLs slice is not empty
	if len(urls) == 0 {
		return "", fmt.Errorf("URLs list cannot be empty")
	}

	// Store combined content
	var allContent strings.Builder
	successCount := 0

	// Scrape each URL
	for _, url := range urls {
		content, err := ScrapeURL(url)
		if err != nil {
			return "", fmt.Errorf("failed to scrape URL %s: %w", url, err)
		}

		// Add separator between URLs
		if successCount > 0 {
			allContent.WriteString("\n---\n")
		}

		allContent.WriteString(content)
		successCount++
	}

	// Check if any URLs were successfully scraped
	if successCount == 0 {
		return "", fmt.Errorf("failed to scrape any URLs")
	}

	return allContent.String(), nil
}
