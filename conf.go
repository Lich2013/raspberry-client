package main

import (
	"github.com/BurntSushi/toml"
	"fmt"
	"os"
)


type Config struct {
	Schema string `toml:"schema"`
	Host string `toml:"host"`
}

func InitConfig() {
	_, err := toml.DecodeFile("./conf/dev.toml", Conf)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Println(Conf)
	wg.Done()
}