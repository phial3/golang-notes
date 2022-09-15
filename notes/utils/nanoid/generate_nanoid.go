package nanoid

// NextId generate 10 bytes id with string
// alphabet form 0-9
func NextId() string {
	f, err := CustomASCII("0123456789", 10)
	if err != nil {
		panic(err)
	}

	return f()
}
