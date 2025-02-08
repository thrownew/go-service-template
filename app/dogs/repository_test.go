package dogs

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"pupa/deps"
	"pupa/tests"
)

func TestIntegrationRepository(t *testing.T) {
	box := tests.NewBox(tests.BoxProvider(
		deps.Provide(NewRepository),
	))

	box.RunT(t, "repository", func(t *testing.T, ctx context.Context, r *Repository) {
		t.Run("DogByName", func(t *testing.T) {
			t.Parallel()
			cases := []struct {
				name    string
				err     error
				goodBoy bool
			}{
				{"Buddy", nil, true},
				{"Undefined", ErrorNotFound, false},
				{"Luna", nil, false},
			}
			for _, c := range cases {
				t.Run(c.name, func(t *testing.T) {
					t.Parallel()
					dog, err := r.DogByName(ctx, c.name)
					if c.err != nil {
						require.Error(t, err)
						assert.ErrorIs(t, err, c.err)
						return
					}
					require.NoError(t, err)
					assert.Equal(t, c.goodBoy, dog.GoodBoy)
				})
			}
		})
	})
}
