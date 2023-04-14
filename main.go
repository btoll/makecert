package main

func main() {
	generateCert(TLSCert{
		EcdsaCurve: "P384",
		Host:       "127.0.0.1,192.168.1.96",
		IsCA:       true,
		RsaBits:    3072,
	})
}
