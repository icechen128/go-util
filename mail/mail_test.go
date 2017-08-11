package mail

import "testing"

func Test_SendMail(t *testing.T) {
	type args struct {
		to      string
		subject string
		body    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				to:      "icechen128@gmail.com",
				subject: "mail test is running",
				body:    `<html><h3>test running</h3></html>`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMail(tt.args.to, tt.args.subject, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("sendMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
