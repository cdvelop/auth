package auth

type db interface {
	ReadObject(table_name string, where_fields map[string]string) map[string]string
}
