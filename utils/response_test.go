package utils

import (
	"reflect"
	"testing"
)

func TestSuccessResponse(t *testing.T) {
	data := map[string]interface{}{"foo": "bar"}
	resp := SuccessResponse(200, "ok", data)

	if resp.Status != 200 {
		t.Errorf("預期 Status 200，實際為 %d", resp.Status)
	}
	if resp.Message != "ok" {
		t.Errorf("預期 Message 'ok'，實際為 %s", resp.Message)
	}
	if !reflect.DeepEqual(resp.Data, data) {
		t.Errorf("預期 Data %v，實際為 %v", data, resp.Data)
	}
	if resp.Error != "" {
		t.Errorf("預期 Error 為空字串，實際為 %s", resp.Error)
	}
}

func TestErrorsResponse(t *testing.T) {
	resp := ErrorsResponse(400, "something wrong")

	if resp.Status != 400 {
		t.Errorf("預期 Status 400，實際為 %d", resp.Status)
	}
	if resp.Error != "something wrong" {
		t.Errorf("預期 Error 'something wrong'，實際為 %s", resp.Error)
	}
	// ErrorResponse does not have a Data field, so no Data check is needed here.
}
