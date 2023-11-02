package ha

var (
	// high availability configs, if ha is acitvated
	// the the control node will not process the configs
	// or sync the configs from the database
	// the default is FALSE
	EnableHA bool
)

func StartManager() {

	if EnableHA {

	}
}
