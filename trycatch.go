package utils

import "fmt"

// TryCatch run func in firs parameter, if panic happens recover in catch function as error
func TryCatch(try func(), catch func(err error)) {
	defer func() {
		if r := recover(); r != nil {
			catch(fmt.Errorf("%v", r))
		}
	}()

	try()
}
