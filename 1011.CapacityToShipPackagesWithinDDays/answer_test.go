package main

import (
	"testing"
)

func TestAnswerShipWithinDaysBase(t *testing.T) {
	n := AnswerShipWithinDays([]int{1,2,3,4,5,6,7,8,9,10}, 5)
	if n == 15 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func TestAnswerShipWithinDaysSence1(t *testing.T) {

	n := AnswerShipWithinDays([]int{3,2,2,4,1,4}, 3)
	if n == 6 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func TestAnswerShipWithinDaysSence2(t *testing.T) {

	n := AnswerShipWithinDays([]int{147,73,265,305,191,152,192,293,309,292,182,157,381,287,73,162,313,366,346,47}, 10)
	if n == 602 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
