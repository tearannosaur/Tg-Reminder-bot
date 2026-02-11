package repository

import (
	"reminder/models"
)

func (r *RepositoryModule) SaveUsers(user models.Users) error {
	query := `INSERT INTO users (id,nick_name,first_name,last_name) VALUES ($1,$2,$3,$4)`
	_, err := r.db.Exec(query, user.Id, user.NickName, user.FirstName, user.LastName)
	if err != nil {
		return err
	}
	return nil
}
