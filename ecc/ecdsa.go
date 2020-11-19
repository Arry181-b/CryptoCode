package ecc

import (
	"BCP/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

//-------------------生成私钥和公钥迷失对===============
func GenerageECDSAKey() (*ecdsa.PrivateKey, error) {
	curve := elliptic.P256()
	pri,err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return  nil, err
	}
	return pri, nil
}

func ECDSASign(pri *ecdsa.PrivateKey, data []byte) (*big.Int,*big.Int, error) {
	hash := utils.SHA256Hash(data)
	r, s, err := ecdsa.Sign(rand.Reader, pri, hash)
	if err != nil {
		return nil, nil, err
	}
	return r,s,  nil
}


func ECDSAVerify(pub ecdsa.PublicKey, data []byte, r *big.Int, s *big.Int)