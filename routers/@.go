package routers

type Delegate struct {
	G      map[string][]string
	VarSet *VarSet
}

type VarSet struct {
	Num  int
	Msg  string
	Nums []int
}

func (d *Delegate) Vars() *VarSet {
	return d.VarSet
}

func (d *Delegate) Add(key, value string) {
	d.G[key] = append(d.G[key], value)
}

func (d *Delegate) Del(key string) {
	delete(d.G, key)
}

func (d *Delegate) Get(key string) string {
	return d.G[key][0]
}

func (d *Delegate) Set(key, value string) {
	d.G[key] = []string{value}
}
