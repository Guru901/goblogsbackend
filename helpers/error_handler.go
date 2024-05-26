package helpers

func HandleError(error error) {
	if error != nil {
		panic(error)
	}
}
