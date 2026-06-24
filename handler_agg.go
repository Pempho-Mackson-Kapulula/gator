package main

import (
	"context"
	"fmt"
	"time"
	"strings"
	"database/sql"

	"github.com/Pempho-Mackson-Kapulula/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1{
		return fmt.Errorf("please enter the duration only")
	}

	durationStr := cmd.args[0]
	timeBetweenRequests , err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}


func scrapeFeeds(s *state){
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil{
		fmt.Printf("could not get feed: %v\n", err)
		return
	}

	fmt.Printf("found feed from: %v\n", nextFeed.Url)

	markedFeed , err := s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil{
		fmt.Printf("could not mark feed as fetched: %v\n", err)
		return
	}

	feed, err := fetchFeed(context.Background(), markedFeed.Url)
	if err != nil{
		fmt.Printf("could not fetch feed: %v\n", err)
		return
	}

	itemsCount := len(feed.Channel.Item)
	fmt.Printf(" -> Fetched %d items from feed\n", itemsCount)

	layout := time.RFC1123Z
	skippedCount := 0

	for _, item := range feed.Channel.Item{
		publishedAt := sql.NullTime{}
		description := sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}

		//parse published at (time)
		t, err := time.Parse(layout,item.PubDate)
		if err == nil {
			publishedAt = sql.NullTime{
				Time: t,
				Valid: true,
			}
		}else {
			publishedAt = sql.NullTime{
				Valid: false, 
			}
		}

		//convert each item into data
		params := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: description,
			FeedID: markedFeed.ID,
			PublishedAt: publishedAt,
		}

		err = s.db.CreatePost(context.Background(), params)
		if err !=  nil {
			// Count duplicate errors instead of printing them
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint"){
				skippedCount++
				continue
			}

			fmt.Printf("could not create post: %v\n", err)
			continue
		}
		
		// Print only for genuine fresh posts saved successfully
		pubDateStr := "unknown date"
		if params.PublishedAt.Valid {
			pubDateStr = params.PublishedAt.Time.Format("2006-01-02 15:04")
		}
		fmt.Printf(" -> Saved Post: %s\n    URL:  %s\n    Date: %s\n", params.Title, params.Url, pubDateStr)
	}	

	// Print the summary of skipped duplicates if any exist
	if skippedCount > 0 {
		fmt.Printf(" -> Skipped %d duplicate posts already in database\n", skippedCount)
	}
}
