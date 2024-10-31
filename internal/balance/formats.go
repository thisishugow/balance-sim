package balance

type Format struct {
	Prefix    string
	Suffix    string
	MinWeight float64
	MaxWeight float64
	Precision int
}

var Formats = map[string]Format{
	"mettler": {
		Prefix:    "S S     ",
		Suffix:    " g\r\n",
		MinWeight: 0.0001,
		MaxWeight: 200.0000,
		Precision: 4,
	},
	"sartorius": {
		Prefix:    "    ",
		Suffix:    " g     \r\n",
		MinWeight: 0.001,
		MaxWeight: 300.000,
		Precision: 3,
	},
}
