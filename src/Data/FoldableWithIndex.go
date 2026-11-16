import r "gopurs/output/gopurs_runtime"

func FoldlWithIndexImpl(f func(r.Value, r.Value, r.Value) r.Value, z r.Value, xs []r.Value) r.Value {
	acc := z
	for i, v := range xs {
		acc = f(r.Int(int64(i)), acc, v)
	}
	return acc
}

func FoldrWithIndexImpl(f func(r.Value, r.Value, r.Value) r.Value, z r.Value, xs []r.Value) r.Value {
	acc := z
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(r.Int(int64(i)), xs[i], acc)
	}
	return acc
}
