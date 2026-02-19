package repository

func (r *RepositoryModule) SaveReminder(userId int, reminder string) error {
	query := `INSERT INTO reminders (user_id,reminder)
VALUES ($1,$2);`
	_, err := r.db.Exec(query, userId, reminder)
	if err != nil {
		return err
	}
	return nil
}
