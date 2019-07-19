package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_action_IsResponsible(t *testing.T) {
	type fields struct {
		isResponsibleFunc func(args []string) bool
		executeFunc       func() int
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "returns true when check-function returns true",
			fields: fields{isResponsibleFunc: func(args []string) bool {
				return true
			}},
			want: true,
		},
		{
			name: "returns false when check-function returns false",
			fields: fields{isResponsibleFunc: func(args []string) bool {
				return false
			}},
			want: false,
		},
		{
			name: "args are passed to check-function",
			args: args{args: []string{"1", "2"}},
			fields: fields{isResponsibleFunc: func(args []string) bool {
				return len(args) == 2
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &action{
				isResponsibleFunc: tt.fields.isResponsibleFunc,
				executeFunc:       tt.fields.executeFunc,
			}
			if got := a.IsResponsible(tt.args.args); got != tt.want {
				t.Errorf("action.IsResponsible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_action_Execute(t *testing.T) {
	resultStub := 10
	a := &action{
		isResponsibleFunc: func([]string) bool { return true },
		executeFunc: func() int {

			return resultStub
		},
	}
	result := a.Execute()
	assert.Exactly(t, resultStub, result)
}

type chainElementStub struct {
	isResponsible   bool
	executionResult int
}

func (m *chainElementStub) IsResponsible(args []string) bool {
	return m.isResponsible
}

func (m *chainElementStub) Execute() int {
	return m.executionResult
}

func Test_actionChain_execute(t *testing.T) {

	tests := []struct {
		name  string
		chain actionChain
		want  int
	}{
		{
			name: "1st chain element is executed",
			chain: actionChain{
				&chainElementStub{isResponsible: true, executionResult: 2},
				&chainElementStub{isResponsible: true, executionResult: 5},
			},
			want: 2,
		},
		{
			name: "2nd chain element is executed",
			chain: actionChain{
				&chainElementStub{isResponsible: false, executionResult: 2},
				&chainElementStub{isResponsible: true, executionResult: 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var noArgs []string
			if got := tt.chain.execute(noArgs); got != tt.want {
				t.Errorf("actionChain.execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_actionChain_execute_emptyChain_panics(t *testing.T) {
	chain := &actionChain{}
	assert.Panics(t, func() {
		chain.execute([]string{})
	})
}
