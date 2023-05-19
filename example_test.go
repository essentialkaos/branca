package branca

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//                  MIT License <https://opensource.org/licenses/MIT>                 //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNewBranca() {
	key := "mysupppadupppasecretkeyforbranca"
	brc, err := NewBranca([]byte(key))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	payload := "MySecretData"
	token, err := brc.EncodeToString([]byte(payload))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Token: %s\n", token)
}

func ExampleBranca_Encode() {
	key := "mysupppadupppasecretkeyforbranca"
	brc, err := NewBranca([]byte(key))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	payload := "MySecretData"
	token, err := brc.Encode([]byte(payload))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Token: %v\n", token)
}

func ExampleBranca_EncodeToString() {
	key := "mysupppadupppasecretkeyforbranca"
	brc, err := NewBranca([]byte(key))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	payload := "MySecretData"
	token, err := brc.EncodeToString([]byte(payload))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Token: %s\n", token)
}

func ExampleBranca_Decode() {
	key := "mysupppadupppasecretkeyforbranca"
	brc, err := NewBranca([]byte(key))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	data := []byte{186, 92, 40, 2, 172, 169, 93, 187, 60, 216, 139, 79, 89,
		25, 68, 191, 235, 113, 237, 55, 133, 168, 158, 255, 160, 36, 98, 222,
		110, 242, 182, 153, 143, 206, 44, 141, 59, 46, 81, 124, 114, 25, 117,
		85, 156, 170, 204, 175, 164, 57, 5, 235, 56, 1, 115, 5, 222}

	token, err := brc.Decode(data)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Payload: %s\n", string(token.Payload()))
	// Output: Payload: MySecretData
}

func ExampleBranca_DecodeString() {
	key := "mysupppadupppasecretkeyforbranca"
	brc, err := NewBranca([]byte(key))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	data := "87yoI2tNmtPGYxQMYExUw6Yn0zRJccwIiMZxAQ7OBNoLl2P2stmAfD1BLvHOIdwmjGIWxnLrNmHLG"
	token, err := brc.DecodeString(data)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Payload: %s\n", string(token.Payload()))
	// Output: Payload: MySecretData
}

func ExampleToken_Payload() {
	key := "mysupppadupppasecretkeyforbranca"
	brc, err := NewBranca([]byte(key))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	data := "87yoI2tNmtPGYxQMYExUw6Yn0zRJccwIiMZxAQ7OBNoLl2P2stmAfD1BLvHOIdwmjGIWxnLrNmHLG"
	token, err := brc.DecodeString(data)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Payload: %s\n", string(token.Payload()))
	// Output: Payload: MySecretData
}

func ExampleToken_Timestamp() {
	key := "mysupppadupppasecretkeyforbranca"
	brc, err := NewBranca([]byte(key))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	data := "87yoI2tNmtPGYxQMYExUw6Yn0zRJccwIiMZxAQ7OBNoLl2P2stmAfD1BLvHOIdwmjGIWxnLrNmHLG"
	token, err := brc.DecodeString(data)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Timestamp: %v\n", token.Timestamp())
}

func ExampleToken_IsExpired() {
	key := "mysupppadupppasecretkeyforbranca"
	brc, err := NewBranca([]byte(key))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	data := "87yoI2tNmtPGYxQMYExUw6Yn0zRJccwIiMZxAQ7OBNoLl2P2stmAfD1BLvHOIdwmjGIWxnLrNmHLG"
	token, err := brc.DecodeString(data)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Token expired: %t", token.IsExpired(0))
	// Output: Token expired: true
}
