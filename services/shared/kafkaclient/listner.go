package kafkaclient

type Listner struct {
}

type ListnerConfig struct {
}

func NewListener(config *ListnerConfig) Listner {
	return Listner{}
}

func (l *Listner) Start() {

}
