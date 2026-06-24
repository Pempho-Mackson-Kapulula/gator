package main

import (
	"context"
	"fmt"
	"strconv"
	"regexp"

	
	"github.com/Pempho-Mackson-Kapulula/gator/internal/database"
)

func stripHTML(s string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(s, "")
}

func handlerBrowse(s *state, cmd command, user database.User)error{
	limit := 2
	if len(cmd.args) == 1 {
		num, err := strconv.Atoi(cmd.args[0])
		if err != nil{
			return fmt.Errorf("could not parse limit: %v", err)
		}

		limit = num
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil{
		return fmt.Errorf("could not get posts from database: %v", err)
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println("Feed:", post.FeedName)

		if post.PublishedAt.Valid {
			fmt.Println("Published:", post.PublishedAt.Time.Format("2006-01-02 15:04"))
		}

		fmt.Println("URL:", post.Url)
		fmt.Println("-----------------------------------")
	}


	return nil
}