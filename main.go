package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/jifpbj/gator/internal/config"
	"github.com/jifpbj/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)

	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := &commands{registeredCommands: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	error := cmds.run(s, command{Name: cmdName, Args: cmdArgs})
	if error != nil {
		log.Fatal(err)
	}
}
