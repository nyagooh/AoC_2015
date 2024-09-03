package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	secretKey := "bgvyzdsv"
	number := 1
	CryptoHash(secretKey, number)

}

func CryptoHash(str string, num int) {
	//global variables
	var fiveZeroNum, sixZeroNum int
	foundFive, foundSix := false, false
	for {
		//converts str and num to a string and concatenates them
		input := fmt.Sprintf("%s%d", str, num)
        //convert the intup to a []byte slice and gets the secret hash code 
		hash := md5.Sum([]byte(input))
        //converts the hash code to a secret hexdecimal code
		hashHex := hex.EncodeToString(hash[:])
        //checks which number combination and string starts with 5 zereos and prints the combination
		if !foundFive && hashHex[:5] == "00000" {
			foundFive = true
			fiveZeroNum = num

		}
		if !foundSix && hashHex[:6] == "000000" {
			foundSix = true
			sixZeroNum = num

		}
		if foundFive && foundSix {
			break
		}
		num++
	}
fmt.Printf("%d\n%d\n",fiveZeroNum,sixZeroNum)
}
