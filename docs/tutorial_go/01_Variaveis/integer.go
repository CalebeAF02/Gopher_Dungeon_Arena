package goguide

// Tipos inteiros em Go incluem signed e unsigned com tamanhos fixos
// e também aliases como byte e rune. O tipo padrão int varia conforme
// a arquitetura (32 ou 64 bits).

func SumIntegers(a int, b int) int {
	return a + b
}

func IntegerTypes() (int8, int16, int32, int64, uint8, uint16, uint32, uint64, byte, rune, uintptr) {
	var a int8 = -120
	var b int16 = 32000
	var c int32 = 2000000000
	var d int64 = -9000000000000000000
	var e uint8 = 250
	var f uint16 = 65000
	var g uint32 = 4000000000
	var h uint64 = 18000000000000000000
	var i byte = 255          // alias de uint8
	var j rune = '✋'          // alias de int32, representa um ponto de código Unicode
	var k uintptr = 123456789 // tipo para guardar valores de ponteiro como inteiro
	return a, b, c, d, e, f, g, h, i, j, k
}

// ConvertUnsigned demonstra conversão explícita entre inteiros uSigned e signed.
func ConvertUnsigned(value uint16) int32 {
	return int32(value)
}
