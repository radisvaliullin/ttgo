package receiver_test

import (
	"testing"
)

type someSmall struct {
	par0 int
	// par1 int
	// par2 int
	// par3 int
	// par4 int
	// par5 int
	// par6 int
	// par7 int
	// par8 int
	par9 int
}

func (s someSmall) valReceivSum() int {
	out := s.par0 + s.par9
	return out
}

func (s *someSmall) pointReceivSum() int {
	out := s.par0 + s.par9
	return out
}

func BenchmarkSomeSmallValueReceiverTest(b *testing.B) {
	s := someSmall{}
	for i := 0; i < b.N; i++ {
		s.valReceivSum()
	}
}

func BenchmarkSomeSmallPointReceiverTest(b *testing.B) {
	s := someSmall{}
	for i := 0; i < b.N; i++ {
		s.pointReceivSum()
	}
}

type someLarge struct {
	par0  int
	par1  int
	par2  int
	par3  int
	par4  int
	par5  int
	par6  int
	par7  int
	par8  int
	par9  int
	par10 int
	par11 int
	par12 int
	par13 int
	par14 int
	par15 int
	par16 int
	par17 int
	par18 int
	par19 int
	par20 int
	par21 int
	par22 int
	par23 int
	par24 int
	par25 int
	par26 int
	par27 int
	par28 int
	par29 int
	par30 int
	par31 int
	par32 int
	par33 int
	par34 int
	par35 int
	par36 int
	par37 int
	par38 int
	par39 int
}

func (s someLarge) valReceivSum() int {
	out := s.par0 + s.par39
	return out
}

func (s *someLarge) pointReceivSum() int {
	out := s.par0 + s.par39
	return out
}

func BenchmarkSomeLargeValueReceiverTest(b *testing.B) {
	s := someLarge{}
	for i := 0; i < b.N; i++ {
		s.valReceivSum()
	}
}

func BenchmarkSomeLargePointReceiverTest(b *testing.B) {
	s := someLarge{}
	for i := 0; i < b.N; i++ {
		s.pointReceivSum()
	}
}
