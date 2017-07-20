package mongo

import (
	"gopkg.in/mgo.v2"
)

// Mongodb encapsulation
type Mongodb struct {
	mgoSession *mgo.Session
	url        string
}

// NewMongodb creat new mongodb
func NewMongodb(url string) *Mongodb {
	mgodb := new(Mongodb)
	mgodb.url = url
	return mgodb
}

func (m *Mongodb) getSession() *mgo.Session {
	if m.mgoSession == nil {
		var err error
		m.mgoSession, err = mgo.Dial(m.url)
		if err != nil {
			panic(err)
		}
	}
	return m.mgoSession.Clone()
}

func (m *Mongodb) db(dbname string) *mgo.Database {
	return m.getSession().DB(dbname)
}

func (m *Mongodb) c(dbname string, collection string) *mgo.Collection {
	return m.db(dbname).C(collection)
}

func (m *Mongodb) insert(dbname string, collection string, document interface{}) error {
	temSession := m.getSession()
	defer temSession.Close()
	temC := temSession.DB(dbname).C(collection)
	return temC.Insert(document)
}
