package configs

func Load() {
	_LoadEnvironmentConf()
	_LoadRedisConf()
}
