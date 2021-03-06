package dbrepo

import (
	"context"
	"time"

	"github.com/lazyspell/Ecommerce_Backend/models"
)

func (m *postgresDBRepo) Authenticate(email string) (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.Users

	query := `select id, first_name, last_name, email, password, "authorization" from users where email = $1`

	person := m.DB.QueryRowContext(ctx, query, email)

	err := person.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Authorization,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *postgresDBRepo) GoogleAuthenticate(email string) (models.GoogleObject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.GoogleObject

	query := `select id, email, name, given_name from google_user where email = $1`

	person := m.DB.QueryRowContext(ctx, query, email)

	err := person.Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.GivenName,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
