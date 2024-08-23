package config_test

import (
	"testing"

	"github.com/Kurichi/go-template/internal/app/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want *config.Config
		err  error
	}{
		{
			name: "Valid configuration",
			want: &config.Config{
				Port: 8080,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := config.New()
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.err, err)
		})
	}
}
