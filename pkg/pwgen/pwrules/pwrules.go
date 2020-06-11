package pwrules

//go:generate go run gen.go

var (
	rules = map[string]rule{}
)

type rule struct {
	minlen    int
	maxlen    int
	required  []string
	allowed   []string
	maxconsec int
}
