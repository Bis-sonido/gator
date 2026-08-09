package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Bis-sonido/gator/internal/config"
	"github.com/Bis-sonido/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
		return
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return
	}
	defer db.Close()
	dbQueries := database.New(db)

	appState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", registerUser)
	cmds.register("reset", resetUser)
	cmds.register("users", getUsers)
	cmds.register("agg", aggCommand)
	cmds.register("addfeed", middlewareLoggedIn(addFeed))
	cmds.register("feeds", feeds)
	cmds.register("follow", middlewareLoggedIn(follow))
	cmds.register("following", middlewareLoggedIn(following))
	cmds.register("unfollow", middlewareLoggedIn(unfollow))
	cmds.register("browse", middlewareLoggedIn(browse))

	if len(os.Args) < 2 {
		log.Fatal("No command provided. Usage: gator <command> [args]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = cmds.run(appState, cmd)
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}
}
