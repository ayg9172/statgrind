package main

const CallAllocationPerTick = 8

type Scheduling struct {
	combat        *Combat
	calls         map[int]func()
	futureCallIds map[int][]int
}

type Scheduler interface {
	Schedule(call func(), tick int)
}

func (scheduling *Scheduling) Schedule(callable func(), tick int) func() {
	id := scheduling.combat.NextId()
	scheduling.calls[id] = callable
	_, exists := scheduling.futureCallIds[tick]
	if !exists {
		scheduling.futureCallIds[tick] = make([]int, CallAllocationPerTick)
	}
	scheduling.futureCallIds[tick] = append(scheduling.futureCallIds[tick], id)

	return func() {
		scheduling.Cancel(id)
	}
}

func (scheduling *Scheduling) ContainsCall(id int) bool {
	_, exists := scheduling.calls[id]
	return exists
}

func (scheduling *Scheduling) Dispatch() {
	callIds, exists := scheduling.futureCallIds[scheduling.combat.CurrentTick]
	if !exists {
		return
	}
	for i := range callIds {
		id := callIds[len(callIds)-i-1]

		if scheduling.ContainsCall(id) {
			scheduling.calls[id]()
			delete(scheduling.futureCallIds, id)
		}
	}
	delete(scheduling.futureCallIds, scheduling.combat.CurrentTick)
}

func (scheduling *Scheduling) Cancel(callableId ...int) {
	for _, id := range callableId {
		if scheduling.ContainsCall(id) {
			delete(scheduling.calls, id)
		}
	}
}
