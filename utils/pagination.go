package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/qimpl/housing/models"
)

// DecodeCursor yo
func DecodeCursor(encodedCursor string) (*models.Cursor, error) {
	byteArray, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return nil, err
	}

	cursorArray := strings.Split(string(byteArray), ",")
	if len(cursorArray) != 2 {
		return nil, errors.New("Given cursor is invalid")
	}

	createdAt, err := time.Parse(time.RFC3339Nano, cursorArray[0])
	if err != nil {
		return nil, err
	}

	return &models.Cursor{
		ID:        cursorArray[1],
		CreatedAt: createdAt,
	}, nil
}

// EncodeCursor yo
func EncodeCursor(uuid uuid.UUID, t time.Time) string {
	return base64.StdEncoding.EncodeToString(
		[]byte(
			fmt.Sprintf(
				"%s,%s",
				t.Format(time.RFC3339Nano),
				uuid.String(),
			),
		),
	)
}

// WritePaginationHeaders YO
func WritePaginationHeaders(w http.ResponseWriter, paginationHeader *models.PaginationHeaders) {
	w.Header().Set("Link", strings.Join(paginationHeader.Link, " "))
	w.Header().Set("Pagination-Cursor", paginationHeader.Cursor)
}
