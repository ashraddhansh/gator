package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ashraddhansh/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Please provide the username")

	}

	userName := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), userName)
	
	if err == nil {
		return fmt.Errorf("user %q already exists", userName)
		os.Exit(1)
	}
	if !errors.Is(err, sql.ErrNoRows){
		return err
	}
	
	user, err := s.db.CreateUser(context.Background(),database.CreateUserParams{
		ID:	uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: userName,
	})

	if err != nil {
		return err
	}

	if err := s.config.SetUser(userName); err != nil {
		return err
	}
	fmt.Printf("User %q created successfully!\n", userName)
	fmt.Printf("User data: %+v\n", user)
	return nil
}
