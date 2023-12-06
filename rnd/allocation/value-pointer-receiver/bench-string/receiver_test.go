package receiver_test

import (
	"testing"
)

type someSmall struct {
	par0 string
	par1 string
	par2 string
	par3 string
	par4 string
	par5 string
	par6 string
	par7 string
	par8 string
	par9 string
}

func (s someSmall) valReceivSum() string {
	out := s.par0 + s.par9
	return out
}

func (s *someSmall) pointReceivSum() string {
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
	par0  string
	par1  string
	par2  string
	par3  string
	par4  string
	par5  string
	par6  string
	par7  string
	par8  string
	par9  string
	par10 string
	par11 string
	par12 string
	par13 string
	par14 string
	par15 string
	par16 string
	par17 string
	par18 string
	par19 string
	par20 string
	par21 string
	par22 string
	par23 string
	par24 string
	par25 string
	par26 string
	par27 string
	par28 string
	par29 string
	par30 string
	par31 string
	par32 string
	par33 string
	par34 string
	par35 string
	par36 string
	par37 string
	par38 string
	par39 string
}

func (s someLarge) valReceivSum() string {
	out := s.par0 + s.par39
	return out
}

func (s *someLarge) pointReceivSum() string {
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
