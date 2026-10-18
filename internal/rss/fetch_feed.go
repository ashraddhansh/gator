package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)

	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	
	req.Header.Set("User-Agent", "gator")

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Unexpected HTTP status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	

	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}

	feed.unescapeHTML()
	return &feed, nil


}

func (feed *RSSFeed) unescapeHTML() {

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

}
