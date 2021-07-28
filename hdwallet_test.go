package main

import (
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

/**
In BIP32 notation, the wallet layout is m/0, m/1, ... and m/i/0, m/i/1, ... for each subchain i if the wallet has subchains.
For example, the path of the fourth address generated is m/3 for a non-subchain wallet.
The path of the fourth address at subchain index two is m/2/3. Note that this is ****FUCKING DIFFERENT**** from the default BIP32 wallet layout.

If you want to use BIP32 default wallet layout
you can submit the extended public key of m/0' (which can only be derived from your master private key) with subchain indexes = [0, 1].
Subchain index 0 represents the external chain (of account 0) and will discover all k keypairs that look like: m/0'/0/k.
Subchain index 1 represents the internal chain (of account 0) and will discover all k keypairs in m/0'/1/k.

*/

func TestHd(t *testing.T) {
	master := "xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jP" +
		"PqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi"

	// Start by getting an extended key instance for the master node.
	// This gives the path:
	//   m
	masterKey, err := hdkeychain.NewKeyFromString(master)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Derive the extended key for account 0.  This gives the path:
	//   m/0H
	acct0, err := masterKey.Derive(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Derive the extended key for the account 0 external chain.  This
	// gives the path:
	//   m/0H/0
	privm00, err := acct0.Derive(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	pubm00, _ := privm00.Neuter()
	t.Log(pubm00, privm00)
	t.Log(pubm00.Address(&chaincfg.TestNet3Params))

	/*
		// Derive the extended key for the account 0 internal chain.  This gives
		// the path:
		//   m/0H/1
		acct0Int, err := acct0.Child(1)
		if err != nil {
			fmt.Println(err)
			return
		}
	*/

}
