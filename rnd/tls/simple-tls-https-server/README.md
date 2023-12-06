# Simple TLS HTTPS Server

## Gen Self-Signed Certificate Authority (CA)
1. Gen private key
1. Gen CRS
1. Gen CA using PK and CRS
1. All steps in one command
    ```
    openssl req -new -newkey rsa:2048 -keyout ./tmp/ca.key -x509 -sha256 -days 365 -out ./tmp/ca.crt
    ```

## Create a conf file for the server certificate
see config/server.cnf

## Gen server certificate using the self-signed CA
1. private key
    ```
    openssl genrsa -out ./tmp/server.key 2048
    ```
1. server CRS
    ```
    openssl req -new -key ./tmp/server.key -out ./tmp/server.csr -config ./config/server.cnf
    ```
1. view of CRS content and verify x509v3
    ```
    openssl req -noout -text -in ./tmp/server.csr
    ```
1. get server certificate signed with CA
    ```
    openssl x509 -req -in ./tmp/server.csr -CA ./tmp/ca.crt -CAkey ./tmp/ca.key \
        -CAcreateserial -out ./tmp/server.crt -days 365 -sha256 -extfile ./config/server.cnf -extensions v3_ext
    ```
