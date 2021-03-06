/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package mongo

import (
	"gopkg.in/mgo.v2"
)

func databaseName() string {
	return "SiCo"
}

func NewDial() (*mgo.Session, error) {
	return &mgo.Session{}, nil
}

func InitDial(address, username, password string) (*mgo.Session, error) {
	if username != "" && password != "" {
		return mgo.Dial(username + ":" + password + "@" + address)
	}
	return mgo.Dial(address)
}
