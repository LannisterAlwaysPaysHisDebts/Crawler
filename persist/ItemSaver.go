package persist

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		// todo: save
	}()
	return out
}
