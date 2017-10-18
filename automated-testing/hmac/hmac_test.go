package main

import "testing"

func TestSign(t *testing.T) {

	signingKey := "secret"

	cases := []struct {
		name          string
		input         string
		expectedOuput string
	}{
		{
			name:          "valid",
			input:         "abc",
			expectedOuput: "mUba1OAOkT_Ivo5dP34RCkqegy-D-wnDRShdeGONig4=",
		},
		{
			name:          "invalid",
			input:         "cba",
			expectedOuput: "8jASuwLmkC3YdSxwfJ1jYlgDPTQBQkLM592MAnXuJHg=",
		},
	}

	for _, c := range cases {
		output, err := sign(signingKey, c.input)
		if err != nil {
			t.Errorf("error signing: %v", err)
		}
		if output != c.expectedOuput {
			t.Errorf("\ncase: %s\ninput: %s\ngot: %s\nwant: %s", c.name, c.input, output, c.expectedOuput)
		}
	}
}

func TestVerify(t *testing.T) {
	signingKey := "secret"

	cases := []struct {
		name          string
		sig           string
		msg           string
		expectedOuput bool
	}{
		{
			name:          "message not altered",
			sig:           "mUba1OAOkT_Ivo5dP34RCkqegy-D-wnDRShdeGONig4=",
			msg:           "abc",
			expectedOuput: true,
		},
		{
			name:          "message altered",
			sig:           "mUba1OAOkT_Ivo5dP34RCkqegy-D-wnDRShdeGONig4=",
			msg:           "cba",
			expectedOuput: false,
		},
	}

	for _, c := range cases {
		output, err := verify(signingKey, c.sig, c.msg)
		if err != nil {
			t.Errorf("error verifying: %v", err)
		}
		if output != c.expectedOuput {
			t.Errorf("\ncase: %s\nsignature: %s\nmessage: %s\ngot: %v\nwant: %v", c.name, c.sig, c.msg, output, c.expectedOuput)
		}
	}
}
