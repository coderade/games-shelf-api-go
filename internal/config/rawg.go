package config

type RawgConfig struct {
	ApiKey      string
	ApiEndpoint string
}

func LoadRawgConfig() RawgConfig {
	return RawgConfig{
		ApiKey:      getEnv("RAWG_API_KEY", ""),
		ApiEndpoint: getEnv("RAWG_API_ENDPOINT", "https://api.rawg.io/api"),
	}
}
