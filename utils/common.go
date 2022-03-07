package utils

func Catch(err error) {
	if err != nil {
		panic(err)
	}
}
