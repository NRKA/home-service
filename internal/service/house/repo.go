package house

import (
	"context"
	"github.com/NRKA/home-service/internal/usecase"
	"github.com/NRKA/home-service/pkg/postgres"
)

const (
	moderator = "moderator"
	approved  = "approved"
)

type Repo struct {
	db *postgres.Database
}

func NewRepo(db *postgres.Database) *Repo {
	return &Repo{
		db: db,
	}
}

func (repo *Repo) Create(ctx context.Context, house usecase.HouseCreateRequest) (usecase.House, error) {
	var response usecase.House

	query := `
		INSERT INTO house (address, year, developer) 
		VALUES ($1, $2, $3) 
		RETURNING *
	`

	err := repo.db.Get(ctx, &response, query, house.Address, house.Year, house.Developer)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (repo *Repo) ClientFlats(ctx context.Context, houseID int) ([]usecase.FlatResponse, error) {
	query := `
		SELECT id, number, house_id, price, rooms, status
		FROM flat
		WHERE house_id = $1 AND status = $2
	`

	var flats []usecase.FlatResponse
	err := repo.db.Select(ctx, &flats, query, houseID, approved)
	if err != nil {
		return flats, err
	}

	return flats, nil
}

func (repo *Repo) ModeratorFlats(ctx context.Context, houseID int) ([]usecase.FlatResponse, error) {
	query := `
		SELECT id, number, house_id, price, rooms, status
		FROM flat
		WHERE house_id = $1
	`

	var flats []usecase.FlatResponse
	err := repo.db.Select(ctx, &flats, query, houseID)
	if err != nil {
		return flats, err
	}

	return flats, nil
}
