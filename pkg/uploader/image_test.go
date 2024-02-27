package uploader

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_imageService_Resize(t *testing.T) {
	type fields struct {
		host      string
		secretKey string
	}
	type args struct {
		path   string
		width  int32
		height int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				host:      "https://lof-images-dev.s3corp.vn",
				secretKey: "123456",
			},
			args: args{
				path:   "image-upload/2023/6/21/1687317088922080563_uefn3w44_title-trangroi-nenden_815_255.png",
				width:  800,
				height: 600,
			},
			want: "https://lof-images-dev.s3corp.vn/image/api/v1/resize/image-upload/2023/6/21/1687317088922080563_uefn3w44_title-trangroi-nenden_815_255.png?width=800&height=600&hash=0007fe67c45b581b8a5d03c44fb82637",
		},
		{
			name: "empty",
			fields: fields{
				host:      "https://lof-images-dev.s3corp.vn",
				secretKey: "123456",
			},
			args: args{
				path:   "",
				width:  800,
				height: 600,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &imageService{
				host:      tt.fields.host,
				secretKey: tt.fields.secretKey,
			}
			require.Equal(t, i.Resize(tt.args.path, tt.args.width, tt.args.height), tt.want)
		})
	}
}
