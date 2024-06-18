package service

import "testing"

func TestJokeService(t *testing.T) {
	t.Run("baseAPI is empty", func(t *testing.T) {
		expectedErr := "invalid or broken baseAPI"
		_, err := GetJoke("")

		if err == nil {
			t.Fatalf("\nexpected %v\ngot	 %v", expectedErr, err)
		}
		if err != nil {
			t.Logf("\nexpected %v\ngot	 %v", expectedErr, err)
		}
	})
	// Starting to think if this is even needed
	// t.Run("no host in requested URL", func(t *testing.T) {
	// 	baseAPI := "http:/github.com/"
	// 	_, err := GetJoke(baseAPI)
	// 	expectedErr := "Could not make a request. Get " + baseAPI + ": http: no Host in request URL"

	// 	if err == nil {
	// 		t.Fatalf("\nexpected %v\ngot	%v", expectedErr, err)
	// 	}
	// 	if err != nil {
	// 		t.Logf("\nexpected %v\ngot	 %v", expectedErr, err)
	// 	}
	// })
}
