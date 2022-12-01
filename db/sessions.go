package db

import (
	"errors"
	"net/http"
	"pood/models"
	"time"
)

func InsertSession(token string, userId int) error {
	_, err := DB.Exec("delete from sessions where user_id = ?", userId)
	if err != nil {
		return err
	}

	_, err = DB.Exec("insert into sessions (id, user_id, created_date) values (?,?, strftime('%s','now'))", token, userId)
	if err != nil {
		return err
	}

	return nil
}

// look closer to time.Time parsing etc...
func CheckSession(r *http.Request) (*models.Session, error) {
	token, err := r.Cookie("session")
	if err != nil {
		// Cookie returns the named cookie provided in the request or ErrNoCookie if not found.
		return nil, err
	}
	session := &models.Session{}
	row := DB.QueryRow("select id, user_id, created_date from sessions where id = ?", token.Value)
	session.User = &models.User{}
	var createDate int64 // unix time stamp
	err = row.Scan(&session.Id, &session.User.Id, &createDate)
	if err != nil {
		return nil, err
	}
	session.User, err = GetFirstNameById(session.User.Id)
	if err != nil {
		return nil, err
	}
	t := time.Unix(int64(createDate), 0) // time.Time

	if session.Id == "" {
		return nil, errors.New("token invalid or expired")
	}

	if t.AddDate(0, 0, 1).Before(time.Now()) {
		err := DeleteSession(session.Id)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("token invalid or expired")
	}

	return session, nil
}

func DeleteSession(token string) error {
	_, err := DB.Exec("delete from sessions where id = ?", token)
	if err != nil {
		return nil
	}

	return nil
}
