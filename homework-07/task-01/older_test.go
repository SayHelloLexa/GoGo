package task01

import "testing"

func TestGetOlder(t *testing.T) {
	type args struct {
		u []User
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				u: []User{
					&Employee{
						Age: 20,
					},
					&Customer{
						Age: 30,
					},
					&Customer{
						Age: 40,
					},
					&Employee{
						Age: 50,
					},
				},
			},

			want: 50,
		},
		
		{
			name: "test 2",
			args: args{
				u: []User{
					&Employee{
						Age: 0,
					},
				},
			},

			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOlder(tt.args.u...); got != tt.want {
				t.Errorf("GetOlder() = %v, want %v", got, tt.want)
			}
		})
	}
}
