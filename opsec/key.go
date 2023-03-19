package opsec

import (
	"encoding/hex"
	"log"

	"github.com/epictetus24/go-util/pkg/enc"
)

var (
	SboxKey *[]byte
)

func NewSBoxKey() []byte {
	key, _ := enc.GenKey(32)
	SboxKey = &key
	keystr := hex.EncodeToString(*SboxKey)
	log.Printf("\n[OPSEC] Using the following randkey %s, in future set one in config.\n", keystr)
	return key
}
