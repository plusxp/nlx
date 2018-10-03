package derrsa

import (
	"strings"
	"testing"
)

var publicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyJJQqF7nvhfWHLhnVkCfM0VPq8YEoyjQsDnm2RMJXSJe2pkUWRbFyqdR6bgmv1ZEtduWMhbKVQHcWE8ZfaYSQmVe92obSwOMYcXL0VS+ouTD0vjAWc0oucZ0U1adndMykgugi9PaWf0I7nCEgYRlsf62sXoF+YqwPiWLPPdzsoBDsmtWOYGfSxtJEucXXbbjiBvRsgipHqCwYsuXEc5TDXqE2iyrgDwVrme9vXxAHGRQgaISUTkOZUcVg3zkGx4SreGz5Hh1iL0VPOA0y4SCtl0a/6ccWxnapfZm03ygoxpsMoYhkFKRP32BYYorPMglck3IeW3fp24gZ7TnqeF29QIDAQAB"
var privateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDIklCoXue+F9YcuGdWQJ8zRU+rxgSjKNCwOebZEwldIl7amRRZFsXKp1HpuCa/VkS125YyFspVAdxYTxl9phJCZV73ahtLA4xhxcvRVL6i5MPS+MBZzSi5xnRTVp2d0zKSC6CL09pZ/QjucISBhGWx/raxegX5irA+JYs893OygEOya1Y5gZ9LG0kS5xddtuOIG9GyCKkeoLBiy5cRzlMNeoTaLKuAPBWuZ729fEAcZFCBohJROQ5lRxWDfOQbHhKt4bPkeHWIvRU84DTLhIK2XRr/pxxbGdql9mbTfKCjGmwyhiGQUpE/fYFhiis8yCVyTch5bd+nbiBntOep4Xb1AgMBAAECggEAJTladNUgR1xfJXdn3k8h/L75WVTn2WwixwIIAPOSglF1rzZH23zpeoHxJuJBdrNJD0NQrb0jEgHFZjJP9czTiLUNmfd8V0CyVhcCi7ghpVO1sDrwX/o+e7QENM9Xc0oWEZrM0gageKEIflAtl2f+A7nVKD0QTM7bhL0tCUSXmWH2jJkYZQlylmMTw6uHOPtHUNT772I1rdBDts5lavVNSiR5v0WgCp3OULneLKvt+AnV4yAoVnkd+D/IGR1don03zXAQak7MgcukMiZ7/4EyPy5ycxHw3EgXv2KY5MQ6WSSpslkUhBCCNYAvXy/yWpINeAj4AbdcLLymHm9mDB/YJQKBgQDn6Yq7NvxOGMKp1QCXSDZE55Teo8ECvgvi6rvZbUqBKcKwNt8e7UAu8ofJiYIR7QgYIo5m07doxj0UyiyNFnbGAt0fSkDaSBoWddT3gfhgr2ziPBj0G353DBERR1MLW4jLClZi4tInwYLaITMTuDDlPwWinnLojPJuZxr9S/M0FwKBgQDdZ29pm9VpHCtTI7t16EgvaPoFS9HN0da2S1VXB2bVFFJSNlUengQcelRJA0rXfcQcpxmonVISB2XDKYzQgc3/5OwVmCZqIuS/YxNGwkV7XZN5M0RTAx4dNV2y19iyBcY6/hNFQIR37jIdss9+LSwY8euZ9a6LVqlZqhs96gu40wKBgQDho2SLipApHAYqX0fX6TaEDB3YENN165p9CY4DmkZnLU5zTLAB0ywjgW2ENPoMnubFNLoAafWFR3uegqaeD0E4nwitrP00Y/qd579xoU8cIxJ5e6d7ulRtV5wzh4XRAxoxErnCNJXsgGbj/HFJumg4Y/0Dn86FXL1zlXW5BPB3uQKBgFKQ2EOSEav5DU8b7DS0QvcrNtHyzHLs5h1lSpYB9J9RzMFL8cJRO0dDd1sjFy31gMUOq8rtoEWI7LIzNIgtYB0lycTj0K+UA6o70Fx65l5zLyPrsncimoBZ/m9Lf0wkSyCEz7rrrHrtIVNh828ieWd0EGCB9m2QC+ecQDx+o5BxAoGBAJhHFABQTyApUGVH5mTsaY4p+sqKtEunLchD2Y6l9uf7l5Wz2NeYmBEuldPZnYDxvXXujtTX8hTUPEXfHnGSDqvfNJBTNO51wqSpXBe2fVTixQklmma/3d2fw2AOhch4d33MYUdmzKiVM6taOa8br9k5FYILG3BCf88KYHOOhDnh"
var invalidKey = "invalid key"

func TestDecodeDEREncodedRSAPublicKeyHappy(t *testing.T) {
	_, err := DecodeDEREncodedRSAPublicKey(strings.NewReader(publicKey))
	if err != nil {
		t.Errorf("error parsing private key %s", err)
	}
}

func TestDecodeDEREncodedRSAPublicKeyException(t *testing.T) {
	invalidKey := "invalidKey"
	_, err := DecodeDEREncodedRSAPublicKey(strings.NewReader(invalidKey))
	if err == nil {
		t.Errorf("parsing invalid public key did not generate a error")
	}
}
func TestDecodeDEREncodedRSAPrivateKeyHappy(t *testing.T) {
	_, err := DecodeDEREncodedRSAPrivateKey(strings.NewReader(privateKey))
	if err != nil {
		t.Errorf("error parsing private key %s", err)
	}
}

func TestDecodeDEREncodedRSAPrivateKeyException(t *testing.T) {
	_, err := DecodeDEREncodedRSAPrivateKey(strings.NewReader(invalidKey))
	if err == nil {
		t.Errorf("parsing invalid private key did not generate a error")
	}
}
