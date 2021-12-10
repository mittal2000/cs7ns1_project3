#!/usr/bin/env sh
# Misbah Rizaee
main() {
    openssl genrsa -out $1.key 1024
    openssl req -new -x509 -days 3650 \
        -key $1.key -out $1.crt \
        -extensions v3_ca \
        -subj "/C=IE/ST=Dublin/L=Dublin/O=TCD/OU=SCSS/CN=root/"
}

main $1
