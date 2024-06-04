![Alt text](image.png)

jwt 主要由3个部分组成:
header.payload.signature

header:
{
    "alg" : "AES256",
    "typ" : "JWT"
}

payload:

{
"sub": "1234567890",
"name": "John Doe",
"admin": true
}

signature:



