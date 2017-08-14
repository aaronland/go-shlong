package charset

import (
       "github.com/thisisaaronland/go-shlong"
       "github.com/thisisaaronland/go-shlong/utils"       
)

type DefaultCharset {
     shlong.Charset
     chars string	
}

func NewDefaultCharset() (shlong.Charset, error) {

     cs := DefaultCharset{
     	chars: "qwrtypsdfghjklzxcvbnm0123456789",
     }

     return &cs, nil
}

func (cs *DefaultCharset) GenerateId(length int) (string, error){
     return RandomFromCharset(cs, length)
}

func (cs *DefaultCharset) Characters() string {
     return cs.chars
}
