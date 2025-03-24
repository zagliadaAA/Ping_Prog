package domain

type Signal struct {
	ID      int
	Address string
	Port    int
	IDUser  int
}

func NewSignal(address string, port int) *Signal {
	return &Signal{
		Address: address,
		Port:    port,
	}
}
