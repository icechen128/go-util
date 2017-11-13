package config

import "testing"

func Test_readFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test.ini",
			args: args{
				filename: "test.ini",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := readFile(tt.args.filename)
			t.Logf("config is %v", defConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
