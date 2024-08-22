// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package box authenticates and encrypts small messages using public-key cryptography.

Box uses Curve25519, XSalsa20 and Poly1305 to encrypt and authenticate
messages. The length of messages is not hidden.

It is the caller's responsibility to ensure the uniqueness of nonces—for
example, by using nonce 1 for the first message, nonce 2 for the second
message, etc. Nonces are long enough that randomly generated nonces have
negligible risk of collision.

Messages should be small because:

1. The whole message needs to be held in memory to be processed.

2. Using large messages pressures implementations on small machines to decrypt
and process plaintext before authenticating it. This is very dangerous, and
this API does not allow it, but a protocol that uses excessive message sizes
might present some implementations with no other choice.

3. Fixed overheads will be sufficiently amortised by messages as small as 8KB.

4. Performance may be improved by working with messages that fit into data caches.

Thus large amounts of data should be chunked so that each message is small.
(Each message still needs a unique nonce.) If in doubt, 16KB is a reasonable
chunk size.

This package is interoperable with NaCl: https://nacl.cr.yp.to/box.html.
Anonymous sealing/opening is an extension of NaCl defined by and interoperable
with libsodium:
https://libsodium.gitbook.io/doc/public-key_cryptography/sealed_boxes.
*/
package box

import (
	cryptorand "crypto/rand"
	"io"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/salsa20/salsa"
)

const (
	// Overhead is the number of bytes of overhead when boxing a message.
	Overhead = secretbox.Overhead

	// AnonymousOverhead is the number of bytes of overhead when using anonymous
	// sealed boxes.
	AnonymousOverhead = Overhead + 32
)

// GenerateKey generates a new public/private key pair suitable for use with
// Seal and Open.
func GenerateKey(rand io.Reader) (publicKey, privateKey *[32]byte, err error) {
	publicKey = new([32]byte)
	privateKey = new([32]byte)
	_, err = io.ReadFull(rand, privateKey[:])
	if err != nil {
		publicKey = nil
		privateKey = nil
		return
	}

	curve25519.ScalarBaseMult(publicKey, privateKey)
	return
}

var zeros [16]byte

// Precompute calculates the shared key between peersPublicKey and privateKey
// and writes it to sharedKey. The shared key can be used with
// OpenAfterPrecomputation and SealAfterPrecomputation to speed up processing
// when using the same pair of keys repeatedly.
func Precompute(sharedKey, peersPublicKey, privateKey *[32]byte) {
	curve25519.ScalarMult(sharedKey, privateKey, peersPublicKey)
	salsa.HSalsa20(sharedKey, &zeros, sharedKey, &salsa.Sigma)
}

// Seal appends an encrypted and authenticated copy of message to out, which
// will be Overhead bytes longer than the original and must not overlap it. The
// nonce must be unique for each distinct message for a given pair of keys.
func Seal(out, message []byte, nonce *[24]byte, peersPublicKey, privateKey *[32]byte) []byte {
	var sharedKey [32]byte
	Precompute(&sharedKey, peersPublicKey, privateKey)
	return secretbox.Seal(out, message, nonce, &sharedKey)
}

// SealAfterPrecomputation performs the same actions as Seal, but takes a
// shared key as generated by Precompute.
func SealAfterPrecomputation(out, message []byte, nonce *[24]byte, sharedKey *[32]byte) []byte {
	return secretbox.Seal(out, message, nonce, sharedKey)
}

// Open authenticates and decrypts a box produced by Seal and appends the
// message to out, which must not overlap box. The output will be Overhead
// bytes smaller than box.
func Open(out, box []byte, nonce *[24]byte, peersPublicKey, privateKey *[32]byte) ([]byte, bool) {
	var sharedKey [32]byte
	Precompute(&sharedKey, peersPublicKey, privateKey)
	return secretbox.Open(out, box, nonce, &sharedKey)
}

// OpenAfterPrecomputation performs the same actions as Open, but takes a
// shared key as generated by Precompute.
func OpenAfterPrecomputation(out, box []byte, nonce *[24]byte, sharedKey *[32]byte) ([]byte, bool) {
	return secretbox.Open(out, box, nonce, sharedKey)
}

// SealAnonymous appends an encrypted and authenticated copy of message to out,
// which will be AnonymousOverhead bytes longer than the original and must not
// overlap it. This differs from Seal in that the sender is not required to
// provide a private key.
func SealAnonymous(out, message []byte, recipient *[32]byte, rand io.Reader) ([]byte, error) {
	if rand == nil {
		rand = cryptorand.Reader
	}
	ephemeralPub, ephemeralPriv, err := GenerateKey(rand)
	if err != nil {
		return nil, err
	}

	var nonce [24]byte
	if err := sealNonce(ephemeralPub, recipient, &nonce); err != nil {
		return nil, err
	}

	if total := len(out) + AnonymousOverhead + len(message); cap(out) < total {
		original := out
		out = make([]byte, 0, total)
		out = append(out, original...)
	}
	out = append(out, ephemeralPub[:]...)

	return Seal(out, message, &nonce, recipient, ephemeralPriv), nil
}

// OpenAnonymous authenticates and decrypts a box produced by SealAnonymous and
// appends the message to out, which must not overlap box. The output will be
// AnonymousOverhead bytes smaller than box.
func OpenAnonymous(out, box []byte, publicKey, privateKey *[32]byte) (message []byte, ok bool) {
	if len(box) < AnonymousOverhead {
		return nil, false
	}

	var ephemeralPub [32]byte
	copy(ephemeralPub[:], box[:32])

	var nonce [24]byte
	if err := sealNonce(&ephemeralPub, publicKey, &nonce); err != nil {
		return nil, false
	}

	return Open(out, box[32:], &nonce, &ephemeralPub, privateKey)
}

// sealNonce generates a 24 byte nonce that is a blake2b digest of the
// ephemeral public key and the receiver's public key.
func sealNonce(ephemeralPub, peersPublicKey *[32]byte, nonce *[24]byte) error {
	h, err := blake2b.New(24, nil)
	if err != nil {
		return err
	}

	if _, err = h.Write(ephemeralPub[:]); err != nil {
		return err
	}

	if _, err = h.Write(peersPublicKey[:]); err != nil {
		return err
	}

	h.Sum(nonce[:0])

	return nil
}
