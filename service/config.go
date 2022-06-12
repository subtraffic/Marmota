package service

import (
	"Marmota/domain/model"
	"Marmota/utils/language"
	"fmt"

	"github.com/spf13/cobra"
)

type callbackfunc func(conf model.Marmota, content interface{})

func UpdateConf(cmd *cobra.Command, args []string) {
	doUpdateConf(cmd, "configfile", func(conf model.Marmota, content interface{}) { conf.ConfigFile = fmt.Sprintf("%s", content) })
	doUpdateConf(cmd, "datafile", func(conf model.Marmota, content interface{}) { conf.DataFile = fmt.Sprintf("%s", content) })
	doUpdateConf(cmd, "host", func(conf model.Marmota, content interface{}) { conf.ServerHost = fmt.Sprintf("%s", content) })
	doUpdateConf(cmd, "port", func(conf model.Marmota, content interface{}) { conf.Port = fmt.Sprintf("%s", content) })
}

func doUpdateConf(cmd *cobra.Command, label string, callback callbackfunc) {
	conf := model.GetMarmotaConf()
	confConfigFile, err := cmd.Flags().GetString(label)
	if err != nil {
		panic(language.SystemError)
	} else {
		callback(*conf, confConfigFile)
	}
}
