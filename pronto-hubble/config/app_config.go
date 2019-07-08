package config

func GetDbConfig() DbConfig {
	return DbConfig{ //read it from file or env
		Host:     "127.0.0.1",
		Port:     27017,
		User:     "pronto",
		Password: "pronto",
		Database: "prontodb",
	}
}
