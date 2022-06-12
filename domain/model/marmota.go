package model

import (
	"Marmota/utils/language"
	"fmt"
	"os"
)

type Marmota struct {
	ServerHost string
	Port       string
	UI         bool
	Server     bool
	ConfigFile string
	DataFile   string
}

var marmotaConf *Marmota

func (m *Marmota) Save(path string) {
	if len(path) == 0 && len(m.ConfigFile) == 0 {
		fmt.Fprintf(os.Stderr, language.FilePathError)
		return
	}

}

func init() {
	if marmotaConf == nil {
		marmotaConf = &Marmota{}
	}
}

func GetMarmotaConf() *Marmota {
	if marmotaConf == nil {
		marmotaConf = &Marmota{}
	}
	return marmotaConf
}
