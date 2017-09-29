package hcashutil_test

import (
	"fmt"
	"math"

	"github.com/HcashOrg/hcashutil"
)

func ExampleAmount() {

	a := hcashutil.Amount(0)
	fmt.Println("Zero Atom:", a)

	a = hcashutil.Amount(1e8)
	fmt.Println("100,000,000 Atoms:", a)

	a = hcashutil.Amount(1e5)
	fmt.Println("100,000 Atoms:", a)
	// Output:
	// Zero Atom: 0 HCASH
	// 100,000,000 Atoms: 1 HCASH
	// 100,000 Atoms: 0.001 HCASH
}

func ExampleNewAmount() {
	amountOne, err := hcashutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := hcashutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := hcashutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := hcashutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 HCASH
	// 0.01234567 HCASH
	// 0 HCASH
	// invalid coin amount
}

func ExampleAmount_unitConversions() {
	amount := hcashutil.Amount(44433322211100)

	fmt.Println("Atom to kCoin:", amount.Format(hcashutil.AmountKiloCoin))
	fmt.Println("Atom to Coin:", amount)
	fmt.Println("Atom to MilliCoin:", amount.Format(hcashutil.AmountMilliCoin))
	fmt.Println("Atom to MicroCoin:", amount.Format(hcashutil.AmountMicroCoin))
	fmt.Println("Atom to Atom:", amount.Format(hcashutil.AmountAtom))

	// Output:
	// Atom to kCoin: 444.333222111 kHCASH
	// Atom to Coin: 444333.222111 HCASH
	// Atom to MilliCoin: 444333222.111 mHCASH
	// Atom to MicroCoin: 444333222111 μHCASH
	// Atom to Atom: 44433322211100 Atom
}
