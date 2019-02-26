package cmd

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestReadConfigFromViper(t *testing.T) {
	expectedKey := "key-viper"
	expectedSecret := "secret-viper"
	viper.Set("application_key", expectedKey)
	viper.Set("application_secret", expectedSecret)

	defer viper.Set("application_key", nil)
	defer viper.Set("application_secret", nil)

	key, secret, err := readCredentials(os.Stdin)

	if err != nil {
		t.Fatal(err)
	}

	if key != expectedKey {
		t.Errorf("expected key to be %q; got %q", expectedKey, key)
	}

	if secret != expectedSecret {
		t.Errorf("expected secret to be %q; got %q", expectedSecret, secret)
	}
}

func TestReadConfigFromReader(t *testing.T) {
	tests := []struct {
		key    string
		secret string
		err    string
	}{
		{"key", "secret", ""},
		{"", "", "key not given"},
		{"key", "", "secret not given"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			defer func(f *os.File) {
				os.Stdout = f
			}(os.Stdout)

			os.Stdout, _ = os.Open(os.DevNull)

			r := strings.NewReader(fmt.Sprintf("%s\n%s\n", tt.key, tt.secret))

			key, secret, err := readCredentials(r)

			if tt.err == "" {
				if err != nil {
					t.Fatal(err)
				}

				if key != tt.key {
					t.Errorf("expected key to be %q; got %q", tt.key, key)
				}

				if secret != tt.secret {
					t.Errorf("expected secret to be %q; got %q", tt.secret, secret)
				}

				return
			}

			if err == nil {
				t.Error("expected error, none received")
			}

			if err.Error() != tt.err {
				t.Errorf("expected error %q; got %q", tt.err, err.Error())
			}
		})
	}
}
