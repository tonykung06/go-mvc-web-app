package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	// "fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Member struct {
	id       int
	email    string
	password string
}

type Session struct {
	id        int
	memberId  int
	sessionId string
}

func (this *Session) Id() int {
	return this.id
}
func (this *Session) MemberId() int {
	return this.memberId
}
func (this *Session) SessionId() string {
	return this.sessionId
}

func (this *Member) Id() int {
	return this.id
}
func (this *Member) Email() string {
	return this.email
}
func (this *Member) Password() string {
	return this.password
}
func (this *Member) SetEmail(val string) {
	this.email = val
}
func (this *Member) SetPassword(val string) {
	this.password = val
}

func GetMember(email, password string) (*Member, error) {
	session, err := GetDbConnection()
	if err == nil {
		defer session.Close()
		pwd := sha256.Sum256([]byte(password))
		result := Member{}
		collection := getCollection(session, "members")
		memberDocument := bson.D{}

		err = collection.Find(bson.M{
			"email":    email,
			"password": hex.EncodeToString(pwd[:]),
		}).One(&memberDocument)

		if err == nil {
			result.SetEmail(memberDocument.Map()["email"].(string))
			result.SetPassword(memberDocument.Map()["password"].(string))
			return &result, nil
		} else {
			return &result, errors.New("Member not found for " + email)
		}
	} else {
		return &Member{}, errors.New("Unable to get db connection")
	}
}

func CreateSession(member *Member) (*Session, error) {
	sessionId := sha256.Sum256([]byte(member.Email() + time.Now().UTC().Format("12:00:00")))
	result := Session{
		memberId:  member.Id(),
		sessionId: hex.EncodeToString(sessionId[:]),
	}

	dbSession, err := GetDbConnection()
	if err == nil {
		defer dbSession.Close()
		sessionCollection := getCollection(dbSession, "sessions")
		err = sessionCollection.Insert(map[string]interface{}{
			"memberId":  result.MemberId(),
			"sessionId": result.SessionId(),
		})

		if err == nil {
			return &result, err
		} else {
			return &Session{}, errors.New("Unable to save session to database")
		}
	} else {
		return &result, errors.New("Unable to get db connection")
	}
}
