package strategy

import "testing"

func TestContext_Do(t *testing.T) {
	context := &Context{}
	strategy1 := &Strategy1{}
	context.Strategy = strategy1
	context.Do()

	strategy2 := &Strategy2{}
	context.Strategy = strategy2
	context.Do()
}
