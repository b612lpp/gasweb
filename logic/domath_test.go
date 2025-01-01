package logic

import (
	"reflect"
	"testing"
)

func TestDoMath(t *testing.T) {
	type args struct {
		v int
		c int
	}
	tests := []struct {
		name           string
		args           args
		wantToFillTank ToFillTankType
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToFillTank, err := DoMath(tt.args.v, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoMath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotToFillTank, tt.wantToFillTank) {
				t.Errorf("DoMath() = %v, want %v", gotToFillTank, tt.wantToFillTank)
			}
		})
	}
}
