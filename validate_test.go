// Copyright 2019 minigo Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package minigo

import "testing"

func TestValidateDo(t *testing.T) {
	userInfo := getTUserInfo()

	validate := NewValidate()
	err := validate.Do(userInfo)

	if err != nil {
		t.Error(err)
	}
}
