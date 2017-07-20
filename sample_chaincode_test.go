package faizal

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status == shim.OK {
		t.Log("Sucess")
		fmt.Println("Success")
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("2", args)
	if res.Status == shim.OK {
		fmt.Println("Invoke succes")
	}

}

func TestInit(t *testing.T) {
	fmt.Println("In Init Testing")
	cc := new(SampleChainCode)
	stub := shim.NewMockStub("InitTest", cc)
	checkInit(t, stub, [][]byte{})
}

func TestInvoke(t *testing.T) {

	fmt.Println("In Invoke Testing")
	cc := new(SampleChainCode)
	stub := shim.NewMockStub("InvokeTesting", cc)
	checkInvoke(t, stub, [][]byte{[]byte("createAsset"), []byte("EIBLC0001"), []byte("FOOD"), []byte("12 kg Apple"), []byte("APPLE101")})
}
