package routinepool

import "github.com/panjf2000/ants/v2"

func NewRoutinePool(cap int) *ants.Pool {
	p, err := ants.NewPool(cap)
	if err != nil {
		panic(err)
	}
	return p
}
