package feature

import "github.com/BlackwonderTF/go-ini/utils"

type Feature interface {
}

func GetFeaturePrefix(line string) string {
	return utils.Match(line, "^[\\s]+")
}
