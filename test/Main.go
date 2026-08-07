package Test_Main

func ArrayFrom1UpTo(n int) []interface{} {
	result := make([]interface{}, n)
	for i := 0; i < n; i++ {
		result[i] = i + 1
	}
	return result
}

func ArrayReplicate(n int) func(interface{}) []interface{} {
	return func(x interface{}) []interface{} {
		result := make([]interface{}, n)
		for i := 0; i < n; i++ {
			result[i] = x
		}
		return result
	}
}

func MkNEArray(nothing interface{}) func(interface{}) func(gopurs_runtime.Value) interface{} {
	return func(just interface{}) func(gopurs_runtime.Value) interface{} {
		return func(arr gopurs_runtime.Value) interface{} {
			if gopurs_runtime.ArrayLength(arr) > 0 {
				return gopurs_runtime.Apply(just.(gopurs_runtime.Value), arr)
			}
			return nothing
		}
	}
}

func FoldMap1NEArray(appendFn interface{}) func(interface{}) func(gopurs_runtime.Value) interface{} {
	return func(f interface{}) func(gopurs_runtime.Value) interface{} {
		return func(arr gopurs_runtime.Value) interface{} {
			fVal := f.(gopurs_runtime.Value)
			appendVal := appendFn.(gopurs_runtime.Value)
			
			acc := gopurs_runtime.Apply(fVal, gopurs_runtime.ArrayAccess(arr, 0))
			for i := 1; i < gopurs_runtime.ArrayLength(arr); i++ {
				mapped := gopurs_runtime.Apply(fVal, gopurs_runtime.ArrayAccess(arr, i))
				acc = gopurs_runtime.Apply2(appendVal, acc, mapped)
			}
			return acc
		}
	}
}
