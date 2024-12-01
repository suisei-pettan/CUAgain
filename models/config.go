package models

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type CUAgain struct {
	Port                         int    `yaml:"port"`
	Password                     string `yaml:"password"`
	AssetsProxy                  bool   `yaml:"assets-proxy"`
	AssetsCache                  bool   `yaml:"assets-cache"`
	LoginAuth                    bool   `yaml:"login-auth"`
	LoginPassword                string `yaml:"login-password"`
	LoginTimeout                 int    `yaml:"login-timeout"`
	GetIpMethod                  string `yaml:"get-ip-method"`
	EnableGlobalHolostarMovement bool   `yaml:"enable-global-holostar-movement"`
	RemoveAngleLimit             bool   `yaml:"remove-angle-limit"`
	RSAPublicKeyPath             string `yaml:"rsa-public-key-path"`
	RSAPrivateKeyPath            string `yaml:"rsa-private-key-path"`
}
type Hololy struct {
	VersionBypass string `yaml:"version-bypass"`
}

type Config struct {
	CUAgain CUAgain `yaml:"cuagain"`
	Hololy  Hololy  `yaml:"hololy"`
}

func defaultConfig() Config {
	return Config{
		CUAgain: CUAgain{
			Port:                         8080,
			Password:                     "114514",
			AssetsProxy:                  false,
			AssetsCache:                  false,
			LoginAuth:                    true,
			LoginPassword:                "Hololive",
			LoginTimeout:                 2880,
			GetIpMethod:                  "0",
			EnableGlobalHolostarMovement: false,
			RemoveAngleLimit:             true,
			RSAPublicKeyPath:             "rsa/rsa_public_key.pem",
			RSAPrivateKeyPath:            "rsa/rsa_private_key.pem",
		},
		Hololy: Hololy{
			VersionBypass: "2.4.8",
		},
	}
}

func GetConfig() Config {
	config := defaultConfig()
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Println("config.yaml not found, using default config")
		return config
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Println("config.yaml parse error, using default config")
		return config
	}
	return config
}
