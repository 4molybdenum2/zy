package sitemap

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"github.com/4molybdenum2/zy/pkg/link"
)

type Links []link.Link

func bfs(urlString string, depth int) ([]string,error) {
	visited := make(map[string]bool)
	q := make([]string, 0)
	q = append(q, urlString)

	for i:= 0; i < depth; i++ {
		if len(q) == 0  {
			break
		}
		// loop through all the children of current urlstring and push them in the queue
		for _, url := range q {
			if(visited[url]) {
				continue
			}
			visited[url] = true
			for _, link := range getLinks(url) {
				q = append(q, link)
			}
		}
	}

	output := make([]string, 0, len(visited))
	for url := range visited {
		output = append(output, url)
	}
	return output,nil
}

func getLinks(siteLink string) []string {
	// create the base url for a link
	resp, err := http.Get(siteLink)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	ls, err := link.Parse(base)
	if(err != nil) {
		log.Fatal(err)
	}

	var ret []string
	for _, l := range ls {
		switch {
			case strings.HasPrefix(l.Href, "/"):
				ret = append(ret, base + l.Href)
			case strings.HasPrefix(l.Href, "http"):
				ret = append(ret, l.Href)
		}
	}
	return filter(ret, base)
}

func filter(links []string, base string) []string {
	p := make([]string, 0)
	for _, link := range links {
		if strings.HasPrefix(link, base) {
			p = append(p, link)
		}
	}
	return p
}

func BuildSitemap(siteLink string, depth int) []string {
	pages,err := bfs(siteLink, depth)
	if err != nil {
		log.Fatal(err)
	}
	return pages
}