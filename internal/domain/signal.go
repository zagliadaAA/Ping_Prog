package domain

type Signal struct {
	ID      int
	Address string
	Port    int
}

func NewSignal(address string, port int) *Signal {
	return &Signal{
		Address: address,
		Port:    port,
	}
}
