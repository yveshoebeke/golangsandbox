package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
//	"gopkg.in/mgo.v2/bson"
	"fmt"
)

// MockSession satisfies Session and act as a mock of *mgo.session.
type MockSession struct{}

// NewMockSession mock NewSession.
func NewMockSession() Session {
	return MockSession{}
}

// Close mocks mgo.Session.Close().
func (fs MockSession) Close() {}

// DB mocks mgo.Session.DB().
func (fs MockSession) DB(name string) DataLayer {
	mockDatabase := MockDatabase{}
	return mockDatabase
}

// MockDatabase satisfies DataLayer and act as a mock.
type MockDatabase struct{}

// C mocks mgo.Database(name).Collection(name).
func (db MockDatabase) C(name string) Collection {
	return MockCollection{}
}

// MockCollection satisfies Collection and act as a mock.
type MockCollection struct{}

// Find mock.
func (fc MockCollection) Find(query interface{}) Query {
	return nil
}

// Count mock.
func (fc MockCollection) Count() (n int, err error) {
	return 10, nil
}

// Insert mock.
func (fc MockCollection) Insert(docs ...interface{}) error {
	return nil
}

// Remove mock.
func (fc MockCollection) Remove(selector interface{}) error {
	return nil
}

// Update mock.
func (fc MockCollection) Update(selector interface{}, update interface{}) error {
	return nil
}

// GetMyDocuments mock.
func (fc MockCollection) GetMyDocuments() ([]interface{}, error) {
	fmt.Println("getmydocuments")
	var documents []interface{}

	content, _ := ioutil.ReadFile(
		os.Getenv("GOPATH") + "/config/config.json")

		fmt.Println(content)

		json.Unmarshal(content, &documents)
/*
	err := c.Find(bson.M{}).All(&documents)
	if err != nil {
		return nil, err
	}
*/
	return documents, nil
}

// MockQuery satisfies Query and act as a mock.
type MockQuery struct{}

// All mock.
func (fq MockQuery) All(result interface{}) error {
	return nil
}

// One mock.
func (fq MockQuery) One(result interface{}) error {
	return nil
}

// Distinct mock.
func (fq MockQuery) Distinct(field string, result interface{}) error {
	return nil
}
