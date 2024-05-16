package utils

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
