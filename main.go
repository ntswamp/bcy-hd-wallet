package main

import (
	"fmt"
	"net/http"

	"github.com/blockcypher/gobcy"
	"github.com/wemeetagain/go-hdwallet"
)

func main() {
	//get blockchian
	bcy := gobcy.API{Token: BLOCKCYPHER_TOKEN, Coin: "bcy", Chain: "test"}
	wallet, _, err := create_hd_wallet(bcy, "tokenlink-hd-wallet")
	if err != nil {
		println(err.Error())
	}

	fmt.Printf("wallet created:%+v\nxPub:%s\n", wallet.Name, wallet.ExtPubKey)

	//list all wallets by token
	walletNames, _ := bcy.ListWallets()
	fmt.Printf("Wallets:%v\n", walletNames)

	/*
		//delete wallet
		err = bcy.DeleteHDWallet("tokenlink-hd-wallet")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Wallet Deleted")
		}
	*/

}

// return xpub
func create_hd_wallet(client gobcy.API, wallet_name string) (wallet gobcy.HDWallet, xpub string, err error) {
	// Generate a random 256 bit seed
	seed, _ := hdwallet.GenSeed(256)
	// Create a master private key
	masterprv := hdwallet.MasterKey(seed)
	// Convert master private key to public key
	masterpub := masterprv.Pub()
	//stringify xpub key
	xpub = masterpub.String()

	wallet, err = client.CreateHDWallet(gobcy.HDWallet{Name: wallet_name, ExtPubKey: xpub})
	if err != nil {
		return
	}
	fmt.Printf("HD Wallet created: %+v\nxPub key:%+v\n", wallet, xpub)
	return
}

func handlePurchase(bcy gobcy.API, writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	id := query.Get("id")

	fmt.Printf("GET: id=%s\n", id)

	wa, addr, err := bcy.GenAddrWallet(CROSSLINK_WALLET)
	if err != nil {
		fmt.Println(err)
	}
	if wa.Addresses[len(wa.Addresses)-1] == addr.Address {
		fmt.Fprintf(writer, `{"code":0,"addr":%s}`, addr.Address)
	}

}

func CollectBalance(bcy gobcy.API) {
	waAddr, err := bcy.GetAddrBal(CROSSLINK_WALLET, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", waAddr.Balance)
	//input:= {inputs:[{"wallet_name":CROSSLINK_WALLET, "wallet_token":BLOCKCYPHER_TOKEN}], value: waAddr.Balance.Int64()}

	//bcy.NewTX()
}
