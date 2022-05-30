package repository

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_setupDBInstance(t *testing.T) {
	tests := []struct {
		name    string
		want    *sql.DB
		wantErr bool
	}{
		{
			name:    "silver",
			want:    GetInstance(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setupDBInstance()
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NotNil(t, got)
		})
	}
}
