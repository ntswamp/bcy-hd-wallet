package main

import (
	"fmt"
	"net/http"

	"github.com/blockcypher/gobcy"
	"github.com/wemeetagain/go-hdwallet"
)

var (
	seller_addr = "CAidSjXaJEgXJhKvdQ6CE9qEMBkZCmtycC"
	seller_pub  = "03381be444b91ca55d2ff7074e2d36650ea040e1bc1c05b9ea48ff39be57c95198"

	buyer_addr = "C5yS3pPG2E5Wfx5uzx5ibWvZxX17uFe1dG"
	buyer_pub  = "02079cff4199262f6bf8e1993b4fc2c65ab525a36c542940ee7e879adb2127bf01"

	faucet_use_addr = "CBzyZEAGmRaxmapYEhErX4kMrN93iaFq5v"
	faucet_private  = "8b92199b665a1f23130f8a40dfc499d82859adf094ef957d17070890627858bb"
	faucet_public   = "03927e6938c23985fa6ade83a6a778e718552632ac298659d8f2a85dd8556a353a"
	faucet_wif      = "Bt1LZERbcwLpLTfDUWm4jnZpn4FrntqEgdVUjTqmidos4AMaB7Hj"
)

func main() {

	//get blockchian client
	bcy := gobcy.API{Token: BLOCKCYPHER_TOKEN, Coin: "bcy", Chain: "test"}
	/**************FAUCET***************
	faucet_addr := gobcy.AddrKeychain{
		Address: "CBzyZEAGmRaxmapYEhErX4kMrN93iaFq5v",
		Private: "8b92199b665a1f23130f8a40dfc499d82859adf094ef957d17070890627858bb",
		Public:  "03927e6938c23985fa6ade83a6a778e718552632ac298659d8f2a85dd8556a353a",
		Wif:     "Bt1LZERbcwLpLTfDUWm4jnZpn4FrntqEgdVUjTqmidos4AMaB7Hj",
	}

	_, err := bcy.Faucet(faucet_addr, 1)
	if err != nil {
		fmt.Println(err)
	}

	***************/

	///****check balance****
	addr, err := bcy.GetAddrBal(buyer_addr, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", addr)

	/*
		//Post New TXSkeleton
		skel, err := bcy.NewTX(gobcy.TempNewTX(faucet_use_addr, buyer_addr, *big.NewInt(190)), false)
		if err != nil {
			fmt.Println(err)
		}
		//Sign it locally
		err = skel.Sign([]string{faucet_private})
		if err != nil {
			fmt.Println(err)
		}
		//Send TXSkeleton
		skel, err = bcy.SendTX(skel)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", skel)
	*/

	/*
			wallet, _, xpri, err := create_hd_wallet(bcy, COMPANY_WALLET)
			if err != nil {
				println(err.Error())
			}

			fmt.Printf("wallet created:%+v\n", wallet)
			fmt.Printf("xpri:%+v\n", xpri)


		//list all wallets by token
		walletNames, _ := bcy.ListWallets()
		fmt.Printf("Wallets:%v\n", walletNames)

			//hdwallet
			wallet, _ := bcy.GetAddrHDWallet(COMPANY_WALLET, nil)
			fmt.Printf("Wallet: %+v\n", wallet)

			derived_wallet, _ := bcy.DeriveAddrHDWallet(COMPANY_WALLET, map[string]string{"count": "1"})
			fmt.Printf("Wallet: %+v\n", derived_wallet)

		//hdwallet
		wallet, _ := bcy.GetAddrHDWallet(COMPANY_WALLET, nil)
		fmt.Printf("Wallet: %+v\n", wallet)


			//delete wallet
			err := bcy.DeleteHDWallet(COMPANY_WALLET)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Wallet Deleted")
			}
	*/

}

// return xpub
func create_hd_wallet(client gobcy.API, wallet_name string) (wallet gobcy.HDWallet, xpub string, xpri string, err error) {
	// Generate a random 256 bit seed
	seed, _ := hdwallet.GenSeed(256)
	// Create a master private key
	masterprv := hdwallet.MasterKey(seed)
	// Convert master private key to public key
	masterpub := masterprv.Pub()

	//stringify xpub key
	xpub = masterpub.String()
	xpri = masterprv.String()

	wallet, err = client.CreateHDWallet(gobcy.HDWallet{Name: wallet_name, ExtPubKey: xpub})
	if err != nil {
		return
	}
	fmt.Printf("HD Wallet created: %+v\nxPub key:%+v\n", wallet, xpub)
	return
}

func derive_payment_address_from_wallet(client gobcy.API, wallet_name string) (string, error) {
	derived_wallet, err := client.DeriveAddrHDWallet(wallet_name, map[string]string{"count": "1"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parital HD Wallet: %+v\n", derived_wallet)
	return "", nil

}

func handlePurchase(bcy gobcy.API, writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	id := query.Get("id")

	fmt.Printf("GET: id=%s\n", id)

	wa, addr, err := bcy.GenAddrWallet(COMPANY_WALLET)
	if err != nil {
		fmt.Println(err)
	}
	if wa.Addresses[len(wa.Addresses)-1] == addr.Address {
		fmt.Fprintf(writer, `{"code":0,"addr":%s}`, addr.Address)
	}

}

func CollectBalance(bcy gobcy.API) {
	waAddr, err := bcy.GetAddrBal(COMPANY_WALLET, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", waAddr.Balance)
	//input:= {inputs:[{"wallet_name":COMPANY_WALLET, "wallet_token":BLOCKCYPHER_TOKEN}], value: waAddr.Balance.Int64()}

	//bcy.NewTX()
}
