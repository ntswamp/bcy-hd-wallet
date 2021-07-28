package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/blockcypher/gobcy"
	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/wemeetagain/go-hdwallet"
)

func main() {
	const (
		company_addr  = "C4zQmUHCZqZpz6meiiiyxLJ1R2eRLDXEMi"
		company_addr2 = "C26eLLR6e26pik8XK8cR736ARbDXJLVdfC"

		faucet_use_addr = "CBzyZEAGmRaxmapYEhErX4kMrN93iaFq5v"
		faucet_private  = "8b92199b665a1f23130f8a40dfc499d82859adf094ef957d17070890627858bb"
		faucet_public   = "03927e6938c23985fa6ade83a6a778e718552632ac298659d8f2a85dd8556a353a"
		faucet_wif      = "Bt1LZERbcwLpLTfDUWm4jnZpn4FrntqEgdVUjTqmidos4AMaB7Hj"
	)

	//get blockchian client
	bcy := gobcy.API{Token: BLOCKCYPHER_TOKEN, Coin: "btc", Chain: "test3"}

	w, pub, pri, _ := create_hd_wallet(bcy, COMPANY_WALLET)
	println(pub, pri, w.ExtPubKey)

	derived_wallet, err := bcy.DeriveAddrHDWallet(COMPANY_WALLET, map[string]string{"count": "1"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parital HD Wallet: %+v\n", derived_wallet)

	/*****delete wallet**********
	err := bcy.DeleteHDWallet(COMPANY_WALLET)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Wallet Deleted")
	}

	/*************FAUCET************
	faucet_addr := gobcy.AddrKeychain{
		Address: "CBzyZEAGmRaxmapYEhErX4kMrN93iaFq5v",
		Private: "8b92199b665a1f23130f8a40dfc499d82859adf094ef957d17070890627858bb",
		Public:  "03927e6938c23985fa6ade83a6a778e718552632ac298659d8f2a85dd8556a353a",
		Wif:     "Bt1LZERbcwLpLTfDUWm4jnZpn4FrntqEgdVUjTqmidos4AMaB7Hj",
	_, err := bcy.Faucet(faucet_addr, 10e6)
	if err != nil {
	fmt.Println(err)

	skel, err := bcy.NewTX(gobcy.TempNewTX(faucet_use_addr, company_addr2, *big.NewInt(2e5)), false)
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


	///****CHECK BALANCE***
	addr, err := bcy.GetAddrBal(COMPANY_WALLET, map[string]string{"omitWalletAddresses": "true"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", addr.Balance.String())



	//*****CREATE HD WALLET************************************************
	w, pub, pri, _ := create_hd_wallet(bcy, COMPANY_WALLET)
	println(pub, pri, w.ExtPubKey)


	//*****DERIVE ADDRESS AND CALLBACK********************************
	derive_payment_address_from_wallet_and_register_callback(bcy, COMPANY_WALLET)

	//*****DERIVE ADDRESS********************************
	derived_wallet, err := bcy.DeriveAddrHDWallet(COMPANY_WALLET, map[string]string{"count": "1"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parital HD Wallet: %+v\n", derived_wallet)


	//************CREATE &&& DERIVE *****
	w, pub, pri, _ := create_hd_wallet(bcy, COMPANY_WALLET)
	println(pub, pri, w.ExtPubKey)

	derived_wallet, err := bcy.DeriveAddrHDWallet(COMPANY_WALLET, map[string]string{"count": "1"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parital HD Wallet: %+v\n", derived_wallet)

	//****** TX ************************



	//***********list all wallets by token************
	walletNames, _ := bcy.ListWallets()
	fmt.Printf("Wallets:%v\n", walletNames)


	//******GET hdwallet**************************
	wallet, _ := bcy.GetAddrHDWallet(COMPANY_WALLET, nil)
	fmt.Printf("Wallet: %+v\n", wallet)

	*/

}

// return xpub
func create_hd_wallet(client gobcy.API, wallet_name string) (wallet gobcy.HDWallet, xpub string, xpri string, err error) {

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
	pubkey, err := acct0.Neuter()
	if err != nil {
		fmt.Println(err)
		return
	}
	xpub = pubkey.String()

	wallet, err = client.CreateHDWallet(gobcy.HDWallet{Name: wallet_name, ExtPubKey: xpub})
	if err != nil {
		return
	}
	fmt.Printf("HD Wallet created: %+v\n", wallet)
	return
}

func derive_payment_address_from_wallet_and_register_callback(client gobcy.API, wallet_name string) string {
	derived_wallet, err := client.DeriveAddrHDWallet(wallet_name, map[string]string{"count": "1"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parital HD Wallet: %+v\n", derived_wallet)
	derived_addr := derived_wallet.Chains[0].ChainAddr[0].Address
	derived_addr_pub := derived_wallet.Chains[0].ChainAddr[0].Public
	_ = derived_addr_pub

	//register callback
	hook, err := client.CreateHook(gobcy.Hook{Event: "tx-confirmation", Address: derived_addr, Confirmations: 1, URL: "https://my.domain.com/callbacks/payments"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", hook)

	return derived_addr

}

/*
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
*/

//curl version of transfer from wallet
func collect_all_balance_from_wallet(bcy gobcy.API, wallet_name string, token string, to_address string) {
	addr, err := bcy.GetAddrBal(COMPANY_WALLET, map[string]string{"omitWalletAddresses": "true"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", addr.Balance.String())

	param := `{"inputs":[{"wallet_name":"` + wallet_name + `", "wallet_token":"` + token + `"}], "outputs":[{"addresses": ["` + to_address + `"], "value": ` + addr.Balance.String() + `}]}`
	println(param)
	body := strings.NewReader(param)
	req, err := http.NewRequest("POST", "https://api.blockcypher.com/v1/bcy/test/txs/new", body)
	if err != nil {
		// handle err
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		fmt.Println(err)
	}
	var data map[string]interface{}
	rbody, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(rbody), &data)
	fmt.Println(data)
	defer resp.Body.Close()
}

func transfer_from_wallet(client gobcy.API, wallet_name string, to_address string, amount int64) {
	//Post New TXSkeleton
	input := gobcy.TXInput{WalletName: wallet_name}
	output := gobcy.TXOutput{Addresses: []string{to_address}, Value: *big.NewInt(amount)}
	tempTx := gobcy.TX{Inputs: []gobcy.TXInput{input}, Outputs: []gobcy.TXOutput{output}}

	skel, err := client.NewTX(tempTx, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Sign it locally
	prvkey, _ := hdwallet.StringChild(COMPANY_WALLET_MASTER_PRIVATE_KEY, 0)
	base58_decoded := base58.Decode(prvkey)
	hex_private_key := hex.EncodeToString(base58_decoded)
	err = skel.Sign([]string{hex_private_key})
	if err != nil {
		fmt.Println(err)
	}
	//Send TXSkeleton
	skel, err = client.SendTX(skel)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", skel)
}
