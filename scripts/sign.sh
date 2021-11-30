#!/usr/bin/env sh

sign() {
    openssl genrsa -out $1.key 2048
    openssl req -new -out $1.csr -key $1.key \
        -reqexts SAN \
        -subj "/C=IE/ST=Dublin/L=Dublin/O=TCD/OU=SCSS/CN=rasp-*.scss.tcd.ie/CN=127.*.*.*"
    openssl ca -in $1.csr -out $1.crt \
        -extensions SAN
        
    rm $1.csr
}

main() {
    sign $1/bundled $2
}

main $1 $2