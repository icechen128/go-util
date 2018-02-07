package mongo

import "github.com/globalsign/mgo"

// Mongodb encapsulation
type Mongodb struct {
	mgoSession *mgo.Session
	url        string
	user       string
	password   string

	session *mgo.Session
}

// NewMongodb creat new mongodb
func NewMongodb(url string, user string, password string) *Mongodb {
	mgodb := new(Mongodb)
	mgodb.url = url
	mgodb.user = user
	mgodb.password = password
	return mgodb
}

func (m *Mongodb) getSession() *mgo.Session {
	if m.mgoSession == nil {
		info := &mgo.DialInfo{
			Addrs:    []string{m.url},
			Username: m.user,
			Password: m.password,
			Source:   "test",
		}
		var err error
		m.mgoSession, err = mgo.DialWithInfo(info)
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
	if m.session == nil {
		m.session = m.getSession()
	}
	temC := m.session.DB(dbname).C(collection)
	return temC.Insert(document)
}

func (m *Mongodb) Close() {
	if m.session != nil {
		m.session.Close()
	}
}
