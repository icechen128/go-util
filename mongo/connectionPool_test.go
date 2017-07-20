package mongo

import (
	"testing"
)

func TestMongodb_insert(t *testing.T) {
	testTem := struct {
		Name string `bson:"username"`
		Age  int64  `bson:"年龄"`
	}{
		Name: "icechen",
		Age:  18,
	}
	type args struct {
		dbname     string
		collection string
		document   interface{}
	}
	tests := []struct {
		name    string
		m       *Mongodb
		args    args
		wantErr bool
	}{
		{
			name:    "test0",
			m:       NewMongodb(""),
			args:    args{dbname: "test", collection: "test0", document: testTem},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.insert(tt.args.dbname, tt.args.collection, tt.args.document); (err != nil) != tt.wantErr {
				t.Errorf("Mongodb.insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
