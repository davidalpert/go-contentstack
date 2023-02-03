package schema

func StrPtr(v string) *string {
	return &v
}

func BoolPtr(v bool) *bool {
	return &v
}

const timestampFormat = "2006-01-02T15:04:05.000Z"
