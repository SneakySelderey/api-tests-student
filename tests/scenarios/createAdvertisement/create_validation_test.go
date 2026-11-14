package createAdvertisement

import (
	"net/http"
	"strings"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/tests/fixtures"
)

func TestCreateAdvertisementInvalidParameters(t *testing.T) {
	owner := fixtures.Login(t)
	cases := []struct {
		name    string
		field   string
		value   string
		omit    bool
		photos  int
		message string
	}{
		{"нет title", "title", "", true, 1, "title is required"},
		{"пустой title", "title", "", false, 1, "title is required"},
		{"title длиннее 50", "title", strings.Repeat("a", 51), false, 1, "title must not exceed 50 characters provided 51"},
		{"нет description", "description", "", true, 1, "description is required"},
		{"пустой description", "description", "", false, 1, "description is required"},
		{"description длиннее 500", "description", strings.Repeat("a", 501), false, 1, "description must be between 1 and 500 characters"},
		{"отрицательная цена", "price", "-1", false, 1, "price must be between 0 and 1000000 inclusive"},
		{"цена выше максимума", "price", "1000001", false, 1, "price must be between 0 and 1000000 inclusive"},
		{"цена не число", "price", "abc", false, 1, "cannot convert string value abc to float"},
		{"нулевое количество", "quantity", "0", false, 1, "quantity must be between 1 and 100 inclusive"},
		{"отрицательное количество", "quantity", "-1", false, 1, "quantity must be between 1 and 100 inclusive"},
		{"количество выше максимума", "quantity", "101", false, 1, "quantity must be between 1 and 100 inclusive"},
		{"количество не число", "quantity", "abc", false, 1, "cannot convert string value to int"},
		{"нет фотографий", "", "", false, 0, "at least 1 photo is required"},
		{"больше трёх фотографий", "", "", false, 4, "maximum 3 photos allowed"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := fixtures.NewCreateRequest(t)
			if tc.field != "" {
				data.Fields[tc.field] = tc.value
			}
			if tc.omit {
				delete(data.Fields, tc.field)
			}
			if tc.photos == 0 {
				data.Photos = nil
			} else {
				for len(data.Photos) < tc.photos {
					data.Photos = append(data.Photos, data.Photos[0])
				}
			}
			err := advertisements.RejectCreateAdvertisement(t, owner.Token, data, http.StatusBadRequest)
			assertions.APIError(t, err, "invalid_request", tc.message)
		})
	}
}
