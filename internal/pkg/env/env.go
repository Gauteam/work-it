package env

import "os"

var (
	DbName     = resolveFlag("DB_NAME", "workit")
	DbHost     = resolveFlag("DB_HOST", "localhost")
	DbPort     = resolveFlag("DB_PORT", "5432")
	DbUser     = resolveFlag("DB_USER", "workituser")
	DbPassword = resolveFlag("DB_PASS", "workitpass")
)

func resolveFlag(envName string, defaultVal string) *string {
	val, ok := os.LookupEnv(envName)
	if !ok {
		return &defaultVal
	}
	return &val
}
