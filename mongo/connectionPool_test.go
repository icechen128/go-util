package mongo

import (
	"sync"
	"testing"
)

func TestMongodb_insert(t *testing.T) {
	testTem := struct {
		Name string `bson:"username"`
		Age  int64  `bson:"年龄"`
		Test int64  `bson:"test"`
	}{
		Name: "icechen1",
		Age:  18,
		Test: 1,
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
			m:       NewMongodb("www.icechen.cn:27017", "test0", "3541213"),
			args:    args{dbname: "test", collection: "test0", document: testTem},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := &sync.WaitGroup{}
			for i := 0; i < 1000; i++ {
				wg.Add(1)
				go func() {
					if err := tt.m.insert(tt.args.dbname, tt.args.collection, tt.args.document); (err != nil) != tt.wantErr {
						t.Errorf("Mongodb.insert() error = %v, wantErr %v", err, tt.wantErr)
					}
					wg.Done()
				}()
			}
			wg.Wait()
			tt.m.Close()
		})
	}
}
