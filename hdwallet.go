package main

import (
	"fmt"

	"github.com/btcsuite/btcutil/hdkeychain"
)

/**
If you want to use BIP32 default wallet layout
you can submit the extended public key of m/0' (which can only be derived from your master private key) with subchain indexes = [0, 1].
Subchain index 0 represents the external chain (of account 0) and will discover all k keypairs that look like: m/0'/0/k.
Subchain index 1 represents the internal chain (of account 0) and will discover all k keypairs in m/0'/1/k.

*/

func hd() {
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

	fmt.Println(acct0.IsPrivate())

}
