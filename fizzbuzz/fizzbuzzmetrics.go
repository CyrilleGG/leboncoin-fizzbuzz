package fizzbuzz

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/cyrillegg/leboncoin-fizzbuzz/database"
)




//		Defining Metrics struct
type Metrics struct {
	IpAddress		string
	FirstInt		int
	SecondInt		int
	Limit			int
	FirstString		string
	SecondString	string
	ParamsHash		string
}




//		Function to hash data
func (m *Metrics) hashParams() error {

	//		Converting integers into strings
	var int1 = strconv.Itoa(m.FirstInt)
	var int2 = strconv.Itoa(m.SecondInt)
	var limit = strconv.Itoa(m.Limit)

	//		Preparing params string before hashing
	var params = "first_int: " + int1 + "second_int: " + int2 + "limit: " + limit + "first_string: " + m.FirstString + "second_string: " + m.SecondString

	//		Hashing and converting hash to string
	var hParams = md5.New()
	hParams.Write([]byte(params))
	var hexaParams = hParams.Sum(nil)
	var hashParamsStr = hex.EncodeToString(hexaParams)

	//		Assigning result
	m.ParamsHash = hashParamsStr

	return nil
}




//		Function to insert metrics into DB
func (m *Metrics) Insert() error {
	err := m.hashParams()
	if err != nil {
		return err
	}

	var query = `
INSERT INTO fizzbuzz_queries ("id", "ip_address", "time", "first_int", "second_int", "limit", "first_string", "second_string", "params_hash")
	VALUES (DEFAULT, $1, DEFAULT, $2, $3, $4, $5, $6, $7)
`

	_, err = database.Open().Exec(query, m.IpAddress, m.FirstInt, m.SecondInt, m.Limit, m.FirstString, m.SecondString, m.ParamsHash)
	if err != nil {
		return err
	}

	return nil
}