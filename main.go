package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/wire"
)

type BitcoinTransaction struct{
	Version int
	Inputs  []TransactionInput
	Outputs []TransactionOutput
	Locktime int
}

type TransactionInput struct{
	PreviousTxHash string
    PreviousTxOutIndex int
    ScriptSig string
    Sequence int

}
type TransactionOutput struct{
	Value int
    ScriptPubKey string

}

func main(){
   transactionhex:=os.Args[1]

	hexString, err:= hex.DecodeString(transactionhex)

	if(err !=nil){
		fmt.Println("Error decoding raw transaction hex:", err)
	}

	var transaction BitcoinTransaction
	var tx wire.MsgTx
	err = tx.Deserialize(bytes.NewReader(hexString))
	if err != nil {
		fmt.Println("Error deserializing raw transaction:", err)
		return
	}
	transaction.Version=int(tx.Version)
	transaction.Locktime=int(tx.LockTime)
	
	

	for _, input := range tx.TxIn {
		// fmt.Printf("Input %d:\n", i)
		// fmt.Printf("  Previous Tx Hash: %s\n", input.PreviousOutPoint.Hash)
		// fmt.Printf("  Previous Tx Index: %d\n", input.PreviousOutPoint.Index)
		// fmt.Printf("  Script Length: %d\n", len(input.SignatureScript))
		// fmt.Println("  Script:", hex.EncodeToString(input.SignatureScript))
		var transactionInput TransactionInput
		transactionInput.PreviousTxHash=hex.EncodeToString(input.PreviousOutPoint.Hash[:])
		transactionInput.PreviousTxOutIndex=int(input.PreviousOutPoint.Index)
		transactionInput.ScriptSig=hex.EncodeToString(input.SignatureScript)

		transaction.Inputs=append(transaction.Inputs, transactionInput)
		
	}
	
	
	for _, output := range tx.TxOut {
		//fmt.Printf("Output %d:\n", i)
		// fmt.Printf("  Value: %d Satoshis\n", output.Value)
		// fmt.Printf("  Script Length: %d\n", len(output.PkScript))
		// fmt.Println("  Script:", hex.EncodeToString(output.PkScript))
		var transactionOutput TransactionOutput
		transactionOutput.ScriptPubKey=hex.EncodeToString(output.PkScript)
		transactionOutput.Value=int(output.Value)
		transaction.Outputs=append(transaction.Outputs, transactionOutput)

	}
	b, err := json.Marshal(transaction)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(b))
	//fmt.Println(transaction)
	// fmt.Println("Transaction Version:", transaction.Version)

	// fmt.Println("Transaction Inputs:",transaction.Inputs)
	// fmt.Println("Transaction Outputs:",transaction.Outputs)
	// fmt.Println("Transaction Locktime:", transaction.Locktime)

}	