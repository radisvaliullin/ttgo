package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

const (
	pkCAPath   = "./tmp/clientcakey.pem"
	certCAPath = "./tmp/clientcacert.pem"
	pkPath     = "./tmp/clientkey.pem"
	certPath   = "./tmp/clientcert.pem"
)

func main() {
	fmt.Println("main")

	// Private Key & Self-Signed Cert
	getPrivateKeyAndSelfSingCert()

}

func getPrivateKeyAndSelfSingCert() {
	// generate CA new key
	privateCAKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("failed to generate private ca key: %v", err)
	}
	// generate client new key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("failed to generate private client key: %v", err)
	}

	// gen cert template
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)

	// СA cert template
	caSerialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate CA serial number: %v", err)
	}
	caTemplate := x509.Certificate{
		SerialNumber: caSerialNumber,
		Subject: pkix.Name{
			Organization:  []string{"galley"},
			Country:       []string{"US"},
			Province:      []string{"CA"},
			Locality:      []string{"Petaluma"},
			StreetAddress: []string{"Technology Ln"},
			PostalCode:    []string{"94954"},
			CommonName:    "root_id",
		},
		IsCA:                  true,
		DNSNames:              []string{"localhost"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(48 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}
	// Client Cert Template
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %v", err)
	}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization:  []string{"galley"},
			Country:       []string{"US"},
			Province:      []string{"CA"},
			Locality:      []string{"Petaluma"},
			StreetAddress: []string{"Technology Ln"},
			PostalCode:    []string{"94954"},
			CommonName:    "client_id",
		},
		DNSNames:  []string{"localhost"},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(3 * time.Hour),

		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	// gen CA cert
	caDerBytes, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &privateCAKey.PublicKey, privateCAKey)
	if err != nil {
		log.Fatalf("failed to create CA certificate: %v", err)
	}
	// gen client cert
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &caTemplate, &privateKey.PublicKey, privateCAKey)
	if err != nil {
		log.Fatalf("failed to create client certificate: %v", err)
	}

	// put CA cert to file
	pemCACert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: caDerBytes})
	if pemCACert == nil {
		log.Fatal("failed to encode CA certificate to PEM")
	}
	if err := os.WriteFile(certCAPath, pemCACert, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %v", certCAPath)

	// put client cert to file
	pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if pemCert == nil {
		log.Fatal("failed to encode client certificate to PEM")
	}
	if err := os.WriteFile(certPath, pemCert, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %v", certPath)

	// put CA key to file
	caPrivBytes, err := x509.MarshalPKCS8PrivateKey(privateCAKey)
	if err != nil {
		log.Fatalf("unable to marshal CA private key: %v", err)
	}
	caPemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: caPrivBytes})
	if caPemKey == nil {
		log.Fatal("failed to encode CA key to PEM")
	}
	if err := os.WriteFile(pkCAPath, caPemKey, 0600); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %v", pkCAPath)

	// put key to file
	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatalf("unable to marshal private key: %v", err)
	}
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
	if pemKey == nil {
		log.Fatal("failed to encode key to PEM")
	}
	if err := os.WriteFile(pkPath, pemKey, 0600); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %v", pkPath)
}
