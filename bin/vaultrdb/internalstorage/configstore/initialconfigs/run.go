package initialconfigs

func SetInitialConfigsIfNotConfiguredAlready() (err error) {

	for _, f := range funcList {
		if err = f(); err != nil {
			return
		}
	}

	return nil
}

var funcList = []func() error{
	calculateencryptionpassphrase,
}
