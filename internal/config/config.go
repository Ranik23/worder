package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)




type Config struct {
	App struct {
		LogLevel       string `yaml:"levelLog"`
		Port           string `yaml:"port"`
		Host		   string `yaml:"host"`
		StoragePath    string `yaml:"storage_path"`
		MigrationsPath string `yaml:"migrations_path"`
	} `yaml:"App"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"DataBase"`
	Dictionary struct {
		Words	 []string `yaml:"words"`
	} `yaml:"Dictionary"`
	Alphabet struct {
		Syllables	[]string	`yaml:"syllables"`
	} `yaml:"Alphabet"`
}


func LoadConfig(path string) (*Config, error) {

	var config Config

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)

	return &config, err
}