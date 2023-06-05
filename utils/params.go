package utils

type QueryParams struct {
	QQQ string
}

func ParseQueryParams(queryParams map[string][]string) (*QueryParams, []string) {
	params := QueryParams{
		QQQ: " ",
	}
	var errStr []string

	for key, value := range queryParams {
		if key == "q" {
			params.QQQ = value[0]
			continue
		}
	}

	return &params, errStr
}
