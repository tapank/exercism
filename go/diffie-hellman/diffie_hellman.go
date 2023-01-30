package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a key randomly.
func PrivateKey(p *big.Int) *big.Int {
	two := big.NewInt(2)
	max := big.NewInt(0).Sub(p, two)
	k, _ := rand.Int(rand.Reader, max)
	return k.Add(k, two)
}

// PublicKey generates a key based on private key.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

// NewPair generates a private, public key pair.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

// SecretKey generates a secret key.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}
