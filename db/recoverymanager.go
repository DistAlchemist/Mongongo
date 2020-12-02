// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package db

// RecoveryManager manages recovery
type RecoveryManager struct {
	//
}

var recoveryManagerInstance *RecoveryManager

// GetRecoveryManager fetch existing instance or
// create a new one.
func GetRecoveryManager() *RecoveryManager {
	if recoveryManagerInstance == nil {
		recoveryManagerInstance = NewRecoveryManager()
	}
	return recoveryManagerInstance
}

// NewRecoveryManager creates a new instance
func NewRecoveryManager() *RecoveryManager {
	r := &RecoveryManager{}
	return r
}

func (r *RecoveryManager) doRecovery() {
	// TODO
}
