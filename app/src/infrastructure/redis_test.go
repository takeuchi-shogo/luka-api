package infrastructure

import (
	"reflect"
	"testing"

	redis "github.com/go-redis/redis/v9"
)

func TestRedis_Get(t *testing.T) {
	type fields struct {
		RDB *redis.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue interface{}
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{redis.NewClient(&redis.Options{
				Addr:     "localhost:6379",
				Password: "",
				DB:       0,
			})},
			args:      args{"test1"},
			wantValue: "test1",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				RDB: tt.fields.RDB,
			}
			gotValue, err := r.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Redis.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Redis.Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
