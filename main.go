// importing package for main
package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type block struct{
	transaction 	string
	nonce 			int
	previousHash	string
}

// Function for creating new block
func NewBlock(transaction string, nonce int, previousHash string) *block {
	b := new(block)
	b.transaction = transaction
	b.nonce = nonce
	b.previous_hash = previousHash	
	return b
}

type BlockList struct {
	list []*block
}

func(ls *BlockList) CreateBlock(transaction string, nonce int, previousHash string)*block {
	b := NewBlock(transaction, nonce, previousHash)
	ls.list = append(ls.list, b)
	return b
}

//Function to create hash of block
func CreateHash(b block) string {
	var pkdStr string
	pkd_str = b.transaction + strconv.Itoa(b.nonce) +  b.previous_hash
	return fmt.Sprintf("%x", sha256.Sum256([]byte(pkd_str)))
}

func listBlocks(ledger BlockList){
	for i:=0; i<len(ledger.list); i++{
		fmt.Printf("%s Block number %d %s\n", strings.Repeat("=", 25),i, strings.Repeat("=", 25))
		fmt.Printf("Transaction:  \t\t %v \n", ledger.list[i].transaction )
		fmt.Printf("Nonce:  \t\t %v \n", ledger.list[i].nonce )
		fmt.Printf("Previous Hash:  \t\t %v \n", ledger.list[i].previous_hash)
	}

}

func changeBlock(ledger BlockList, transaction string, nonce int, search int){
	ledger.list[search].transaction = transaction
	ledger.list[search].nonce = nonce
	fmt.Printf("Transaction has been updated!\n")
}
// Function to verify the integrity of chain
func VerifyChain(ledger BlockList) bool{
	for i:=1; i<len(ledger.list); i++{
		var temp string
		temp = CreateHash(*ledger.list[i-1])
		if temp == ledger.list[i].previous_hash{
			continue
		}else{
			fmt.Printf("Varification Failed, Blockchain has been modified!\n")
			return false
		}
	}
	return true
}



func main(){

	in:= bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	ledger:= new(BlockList)
	var counter,option int	
	for{
		fmt.Println("Hello Welcome to Blockchain\n")
	fmt.Printf("\tMenu\n")
	fmt.Printf("1. Add new Block\n2. List Blocks\n3. Change Block\n4. Verify Chain\n5. Exit\n")
	fmt.Printf("Enter your choice: \n")
	var previousHash string
	var nonce int
	
	fmt.Scanln(&option)
	if option==1{
		if len(ledger.list) == 0{
			fmt.Printf("Please enter the value for current transaction: \n")
			currentTransaction, err := in.ReadString('\n')
			_ = err
			nonce = rand.Intn(1000 - 1 + 1) + 1
			previousHash = ""
			ledger.CreateBlock(current_transaction, nonce, previousHash)
			counter++
		}else{
			fmt.Printf("Please enter the value for current transaction: \n")
			current_transaction, err := in.ReadString('\n')
			_ = err
			nonce = rand.Intn(1000 - 1 + 1) + 1
			previousHash = CreateHash(*ledger.list[counter-1])
			ledger.CreateBlock(current_transaction, nonce, previousHash)
			counter++
		}
		
	}else if option==2{
		listBlocks(*ledger)	
	}else if option==3{
		var search int
		var newTransaction string
		fmt.Printf("Please enter Transaction Id to change: \n")
		fmt.Scanln(&search)
		fmt.Printf("Please enter the value for new transaction: \n")
		newTransaction, err := in.ReadString('\n')
		_ = err
		fmt.Printf("Please enter the value for new nonce: \n")
		fmt.Scanln(&nonce)
		if search < counter{
			changeBlock(*ledger,new_transaction,nonce,search)
		}else{
			fmt.Printf("Transaction Id out of bound!\n")
		}

	}else if option==4{
			if VerifyChain(*ledger) == true{
				fmt.Printf("Verification Successful!\n")
			}
	}else if option==5{
			os.Exit(0)
	}else{
		fmt.Printf("\tInvalid Option!\n")
	}
	
	}
	

		
	


}

