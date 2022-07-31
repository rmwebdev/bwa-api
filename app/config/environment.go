package config

// Environment variables
var Environment map[string]interface{} = map[string]interface{}{
	"port":              9000,
	"endpoint":          "/api/v1/master",
	"environment":       "development",
	"db_host":           "postgres",
	"db_port":           5432,
	"db_user":           "postgres",
	"db_pass":           "postgres",
	"db_name":           "postgres",
	"db_table_prefix":   "",
	"redis_host":        "redis",
	"redis_port":        6379,
	"redis_pass":        "",
	"redis_index":       0,
	"prefork":           false,
	"language":          "en",
	"aes":               "AES_CIPHER_SECRET_KEY_KUDU_32BIT",
	"salt":              "SALT",
	"agent_id":          "8d5fb953-73af-4a07-b450-f8641b40e383", // Dummy Agent ID
	"user_id":           "db24d53c-7d36-4770-8598-dc36174750af", // Dummy User ID
	"cache_ttl_seconds": 60 * 60,
}
