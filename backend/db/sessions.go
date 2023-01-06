package db

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"pood/models"
	"time"
)

func InsertSession(token string, userId int) error {
	_, err := DB.Exec("delete from sessions where user_id = ?", userId)
	if err != nil {
		return err
	}
	_, err = DB.Exec("insert into sessions (id, user_id, date_created) values (?,?, strftime('%s','now'))", token, userId)
	if err != nil {
		return err
	}
	return nil
}

// look closer to time.Time parsing etc...
func CheckSession(r *http.Request) (*models.Session, error) {
	token, err := r.Cookie("session")
	if err != nil {
		log.Println(err.Error())
		// response with 401 ? unauthorized???
		// Cookie returns the named cookie provided in the request or ErrNoCookie if not found.
		return nil, err
	}
	session := &models.Session{}
	session.User = &models.User{}
	var createDate int64 // unix time stamp
	row := DB.QueryRow("select id, user_id, date_created from sessions where id = ?", token.Value)
	err = row.Scan(&session.Id, &session.User.Id, &createDate)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	// Why I need this FirstName ??
	session.User, err = GetFirstNameById(session.User.Id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if session.Id == "" {
		log.Println(err.Error())
		return nil, errors.New("token invalid or expired")
	}

	t := time.Unix(int64(createDate), 0) // time.Time
	if t.AddDate(0, 0, 1).Before(time.Now()) {
		err := DeleteSession(session.Id)
		if err != nil {
			log.Println(err.Error())

			return nil, err
		}
		log.Println(err.Error())
		return nil, errors.New("token invalid or expired")
	}
	return session, nil
}

func CheckAdminSession(r *http.Request) (*models.Session, error) {
	token, err := r.Cookie("session")
	if err != nil {
		return nil, models.ErrNoRecord
	}
	session := &models.Session{}
	session.User = &models.User{}
	var createdDate int64
	row := DB.QueryRow("SELECT id, user_id, date_created FROM sessions WHERE id = ?", token.Value)
	err = row.Scan(&session.Id, &session.User.Id, &createdDate)
	if err != nil {
		return nil, models.ErrNoRecord
	}

	//check if session belongs to the Admin
	user := &models.User{}
	row = DB.QueryRow("SELECT is_admin FROM users WHERE id = ?", session.User.Id)
	err = row.Scan(&user.IsAdmin)

	if err != nil || user.IsAdmin != 1 {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("SELECT is_admin err - %v", err)
			return nil, models.ErrUnauthorized
		}
		return nil, models.ErrUnauthorized
	}

	t := time.Unix(int64(createdDate), 0) // time.Time
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
