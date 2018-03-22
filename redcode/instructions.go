package redcode

// AddressMode identifies operand addressing modes
type AddressMode int

// Opcode identifies opcodes
type Opcode int

// Operation identifies compile-time operations
type Operation int

// Modifier identifies the instruction modifier
type Modifier int

// enumerated address mode values
const (
	Immediate AddressMode = iota
	Relative
	Indirect
	DecrementIndirect
)

// enumerated opcode values
const (
	OpDat Opcode = iota
	OpMov
	OpAdd
	OpSub
	OpJmp
	OpJmz
	OpDjn
	OpCmp
	OpSpl
	OpEnd
)

// enumerated operations
const (
	Number Operation = iota
	Label
	Add
	Sub
)

// enumerate modifiers
const (
	ModUnknown Modifier = iota
)

// Instruction is a single redcode instruction
type Instruction struct {
	Label    string
	Opcode   Opcode
	Modifier Modifier
	A, B     Operand
}

// Operand is a single redcode operand
type Operand struct {
	Mode       AddressMode
	Expression *Expression
}

// Expression is an operand expression
type Expression struct {
	Operation   Operation
	Number      int
	Label       string
	Left, Right *Expression
}
