package Data_Traversable

import "gopurs/output/gopurs_runtime"

func TraverseArrayImpl(
	apply func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value,
	mapFn func(func(gopurs_runtime.Value) gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value,
	pure func(gopurs_runtime.Value) gopurs_runtime.Value,
	concat2 func(gopurs_runtime.Value) func(gopurs_runtime.Value) gopurs_runtime.Value,
	f func(gopurs_runtime.Value) gopurs_runtime.Value,
	arrayVal []gopurs_runtime.Value,
) gopurs_runtime.Value {
	array1 := func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Array([]gopurs_runtime.Value{a})
	}
	array2 := func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array([]gopurs_runtime.Value{a, b})
		})
	}
	array3 := func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Array([]gopurs_runtime.Value{a, b, c})
			})
		})
	}
	var goFn func(int, int) gopurs_runtime.Value
	goFn = func(bot, top int) gopurs_runtime.Value {
		switch top - bot {
		case 0:
			return pure(gopurs_runtime.Array([]gopurs_runtime.Value{}))
		case 1:
			return mapFn(array1, f(arrayVal[bot]))
		case 2:
			return apply(mapFn(array2, f(arrayVal[bot])), f(arrayVal[bot+1]))
		case 3:
			return apply(apply(mapFn(array3, f(arrayVal[bot])), f(arrayVal[bot+1])), f(arrayVal[bot+2]))
		default:
			pivot := bot + ((top-bot)/4)*2
			return apply(mapFn(func(x gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(concat2(x))
			}, goFn(bot, pivot)), goFn(pivot, top))
		}
	}
	return goFn(0, len(arrayVal))
}
