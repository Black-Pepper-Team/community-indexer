package core

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/iden3/go-iden3-crypto/babyjub"
)

func TestBuildCredential(t *testing.T) {
	contractId := common.HexToAddress("0x0B306BF915C4d645ff596e518fAf3F9669b97016")
	nftOwner := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	x, ok := new(big.Int).SetString("14302959900069061493388322109831425126447754027481418724221945884581763476571", 10)
	if !ok {
		log.Fatalf("failed to set x")
	}

	y, ok := new(big.Int).SetString("19133337649544864758584673172563393357921375750532468744093033525654718926514", 10)
	if !ok {
		log.Fatalf("failed to set y")
	}

	pk := babyjub.PublicKey(babyjub.Point{
		X: x,
		Y: y,
	})

	credentialID, err := buildCredentialId(
		contractId,
		big.NewInt(0),
		nftOwner,
		pk.X,
		pk.Y,
		5017,
	)
	if err != nil {
		log.Fatalf("failed to build credential id: %v", err)
	}

	hexCredentialID := hex.EncodeToString(credentialID.Bytes())

	fmt.Println(hexCredentialID)
}

func TestMessageEncoding(t *testing.T) {
	message := "Hello World!"

	fmt.Println(hex.EncodeToString(crypto.Keccak256([]byte(message))))

	hashInt := new(big.Int).SetBytes(crypto.Keccak256([]byte(message)))

	mask := new(big.Int).SetBytes([]byte{0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	result := new(big.Int).And(hashInt, mask)

	fmt.Println(result)
}
