package common

import "testing"

func TestGraphNode_Ctor(t *testing.T) {
	// arrange
	want := NewLocation(0, 0, 0)
	// act
	node := newNode(want)
	have := node.getLocation()
	// assert
	if !want.Compare(have) {
		t.Errorf("")
	}
}

func TestGraphNode_Ctor1(t *testing.T) {
	// arrange
	want := NewLocation(^uint(0), ^uint(0), ^uint(0))
	// act
	node := newNode(want)
	have := node.getLocation()
	// assert
	if !want.Compare(have) {
		t.Errorf("")
	}
}

func TestGraphNode_GetLocation(t *testing.T) {
	// arrange
	node := newNode(NewLocation(0, 0, 0))
	want := NewLocation(0, 0, 0)
	// act
	have := node.getLocation()
	// assert
	if !want.Compare(have) {
		t.Errorf("")
	}
}

func TestGraphNode_GetLocation1(t *testing.T) {
	// arrange
	node := newNode(NewLocation(^uint(0), ^uint(0), ^uint(0)))
	want := NewLocation(^uint(0), ^uint(0), ^uint(0))
	// act
	have := node.getLocation()
	// assert
	if !want.Compare(have) {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(0, 0, 0))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor0(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(64, 64, 64))
	node2 := newNode(NewLocation(64, 64, 64))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor1(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(^uint(0), ^uint(0), ^uint(0)))
	node2 := newNode(NewLocation(^uint(0), ^uint(0), ^uint(0)))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor2(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(0, 0, 2))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor3(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(0, 2, 0))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor4(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(2, 0, 0))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor5(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(2, 2, 0))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor6(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(2, 0, 2))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor7(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(0, 2, 2))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor8(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(1, 1, 2))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor9(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(1, 2, 1))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor10(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(2, 1, 1))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor11(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(1, 1, 1))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor12(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(3, 3, 2))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor13(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(3, 2, 3))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor14(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(2, 3, 3))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor15(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(3, 3, 3))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor16(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(2, 2, 3))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if !have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor17(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(2, 3, 2))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if !have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor18(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(3, 2, 2))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if !have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor19(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(2, 2, 1))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if !have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor20(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(2, 1, 2))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if !have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor21(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(1, 2, 2))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if !have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor22(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(^uint(0), ^uint(0), ^uint(0)))
	// act
	have1 := node1.isNeighbor(node2)
	have2 := node2.isNeighbor(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_IsNeighbor23(t *testing.T) {
	// arrange
	node := newNode(NewLocation(0, 0, 0))
	// act
	have := node.isNeighbor(nil)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphNode_Connect(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(0, 0, 1))
	// act
	have, _, _ := node1.connect(node2)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphNode_Connect1(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(0, 0, 1))
	wasSuccessful, node1, node2 := node1.connect(node2)

	if !wasSuccessful {
		t.Errorf("")
	}
	// act
	have, node2, node1 := node2.connect(node1)
	// assert
	if have {
		t.Errorf("%v reverse connenct %v", node1, node2)
	}
}

func TestGraphNode_Connect2(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(0, 0, 1))
	wasSuccessful, node1, node2 := node1.connect(node2)

	if !wasSuccessful {
		t.Errorf("")
	}
	// act
	have, node1, node2 := node1.connect(node2)
	// assert
	if have {
		t.Errorf("%v double connenct %v", node1, node2)
	}
}

func TestGraphNode_Connect3(t *testing.T) {
	// arrange
	node := newNode(NewLocation(0, 0, 0))
	// act
	have, _, _ := node.connect(node)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphNode_Connect5(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	// act
	have, _, result := node1.connect(nil)
	// assert
	if have || result != nil {
		t.Errorf("")
	}
}

func TestGraphNode_Connect6(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(1, 0, 1))
	// act
	have1, node1, node2 := node1.connect(node2)
	have2, _, _ := node2.connect(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_Connect7(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(1, 1, 0))
	// act
	have1, node1, node2 := node1.connect(node2)
	have2, _, _ := node2.connect(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_Connect8(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(1, 1, 1))
	// act
	have1, node1, node2 := node1.connect(node2)
	have2, _, _ := node2.connect(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_Connect9(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(2, 2, 0))
	// act
	have1, node1, node2 := node1.connect(node2)
	have2, _, _ := node2.connect(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_Connect10(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(2, 0, 2))
	// act
	have1, node1, node2 := node1.connect(node2)
	have2, _, _ := node2.connect(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_Connect11(t *testing.T) {
	// arrange
	node1 := newNode(NewLocation(2, 2, 2))
	node2 := newNode(NewLocation(0, 2, 2))
	// act
	have1, node1, node2 := node1.connect(node2)
	have2, _, _ := node2.connect(node1)
	// assert
	if have1 || have1 != have2 {
		t.Errorf("")
	}
}

func TestGraphNode_Connect12(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(2, 2, 2))
	node1 := newNode(NewLocation(1, 2, 2))
	node2 := newNode(NewLocation(2, 1, 2))
	node3 := newNode(NewLocation(2, 2, 1))
	node4 := newNode(NewLocation(3, 2, 2))
	node5 := newNode(NewLocation(2, 3, 2))
	node6 := newNode(NewLocation(2, 2, 3))
	// act
	have1, sut, node1 := sut.connect(node1)
	have2, sut, node2 := sut.connect(node2)
	have3, sut, node3 := sut.connect(node3)
	have4, sut, node4 := sut.connect(node4)
	have5, sut, node5 := sut.connect(node5)
	have6, sut, node6 := sut.connect(node6)
	wantList := []Node{
		node1, node2, node3, node4, node5, node6,
	}
	haveList := sut.getConnected()
	// assert
	if !have1 || have1 != have2 ||
		have2 != have3 || have3 != have4 ||
		have4 != have5 || have5 != have6 {
		t.Errorf("")
	}

	for i, elem := range wantList {
		if !elem.hardCompare(haveList[i]) {
			t.Errorf("")
		}
	}
}

func TestGraphNode_Disconnect(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(2, 2, 2))
	// act
	have, _, result := sut.disconnect(nil)
	// assert
	if have || result != nil {
		t.Errorf("")
	}
}

func TestGraphNode_Disconnect1(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(2, 2, 2))
	// act
	have, sut, result := sut.disconnect(sut)
	// assert
	if have || !sut.hardCompare(result) {
		t.Errorf("")
	}
}

func TestGraphNode_Disconnect2(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(2, 2, 2))
	node1 := newNode(NewLocation(2, 2, 2))
	// act
	have, sut, node1 := sut.disconnect(node1)
	// assert
	if have || !sut.hardCompare(node1) {
		t.Errorf("")
	}
}

func TestGraphNode_Disconnect3(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(2, 2, 2))
	node1 := newNode(NewLocation(2, 2, 1))
	// act
	have, _, _ := sut.disconnect(node1)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphNode_Disconnect4(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(2, 2, 2))
	node1 := newNode(NewLocation(2, 2, 1))
	res, sut, node1 := sut.connect(node1)

	if !res {
		t.Errorf("")
	}
	// act
	have, _, _ := sut.disconnect(node1)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphNode_Disconnect5(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(2, 2, 2))
	node1 := newNode(NewLocation(2, 2, 1))
	res, sut, node1 := sut.connect(node1)

	if !res {
		t.Errorf("")
	}
	// act
	have1, sut, _ := sut.disconnect(node1)
	have2, _, _ := sut.disconnect(node1)
	// assert
	if !have1 || have2 {
		t.Errorf("")
	}
}

func TestGraphNode_Disconnect6(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(2, 2, 2))
	node1 := newNode(NewLocation(2, 2, 1))
	res, sut, node1 := sut.connect(node1)

	if !res {
		t.Errorf("")
	}
	// act
	have1, sut, node1 := sut.disconnect(node1)
	have2, _, _ := node1.disconnect(sut)
	// assert
	if !have1 || have2 {
		t.Errorf("")
	}
}

func TestGraphNode_Disconnect7(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(2, 2, 2))
	node1 := newNode(NewLocation(2, 2, 1))
	node2 := newNode(NewLocation(2, 1, 2))
	node3 := newNode(NewLocation(1, 2, 2))
	node4 := newNode(NewLocation(2, 2, 3))
	node5 := newNode(NewLocation(2, 3, 2))
	node6 := newNode(NewLocation(3, 2, 2))
	_, sut, node1 = sut.connect(node1)
	_, sut, node2 = sut.connect(node2)
	_, sut, node3 = sut.connect(node3)
	_, sut, node4 = sut.connect(node4)
	_, sut, node5 = sut.connect(node5)
	_, sut, node6 = sut.connect(node6)
	// act
	have2, sut, node2 := sut.disconnect(node2)
	have22, _, sut := node2.disconnect(sut)
	wantList2 := []Node{node1, node3, node4, node5, node6}
	haveList2 := sut.getConnected()

	have1, sut, node1 := sut.disconnect(node1)
	have11, _, sut := node1.disconnect(sut)
	wantList1 := []Node{node3, node4, node5, node6}
	haveList1 := sut.getConnected()

	have5, sut, node5 := sut.disconnect(node5)
	have55, _, sut := node5.disconnect(sut)
	wantList5 := []Node{node3, node4, node6}
	haveList5 := sut.getConnected()

	have3, sut, node3 := sut.disconnect(node3)
	have33, _, sut := node3.disconnect(sut)
	wantList3 := []Node{node4, node6}
	haveList3 := sut.getConnected()

	have4, sut, node4 := sut.disconnect(node4)
	have44, _, sut := node4.disconnect(sut)
	wantList4 := []Node{node6}
	haveList4 := sut.getConnected()

	have6, sut, node6 := sut.disconnect(node6)
	have66, _, sut := node6.disconnect(sut)
	wantList6 := make([]Node, 0)
	haveList6 := sut.getConnected()
	// assert
	if !have1 || have11 || !have2 || have22 ||
		!have3 || have33 || !have4 || have44 ||
		!have5 || have55 || !have6 || have66 {
		t.Errorf("")
	}

	for i, elem := range wantList1 {
		if !elem.hardCompare(haveList1[i]) {
			t.Errorf("")
		}
	}

	for i, elem := range wantList2 {
		if !elem.hardCompare(haveList2[i]) {
			t.Errorf("")
		}
	}

	for i, elem := range wantList3 {
		if !elem.hardCompare(haveList3[i]) {
			t.Errorf("")
		}
	}

	for i, elem := range wantList4 {
		if !elem.hardCompare(haveList4[i]) {
			t.Errorf("")
		}
	}

	for i, elem := range wantList5 {
		if !elem.hardCompare(haveList5[i]) {
			t.Errorf("")
		}
	}

	for i, elem := range wantList6 {
		if !elem.hardCompare(haveList6[i]) {
			t.Errorf("")
		}
	}
}

func TestGraphNode_Compare(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	// act
	have := sut.compare(nil)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphNode_Compare1(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	// act
	have := sut.compare(sut)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphNode_Compare2(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	node1 := newNode(NewLocation(0, 0, 0))
	node2 := newNode(NewLocation(2, 2, 2))
	_, node1, _ = node1.connect(node2)
	// act
	have := sut.compare(node1)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphNode_Compare3(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	node1 := newNode(NewLocation(0, 0, 0))
	// act
	have := sut.compare(node1)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphNode_Compare4(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	node1 := newNode(NewLocation(0, 0, 1))
	// act
	have := sut.compare(node1)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphNode_HardCompare(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	// act
	have := sut.hardCompare(nil)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphNode_HardCompare1(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	// act
	have := sut.hardCompare(sut)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphNode_HardCompare2(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	node1 := newNode(NewLocation(0, 0, 0))
	_, sut, _ = sut.connect(node1)
	// act
	have := sut.hardCompare(sut)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphNode_HardCompare3(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	node1 := newNode(NewLocation(0, 0, 0))
	// act
	have := sut.hardCompare(node1)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphNode_HardCompare4(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	node1 := newNode(NewLocation(0, 0, 4))
	// act
	have := sut.hardCompare(node1)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphNode_HardCompare5(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	node1 := newNode(NewLocation(0, 0, 1))
	_, sut, node1 = sut.connect(node1)
	// act
	have := sut.hardCompare(node1)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphNode_HardCompare6(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	sut2 := newNode(NewLocation(0, 0, 0))
	node1 := newNode(NewLocation(0, 0, 1))
	node2 := newNode(NewLocation(0, 1, 0))
	_, sut, _ = sut.connect(node1)
	_, sut2, _ = sut2.connect(node2)
	// act
	have := sut.hardCompare(sut2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphNode_HardCompare7(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	sut2 := newNode(NewLocation(0, 0, 0))
	node1 := newNode(NewLocation(0, 0, 1))
	_, sut, _ = sut.connect(node1)
	// act
	have := sut.hardCompare(sut2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphNode_HardCompare8(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	sut2 := newNode(NewLocation(0, 0, 0))
	node1 := newNode(NewLocation(0, 0, 1))
	_, sut, _ = sut.connect(node1)
	_, sut2, _ = sut2.connect(node1)
	// act
	have := sut.hardCompare(sut2)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphNode_HardCompare9(t *testing.T) {
	// arrange
	sut := newNode(NewLocation(0, 0, 0))
	sut2 := newNode(NewLocation(0, 0, 0))
	node1 := newNode(NewLocation(0, 0, 1))
	node2 := newNode(NewLocation(0, 0, 1))
	_, sut, _ = sut.connect(node1)
	_, sut2, _ = sut2.connect(node2)
	// act
	have := sut.hardCompare(sut2)
	// assert
	if !have {
		t.Errorf("")
	}
}
