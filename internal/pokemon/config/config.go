package config

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	defCfg      map[string]string
	initialized = false
)

// initialize this configuration
func initialize() {
	viper.SetConfigFile(".env")
	viper.SetEnvPrefix("POKEMON")
	viper.ReadInConfig()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	defCfg = make(map[string]string)

	defCfg["server.log.level"] = viper.GetString(`POKEMON_SERVER_LOG_LEVEL`) // valid values are trace, debug, info, warn, error, fatal

	defCfg["server.address"] = fmt.Sprintf("%s:%s", viper.GetString(`POKEMON_SERVER_HOST`), viper.GetString(`POKEMON_SERVER_PORT`))
	defCfg["server.host"] = viper.GetString(`POKEMON_SERVER_HOST`)
	defCfg["server.port"] = viper.GetString(`POKEMON_SERVER_PORT`)
	defCfg["db.mysql.host"] = viper.GetString(`POKEMON_DB_MYSQL_HOST`)
	defCfg["db.mysql.port"] = viper.GetString(`POKEMON_DB_MYSQL_PORT`)
	defCfg["db.mysql.user"] = viper.GetString(`POKEMON_DB_MYSQL_USER`)
	defCfg["db.mysql.password"] = viper.GetString(`POKEMON_DB_MYSQL_PASSWORD`)
	defCfg["db.mysql.name"] = viper.GetString(`POKEMON_DB_MYSQL_NAME`)
	defCfg["db.mysql.maxidle"] = viper.GetString(`POKEMON_DB_MYSQL_MAXIDLE`)
	defCfg["db.mysql.maxopen"] = viper.GetString(`POKEMON_DB_MYSQL_MAXOPEN`)

	for k := range defCfg {
		err := viper.BindEnv(k)
		if err != nil {
			log.Errorf("Failed to bind env \"%s\" into configuration. Got %s", k, err)
		}
	}

	initialized = true
}

// SetConfig put configuration key value
func SetConfig(key, value string) {
	viper.Set(key, value)
}

// Get fetch configuration as string value
func Get(key string) string {
	if !initialized {
		initialize()
	}
	ret := viper.GetString(key)
	if len(ret) == 0 {
		if ret, ok := defCfg[key]; ok {
			return ret
		}
		log.Debugf("%s config key not found", key)
	}
	return ret
}

// GetBoolean fetch configuration as boolean value
func GetBoolean(key string) bool {
	if len(Get(key)) == 0 {
		return false
	}
	b, err := strconv.ParseBool(Get(key))
	if err != nil {
		panic(err)
	}
	return b
}

// GetInt fetch configuration as integer value
func GetInt(key string) int {
	if len(Get(key)) == 0 {
		return 0
	}
	i, err := strconv.ParseInt(Get(key), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

// GetFloat fetch configuration as float value
func GetFloat(key string) float64 {
	if len(Get(key)) == 0 {
		return 0
	}
	f, err := strconv.ParseFloat(Get(key), 64)
	if err != nil {
		panic(err)
	}
	return f
}

// Set configuration key value
func Set(key, value string) {
	defCfg[key] = value
}
