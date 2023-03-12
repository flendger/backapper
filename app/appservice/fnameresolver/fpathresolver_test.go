package fnameresolver

import (
	"backapper/app"
	"testing"
	"time"
)

func Test_resolve(t *testing.T) {
	type args struct {
		curApp *app.App
		fName  string
		fTime  time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test with prefix",
			args: args{
				curApp: &app.App{
					Name:     "test_app",
					FilePath: "/home/app/test_app/filename.jar",
					ArcDir:   "/home/app/archive"},
				fName: "filename.jar",
				fTime: time.Date(2023, 3, 10, 22, 10, 15, 0, time.Local)},
			want: "/home/app/archive/filename_2023-03-10 22_10_15.jar"},
		{
			name: "test without prefix",
			args: args{
				curApp: &app.App{
					Name:     "test_app_2",
					FilePath: "/home/app/test_app/filename",
					ArcDir:   "/home/app/archive"},
				fName: "filename",
				fTime: time.Date(2023, 3, 10, 22, 10, 15, 0, time.Local)},
			want: "/home/app/archive/filename_2023-03-10 22_10_15"},
		{
			name: "test with two pointers",
			args: args{
				curApp: &app.App{
					Name:     "test_app_3",
					FilePath: "/home/app/test_app/filename.app.jar",
					ArcDir:   "/home/app/archive"},
				fName: "filename.app.jar",
				fTime: time.Date(2023, 3, 10, 22, 10, 15, 0, time.Local)},
			want: "/home/app/archive/filename.app_2023-03-10 22_10_15.jar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Resolve(tt.args.curApp, tt.args.fName, tt.args.fTime); got != tt.want {
				t.Errorf("Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
