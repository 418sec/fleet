package registry

import (
	"path"

	"github.com/coreos/fleet/unit"
)

const (
	statePrefix = "/state/"
)

// Get the current UnitState of the provided Job's Unit
func (r *Registry) getUnitState(jobName string) *unit.UnitState {
	key := path.Join(r.keyPrefix, statePrefix, jobName)
	resp, err := r.etcd.Get(key, false, true)

	// Assume the error was KeyNotFound and return an empty data structure
	if err != nil {
		return nil
	}

	var state unit.UnitState
	//TODO: Handle the error generated by unmarshal
	unmarshal(resp.Node.Value, &state)
	return &state
}

// Persist the changes in a provided Machine's Job
func (r *Registry) SaveUnitState(jobName string, unitState *unit.UnitState) {
	key := path.Join(r.keyPrefix, statePrefix, jobName)
	//TODO: Handle the error generated by marshal
	json, _ := marshal(unitState)
	r.etcd.Set(key, json, 0)
}

// Delete the state from the Registry for the given Job's Unit
func (r *Registry) RemoveUnitState(jobName string) error {
	key := path.Join(r.keyPrefix, statePrefix, jobName)
	_, err := r.etcd.Delete(key, false)
	return err
}
