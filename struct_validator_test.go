package utils

import (
	"testing"
)

func TestValidate(t *testing.T) {
	type address struct {
		Street   string `json:"street" required:"true" minlength:"3" `
		Province string `json:"province" required:"true" format:"alphabet" maxlength:"5"`
	}
	type education struct {
		Level    string `json:"level" required:"true" maxlength:"5"`
		Instance string `json:"instance" required:"true" format:"alphabet"`
		Detail   struct {
			NIS string `json:"nis" required:"true" format:"numeric"`
		} `json:"detail" required:"true"`
	}
	type People struct {
		ID        string      `json:"id" required:"true" format:"numeric"`
		Name      string      `json:"name" required:"true"`
		Alias     string      `json:"alias" required:"true" format:"alphanumeric"`
		Birthday  string      `json:"birthday" required:"true" format:"date,yyyy-mm-dd"`
		School    string      `json:"school" maxlength:"5"`
		Gender    string      `json:"gender" in:"L,P"`
		Address   *address    `json:"address"`
		Education []education `json:"education"`
	}
	type fake1 struct {
		Fk string `json:"fk" format:"inv"`
	}
	type fake2 struct {
		Datefake string `json:"dateFake" format:"date,inv"`
	}

	tests := []struct {
		name    string
		args    interface{}
		wantErr bool
	}{
		{
			name: "Testcase #1: Positive",
			args: People{
				ID: "0912", Name: "AgungDP", Alias: "agungdp22", Birthday: "2018-06-05",
				Education: []education{education{
					Level: "S1", Instance: "IPB", Detail: struct {
						NIS string `json:"nis" required:"true" format:"numeric"`
					}{NIS: "873892374"},
				}},
			},
			wantErr: false,
		},
		{
			name:    "Testcase #2: Negative",
			args:    []education{education{Level: "S111111111"}},
			wantErr: true,
		},
		{
			name:    "Testcase #3: Negative",
			args:    address{Street: "a", Province: "jdshjsdf jdhfjds kjdfhjdsfhakas"},
			wantErr: true,
		},
		{
			name:    "Testcase #4: Negative",
			args:    &People{Gender: "undefined"},
			wantErr: true,
		},
		{
			name:    "Testcase #5: Negative, undefined tag",
			args:    fake1{Fk: "inv"},
			wantErr: true,
		},
		{
			name:    "Testcase #6: Negative, undefined date format",
			args:    fake2{Datefake: "inv"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator()
			err := validator.Validate(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("\x1b[31;1m ValidateModel() error = %v, wantErr %v\x1b[0m", err, tt.wantErr)
			}
		})
	}
}
