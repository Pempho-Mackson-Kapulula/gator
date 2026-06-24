package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Pempho-Mackson-Kapulula/gator/internal/config"
	"github.com/Pempho-Mackson-Kapulula/gator/internal/database"
	_ "github.com/lib/pq"
)

// state struct
type state struct {
	cfg     *config.Config
	db *database.Queries
}

func main() {
	//read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	//create db connection
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	//create bdQueries for stateInstance
	dbQueries := database.New(db)

	// create state and command instances
	programState := &state{
		cfg:     &cfg,
		db: dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	//register commands
	cmds.register("login",handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed",  middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFeedFollows))
	cmds.register("following", middlewareLoggedIn(handlerListFeedFollows))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		log.Fatal("invalid inputs")
	}

	cmdName := os.Args[1]
	args := os.Args[2:]

	command := command{
		name: cmdName,
		args: args,
	}

	err = cmds.run(programState, command)
	if err != nil {
		log.Fatal(err)
	}

}
