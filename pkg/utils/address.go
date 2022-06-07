package utils

import (
	"github.com/civet148/log"
	"github.com/ethereum/go-ethereum/common"
)

func EtherAddr(addr string) (common.Address, error) {
	oa, err := common.NewMixedcaseAddressFromString(addr)
	if err != nil {
		log.Errorf("address [%s] can't convert to ethereum common address, error [%s]", addr, err.Error())
		return common.Address{}, err
	}
	return oa.Address(), nil
}
