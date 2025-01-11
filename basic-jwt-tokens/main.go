package main

import (
	"basic-jwt-tokens/authentication"
	"fmt"
)

/*
There are basically 2 ways of authorization
 1. Using Cookies
 2. Using Json web tokens

How do Cookies work:-
 1. User sends the user id password to the server
 2. Server checks the id password and sends the cookie back session id to user and stores this in server memory as well
 3. The user sends a request to the server with the session id. The server does a lookup of this session id and check if they can make a request or not.
 4. The server would send appropriate response to the user

The main disadvantage with Cookies
 1. The server needs to store the session id in server memory and this cookies cannot be pooled between multiple servers
 2. Stealing cookies is very easy for a middle man

How do JWT work:-
 1. User sends the user id password to the server
 2. Server checks the id password and would send a JWT with its information baked into the token. So the server need not store any information with it
 3. The user will send a request to the server with the JWT. The server would check weather the JWT is valid or not and also extract the information baked into it.
 4. The server would send appropriate response to the user

The main advantages with JWT is:-
 1. The server need not store the session id or related information in memory.
 2. JWT tokens are hard to go around due to the information stored inside it and also it can be used with multiple servers as pooling is possible.
    We can extract information from JWT tokens using parsing the JWT and reversing it.
*/

func main() {
	jwtToken, err := authentication.CreateToken("param")

	if err != nil {
		fmt.Println("error -> ", err)
	}

	fmt.Println("JWT Token -> ", jwtToken)

	token, err := authentication.VerifyToken(jwtToken)

	if err != nil {
		fmt.Println("error => ", err)
	}

	fmt.Printf("Token -> %+v \n", token.Claims)
}
