package main

import (
	"reflect"
	"testing"
)

func TestGame_GetPawns(t *testing.T) {
	type args struct {
		c Color
	}
	tests := []struct {
		name string
		g    Game
		args args
		want Bitboard
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.GetPawns(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Game.GetPawns() = %v, want %v", got, tt.want)
			}
		})
	}
}
