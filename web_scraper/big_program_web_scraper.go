// Author: Carlos Moncada Soto, Abby Baron, Ken Lam
// Program Name: big_program_web_scraper.go
// Date: 04/07/2025
// Description: A concurrent web scraper that takes a list of websites, crawls each one,
// finds all <a href="..."> links recursively, and stores the results in a text file called links.txt
// Each link is printed to the terminal and saved to the file.
// Usage: First install: go get golang.org/x/net/html, then go run big_program_web_scraper.go

// References:
// - https://pkg.go.dev/net/http for handling HTTP GET requests
// - https://www.golang-book.com/books/intro/10 for concurrency and goroutines
// - https://github.com/neo-liang-sap/book/blob/master/Go/The.Go.Programming.Language.pdf for Go
// - https://medium.com/@datajournal/parse-html-in-golang-83c882576a0a for understanding the recursion behind parsing

// Packages group related functions and consist of all files within the same directory
package main

// packages used:
import (
	"bufio"    // For buffered writing to files
	"fmt"      // For printing to terminal and formatting strings
	"io"       // For I/O
	"net/http" // For making GET requests to URLs
	"os"       // For creating file
	"sync"     // For concurrency wait groups

	"golang.org/x/net/html" // For parsing HTML (must be installed separately with: go get golang.org/x/net/html)
)

// parseAnchorTags takes an HTML document (as an io.Reader) and extracts all <a href="..."> links
// It returns a slice of strings containing the link values
func parseAnchorTags(r io.Reader) []string {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println("HTML parsing failed:", err)
		return nil
	}

	var links []string

	// recursive function that traverses the HTML tree and collects all hrefs
	// needs recursion because it is a tree structure, and works best
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		// Recurse through children with this call, base case is here
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	traverse(doc)
	return links
}

// scrape makes an HTTP request to a single URL, parses its links, and sends each one into a channel.
// It decrements the wait group counter when done. This function runs as a goroutine.
func scrape(url string, channel chan<- string, wait_group *sync.WaitGroup) {
	defer wait_group.Done() // Mark this goroutine as done when function exits

	res, err := http.Get(url) // GET request
	if err != nil {
		channel <- fmt.Sprintf("Error: %s for %s", url, err)
		return
	}
	defer res.Body.Close() // Close the response body to free resources

	anchors := parseAnchorTags(res.Body) // Extract all <a href="..."> (which are links)
	for _, link := range anchors {
		channel <- fmt.Sprintf("%s %s", url, link) // Send each link to the channel
	}
}

// Creates a file to store the scraped links and defines the list of sites to scrape
// It returns a file handle, a writer object, and the site URLs
func setup() (*os.File, *bufio.Writer, []string) {
	file, err := os.Create("links.txt") // Creates the file
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return nil, nil, nil
	}

	writer := bufio.NewWriter(file) // Buffered writer for efficient file writing into links.txt

	// List of websites to scrape
	sites := []string{
		"https://golang.org",
		"https://www.hamilton.edu",
		"https://www.hamilton.edu/academics/departments/computer-science",
	}

	return file, writer, sites
}

// Launches a goroutine for each site to scrape links concurrently
// It returns a channel where all links will be sent
func webscrapingThings(sites []string) chan string {
	linkStream := make(chan string) // channel for links, provides a way to synchronize execution
	var wait_group sync.WaitGroup   // synchronizes the routines so we end all at the same time

	// Launch a scraping goroutine for each site
	for _, site := range sites {
		wait_group.Add(1)
		go scrape(site, linkStream, &wait_group)
	}

	// waits for all routines and then closes the channel once all goroutines are done
	go func() {
		wait_group.Wait()
		close(linkStream)
	}()

	return linkStream
}

// main is the entry point of the program, and creates the whole program
func main() {
	file, writer, sites := setup()
	if file == nil || writer == nil || sites == nil {
		fmt.Println("Setup failed. Exiting.")
		return
	}
	defer file.Close()

	linkStream := webscrapingThings(sites)

	// Read from the channel, print each link, and write it to the file
	for link := range linkStream {
		fmt.Println(link)
		writer.WriteString(link + "\n")
	}

	// Ensure all data is written to the file
	writer.Flush()

	fmt.Println("\nAll links have been saved to links.txt :)")
}
