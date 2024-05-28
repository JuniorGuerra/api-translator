package service

import "testing"

func TestTranslatorService(t *testing.T) {
	type args struct {
		info Info
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "", args: struct{ info Info }{info: Info{
			CurrentLang: "es",
			ToLang:      "en",
			Text:        "Hola",
		}}, want: "Hello", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TranslatorService(tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("TranslatorService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TranslatorService() got = %v, want %v", got, tt.want)
			}
		})
	}
}
