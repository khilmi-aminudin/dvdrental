package repository

import (
	"context"
	"database/sql"
	"dvdrental/helper"
	"dvdrental/models/entity"
	"errors"
	"time"
)

type ActorRepository interface {
	Create(ctx context.Context, tx *sql.Tx, actor entity.Actor) entity.Actor
	Update(ctx context.Context, tx *sql.Tx, actor entity.Actor) entity.Actor
	Delete(ctx context.Context, tx *sql.Tx, actor entity.Actor) error
	Find(ctx context.Context, tx *sql.Tx, actor entity.Actor) entity.Actor
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Actor
}

var sqlQuery string

type actorRepository struct{}

func NewActorRepository() ActorRepository {
	return &actorRepository{}
}

func (repository *actorRepository) Create(ctx context.Context, tx *sql.Tx, actor entity.Actor) entity.Actor {
	sqlQuery = "INSERT INTO actor(first_name,last_name) VALUES($1,$2) RETURNING *;"

	result, err := tx.ExecContext(ctx, sqlQuery, actor.FirstName, actor.LastName)
	helper.LogErrorWithFields(err, "actorParams", actor)

	id, err := result.LastInsertId()
	helper.LogError(err)

	return entity.Actor{
		ActorID:    id,
		FirstName:  actor.FirstName,
		LastName:   actor.LastName,
		LastUpdate: time.Now(),
	}
}

func (repository *actorRepository) Update(ctx context.Context, tx *sql.Tx, actor entity.Actor) entity.Actor {
	sqlQuery = "UPDATE actor SET first_name = $2, last_name = $3 WHERE actor_id = $1 RETURNING *;"

	result, err := tx.ExecContext(ctx, sqlQuery, actor.ActorID, actor.FirstName, actor.LastName)
	helper.LogErrorWithFields(err, "actorParams", actor)

	affects, err := result.RowsAffected()
	helper.LogError(err)

	if affects != 1 {
		return entity.Actor{}
	}

	return entity.Actor{
		ActorID:    actor.ActorID,
		FirstName:  actor.FirstName,
		LastName:   actor.LastName,
		LastUpdate: time.Now(),
	}

}

func (repository *actorRepository) Delete(ctx context.Context, tx *sql.Tx, actor entity.Actor) error {
	sqlQuery = "DELETE FROM actor WHERE actor_id = $1;"

	result, err := tx.ExecContext(ctx, sqlQuery, actor.ActorID)
	if err != nil {
		helper.LogError(err)
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		helper.LogError(err)
		return err
	}

	if affect != 1 {
		newError := errors.New("no actor data deleted")
		helper.LogError(newError)
		return newError
	}

	return nil
}

func (repository *actorRepository) Find(ctx context.Context, tx *sql.Tx, actor entity.Actor) entity.Actor {
	sqlQuery = "SELECT actor_id, first_name, last_name, last_update FROM actor WHERE actor_id = $1;"

	row := tx.QueryRowContext(ctx, sqlQuery, actor.ActorID)

	err := row.Err()
	helper.LogError(err)

	var result entity.Actor

	err = row.Scan(
		&result.ActorID,
		&result.FirstName,
		&result.LastName,
		&result.LastUpdate,
	)

	helper.LogError(err)

	return result
}

func (repository *actorRepository) FindAll(ctx context.Context, tx *sql.Tx) []entity.Actor {
	sqlQuery = "SELECT actor_id, first_name, last_name, last_update FROM actor;"

	rows, err := tx.QueryContext(ctx, sqlQuery)
	helper.LogError(err)

	var result []entity.Actor

	for rows.Next() {
		var actor entity.Actor
		err = rows.Scan(
			actor.ActorID,
			actor.FirstName,
			actor.LastName,
			actor.LastUpdate,
		)
		helper.LogError(err)
		result = append(result, actor)
	}

	return result
}
