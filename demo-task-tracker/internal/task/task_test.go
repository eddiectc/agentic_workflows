package task

import "testing"

func TestNewTask(t *testing.T) {
	tests := []struct {
		name    string
		desc    string
		wantErr bool
	}{
		{
			name:    "valid description",
			desc:    "Buy groceries",
			wantErr: false,
		},
		{
			name:    "empty description",
			desc:    "",
			wantErr: true,
		},
		{
			name:    "whitespace only",
			desc:    "   ",
			wantErr: true,
		},
		{
			name:    "description with leading/trailing spaces",
			desc:    "  Trim me  ",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := NewTask(tt.desc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if task.Done {
					t.Error("new task should not be done")
				}
				if task.ID == "" {
					t.Error("new task should have an ID")
				}
			}
		})
	}
}
