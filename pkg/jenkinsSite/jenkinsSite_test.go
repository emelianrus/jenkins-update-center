package jenkinsSite

import (
	"reflect"
	"testing"
)

func TestNewJenkinsSite(t *testing.T) {
	tests := []struct {
		name string
		want JenkinsSite
	}{
		{
			name: "init test",
			want: NewJenkinsSite(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := NewJenkinsSite(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJenkinsSite() = %v, want %v", got, tt.want)
			}
		})
	}
}
