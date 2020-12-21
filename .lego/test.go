package _lego

import (
"crypto"
"crypto/ecdsa"
"crypto/elliptic"
"crypto/rand"
"fmt"
"github.com/go-acme/lego/v4/certcrypto"
"github.com/go-acme/lego/v4/certificate"
"github.com/go-acme/lego/v4/lego"
"github.com/go-acme/lego/v4/providers/dns"
"github.com/go-acme/lego/v4/registration"
"log"
)

type CertUser struct {
	Email string
	Registration *registration.Resource
	key crypto.PrivateKey
}

func (u *CertUser) GetEmail() string {
	return u.Email
}

func (u CertUser) GetRegistration() *registration.Resource {
	return u.Registration
}

func (u *CertUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func main() {

	// Create user. New accounts need an email and private key to start
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	certUser := CertUser{
		Email:        "ben@grewelltech.com",
		key:          privateKey,
	}

	config := lego.NewConfig(&certUser)

	config.Certificate.KeyType = certcrypto.RSA2048

	// A client facilitages communication with the CA server.
	client, err := lego.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// Get DNS provider
	provider, err := dns.NewDNSChallengeProviderByName("cloudflare")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Challenge.SetDNS01Provider(provider)
	if err != nil {
		log.Fatal(err)
	}

	// New users need to register
	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		log.Fatal(err)
	}
	certUser.Registration = reg

	// Request
	request := certificate.ObtainRequest{
		Domains: []string{"a.hosts.attackmap.grewelltech.com"},
		Bundle: true,
	}
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", certificates)
}
