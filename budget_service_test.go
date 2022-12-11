package YileTDD

import (
	"testing"
	"time"
)



type FakeRepo struct {}

func (f FakeRepo) GetAll() []Budget {
	return []Budget{
		{
			yearMonth: "202212",
			amount:    310,
		},
		{
			yearMonth: "202211",
			amount:    30,
		},
		{
			yearMonth: "202210",
			amount:    3100,
		},
		{
			yearMonth: "202209",
			amount:    0,
		},
	}
}

func GivenFakeTime(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return t
}

func TestBudgetService_Query(t *testing.T) {
	type fields struct {
		br IBudgetRepo
	}
	type args struct {
		start time.Time
		end   time.Time
	}
	var tests = []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name:   "同日",
			fields: fields{br: FakeRepo{}},
			args:   args{
				start: GivenFakeTime("2022-12-01"),
				end:   GivenFakeTime("2022-12-01"),
			},
			want:   10,
		},
		{
			name:   "跨月",
			fields: fields{br: FakeRepo{}},
			args:   args{
				start: GivenFakeTime("2022-10-31"),
				end:   GivenFakeTime("2022-12-03"),
			},
			want:   160,
		},
		{
			name:   "單年",
			fields: fields{br: FakeRepo{}},
			args:   args{
				start: GivenFakeTime("2022-01-01"),
				end:   GivenFakeTime("2022-12-31"),
			},
			want:   3440,
		},
		{
			name:   "起訖相反",
			fields: fields{br: FakeRepo{}},
			args:   args{
				start: GivenFakeTime("2022-05-03"),
				end:   GivenFakeTime("2022-05-01"),
			},
			want:   0,
		},
		{
			name:   "沒資料",
			fields: fields{br: FakeRepo{}},
			args:   args{
				start: GivenFakeTime("2022-05-01"),
				end:   GivenFakeTime("2022-05-03"),
			},
			want:   0,
		},
		{
			name:   "跨月比例",
			fields: fields{br: FakeRepo{}},
			args:   args{
				start: GivenFakeTime("2022-11-30"),
				end:   GivenFakeTime("2022-12-03"),
			},
			want:   31,
		},
		{
			name:   "預算0",
			fields: fields{br: FakeRepo{}},
			args:   args{
				start: GivenFakeTime("2022-09-01"),
				end:   GivenFakeTime("2022-09-30"),
			},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := BudgetService{
				repo: tt.fields.br,
			}
			if got := bs.Query(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}