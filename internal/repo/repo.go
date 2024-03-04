package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/rs/zerolog"

	"github.com/mrbelka12000/webhooker/internal/models"
)

type WebHooker struct {
	db *sql.DB
	lg zerolog.Logger
}

func NewWebHooker(db *sql.DB, log zerolog.Logger) *WebHooker {
	return &WebHooker{
		db: db,
		lg: log,
	}
}

func (wh *WebHooker) Create(ctx context.Context, data *models.Data) (err error) {

	_, err = wh.db.ExecContext(ctx, `
	INSERT INTO web_hooks 
		(callback_url,http_method, params, body, end_time) 
	VALUES 
	    ($1,$2,$3,$4,$5)`,
		data.CallbackURL,
		data.Method,
		data.Params,
		data.Body,
		data.EndTime,
	)

	return
}

func (wh *WebHooker) List(ctx context.Context, pars models.DataListPars) ([]models.Data, error) {

	var err error

	var filterValues []interface{}
	querySelect := `
	SELECT callback_url,http_method, params, body, end_time
`
	queryFrom := ` FROM web_hooks `
	queryWhere := ` WHERE 1 = 1`
	queryOffset := fmt.Sprintf(" OFFSET %v", pars.Offset)
	queryLimit := fmt.Sprintf(" LIMIT %v", pars.Limit)
	queryOrderBy := " order by end_time asc"

	rows, err := wh.db.QueryContext(ctx, querySelect+queryFrom+queryWhere+queryOrderBy+queryOffset+queryLimit, filterValues...)
	if err != nil {
		return nil, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	var result []models.Data

	for rows.Next() {
		var d models.Data

		err := rows.Scan(
			&d.CallbackURL,
			&d.Method,
			&d.Params,
			&d.Body,
			&d.EndTime,
		)
		if err != nil {
			return nil, fmt.Errorf("rows scan: %w", err)
		}

		result = append(result, d)
	}

	return result, nil
}

func reverseWords(s string) string {
	arr := strings.Split(s, " ")

	var result strings.Builder
	for i := 0; i < len(arr); i++ {
		revStr := make([]byte, len(arr[i]))

		for j := 0; j < len(arr[i]); j++ {
			revStr[j] = arr[i][len(arr[i])-j]
		}

		result.Write(revStr)
		if i != len(arr)-1 {
			result.Write([]byte(" "))
		}
	}

	return result.String()
}
