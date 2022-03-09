package helpers

func Catch(err error) {
	if err != nil {
		panic(err)
	}
}
