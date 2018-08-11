package prototype

import (
	"testing"
)

func TestClone(t *testing.T) {
	//shirtCache := GetShirtsCloner()
	//shirtCache := Shirt{12.00, "empty", White}
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Not implemented yet")
	}

	//a, _ := GetShirtsCloner().GetClone(2)A
	//item1, err := shirtCache.GetClone(White)
	//if err != nil {
	//	t.Error(err)
	//}

	//if item1 == whitePrototype {
		//t.Error("item1 cannot be equal to the white prototype")
	//}
}
