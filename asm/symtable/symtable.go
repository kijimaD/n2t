package symtable

type table map[string]int

// 3種類のシンボルがある。組み込みシンボル、ラベルシンボル、変数シンボル

type Symtable struct {
	table    table
	nextAddr int
}

func NewTable() Symtable {
	var t = Symtable{
		// 最初の5つは指定するためのシンボルが2つある。
		// R2とARGはRAM[2]を参照することに用いることができる
		table: table{
			"SP":     0,
			"LCL":    1,
			"ARG":    2,
			"THIS":   3,
			"THAT":   4,
			"R0":     0,
			"R1":     1,
			"R2":     2,
			"R3":     3,
			"R4":     4,
			"R5":     5,
			"R6":     6,
			"R7":     7,
			"R8":     8,
			"R9":     9,
			"R10":    10,
			"R11":    11,
			"R12":    12,
			"R13":    13,
			"R14":    14,
			"R15":    15,
			"SCREEN": 16384,
			"KBD":    24576,
		},
		nextAddr: 16, // 変数シンボルのためのアドレスは16から始まる
	}
	return t
}

func (t *Symtable) AddEntry(symbol string, address int) {
	t.table[symbol] = address
}

func (t *Symtable) AddVariable(symbol string) {
	t.AddEntry(symbol, t.nextAddr)
	t.nextAddr++
}

func (t *Symtable) Contains(symbol string) bool {
	_, ok := t.table[symbol]
	return ok
}

func (t *Symtable) GetAddress(symbol string) int {
	return t.table[symbol]
}
