package ssh

import "testing"

func TestPushCMD(t *testing.T) {
	type args struct {
		user     string
		password string
		host     string
		port     int
		cmds     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				user:     "work",
				password: "wwrglb",
				host:     "180.97.80.42",
				port:     22,
				cmds: []string{"cd icechen/idatacollector",
					"pkill idatacollector",
					"/usr/bin/nohup ./idatacollector > ./nohup.log &",
					"ps -e | grep idata"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PushCMD(tt.args.user, tt.args.password, tt.args.host, tt.args.port, tt.args.cmds...); (err != nil) != tt.wantErr {
				t.Errorf("PushCMD() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
