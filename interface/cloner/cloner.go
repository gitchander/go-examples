package cloner

type Cloner interface {
	Clone() Cloner
}
