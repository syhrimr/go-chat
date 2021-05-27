package acc

import (
	"time"

	"github.com/lolmourne/go-accounts/model"
)

func (dbr *DBResource) Register(username string, password string, salt string) error {
	query := `
		INSERT INTO
			account
		(
			username,
			password,
			salt,
			created_at,
			profile_pic
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5
		)
	`

	_, err := dbr.db.Exec(query, username, password, salt, time.Now(), "")
	if err != nil {
		return err
	}

	return nil
}

func (dbr *DBResource) GetUserByUserID(userID int64) (model.User, error) {
	query := `
	SELECT 
		user_id,
		username,
		password,
		salt,
		created_at,
		profile_pic
	FROM
		account
	WHERE
		user_id = $1
	`

	var user UserDB
	err := dbr.db.Get(&user, query, userID)
	if err != nil {
		return model.User{}, nil
	}

	return model.User{
		UserID:     user.UserID.Int64,
		Username:   user.UserName.String,
		Password:   user.Password.String,
		Salt:       user.Salt.String,
		CreatedAt:  user.CreatedAt,
		ProfilePic: user.ProfilePic.String,
	}, nil
}

func (dbr *DBResource) GetUserByUserName(userName string) (model.User, error) {
	query := `
	SELECT 
		user_id,
		username,
		password,
		salt,
		created_at,
		profile_pic
	FROM
		account
	WHERE
		username = $1
	`

	var user UserDB
	err := dbr.db.Get(&user, query, userName)
	if err != nil {
		return model.User{}, nil
	}

	return model.User{
		UserID:     user.UserID.Int64,
		Username:   user.UserName.String,
		Password:   user.Password.String,
		Salt:       user.Salt.String,
		CreatedAt:  user.CreatedAt,
		ProfilePic: user.ProfilePic.String,
	}, nil
}

func (dbr *DBResource) UpdateUserProfpic(userID int64, newProfpic string) error {
	query := `
		UPDATE
			account
		SET 
		    profile_pic = $1
		WHERE
			user_id = $2
	`

	_, err := dbr.db.Exec(query, newProfpic, userID)
	if err != nil {
		return err
	}

	return nil
}

func (dbr *DBResource) UpdateUserPassword(userID int64, newPassword string) error {
	query := `
		UPDATE
			account
		SET 
		    password = $1
		WHERE
			user_id = $2
	`

	_, err := dbr.db.Exec(query, newPassword, userID)
	if err != nil {
		return err
	}

	return nil
}

func (dbr *DBResource) UpdateUserName(userID int64, newUsername string) error {
	query := `
		UPDATE
			account
		SET 
		    username = $1
		WHERE
			user_id = $2
	`

	_, err := dbr.db.Exec(query, newUsername, userID)
	if err != nil {
		return err
	}

	return nil
}
