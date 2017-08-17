package g

import (
	"github.com/BurntSushi/toml"
	"fmt"
	"os"
)

func InitConfig() {
	_, err := toml.DecodeFile("./conf/dev.toml", Conf)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}

func Init()  {
	InitConfig()
}
