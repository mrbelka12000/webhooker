package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/mrbelka12000/webhooker/internal/models"
)

func TestApp_CreateWebHook(t *testing.T) {
	t.Skip("dev purpose only")

	ctx := context.Background()

	cases := []struct {
		name string
		req  *models.Data

		isError bool
	}{
		{
			name: "ok",
			req: &models.Data{
				CallbackURL: "https://google.com",
				Body: `
					{"a":"b"}
				`,
				EndTime: time.Now().AddDate(0, 2, 0),
			},
			isError: false,
		},
		{
			name: "error, invalid end time",
			req: &models.Data{
				CallbackURL: "https://google.com",
				EndTime:     time.Now().AddDate(0, -2, 0),
			},
			isError: true,
		},
		{
			name: "error, can not parse url",
			req: &models.Data{
				CallbackURL: "https://goo gle.com",
				EndTime:     time.Now().AddDate(0, 2, 0),
			},
			isError: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			err := app.uc.Create(ctx, tc.req)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
