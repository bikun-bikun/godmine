package flags

import "github.com/urfave/cli"

var Profile = cli.StringFlag{
	Name:  "profile, p",
	Usage: "Change Profile",
	Value: "default",
}

var Create = cli.StringFlag{
	Name:  "create, c",
	Usage: "Create",
	Value: "",
}

var Endpoit = cli.StringFlag{
	Name:  "endpoint, ep",
	Value: "",
}

var ApiKey = cli.StringFlag{
	Name:  "apikey, ak",
	Value: "",
}

var Name = cli.StringFlag{
	Name:  "name, n",
	Value: "default",
}
