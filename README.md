# Go-examples :books:
#### :loudspeaker: This GitHub repository contains code examples for those who are new to go programming. 

![This is an image](https://www.softwebsolutions.com/wp-content/uploads/2020/10/golang-Programing.jpg)

IMPLEMENTATIONS

1. **JWT - json-web-token-implementation/jwt.go**

JSON Web Token (JWT) is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. This information can be verified and trusted because it is digitally signed. JWTs can be signed using a secret (with the HMAC algorithm) or a public/private key pair using RSA or ECDSA.
Although JWTs can be encrypted to also provide secrecy between parties, we will focus on signed tokens. Signed tokens can verify the integrity of the claims contained within it, while encrypted tokens hide those claims from other parties. When tokens are signed using public/private key pairs, the signature also certifies that only the party holding the private key is the one that signed it.

For more information, visit the official https://jwt.io

In our JWT implementation in Golang:
- `/login` -> endpoint uses for authentication and authorization client.Client first have to request to /login endpoint with the given username:password with JSON format.
- `/home`  -> for Hello, user.

All response codes using like below:
- 200 - [OK]
- 400 - [Bad Request]
- 401 - [Unauthorized]

