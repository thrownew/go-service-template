package dogs

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	srv "pupa/servers/grpc/pupa"
	"pupa/tests"
)

func TestIntegrationServer(t *testing.T) {
	box := tests.NewBox()
	box.RunT(t, "server", func(t *testing.T, ctx context.Context, cli srv.DogServiceClient) {
		t.Run("DogIsGoodBoyV1", func(t *testing.T) {
			t.Parallel()
			cases := []struct {
				name      string
				errorCode codes.Code
				goodBoy   bool
			}{
				{"Buddy", codes.OK, true},
				{"Undefined", codes.NotFound, false},
				{"Luna", codes.OK, false},
			}
			for _, c := range cases {
				t.Run(c.name, func(t *testing.T) {
					t.Parallel()
					resp, err := cli.DogIsGoodBoyV1(ctx, &srv.DogIsGoodBoyV1Request{Name: c.name})
					if c.errorCode != codes.OK {
						require.Error(t, err)
						assert.Equal(t, c.errorCode, status.Code(err))
						return
					}
					require.NoError(t, err)
					assert.Equal(t, c.goodBoy, resp.GetIsGoodBoy())
				})
			}
		})
	})
}
