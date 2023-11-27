package interfaces 

type SetextAble interface {
	SetText(text string)
}

type AddressAble interface {
	Address() string
}

type MethodAble interface {
	Method() string
}

type HeaderListAble interface {
	HeaderList() []string
}