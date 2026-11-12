package provider

import "testing"

func TestOriginalCreationCaptureDefaultPlatformBehavior(t *testing.T) {
	t.Setenv("XCSH_ORIGINAL_CREATION_SCOPE_FILE", "")
	capture, err := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
	if err != nil || capture != nil {
		t.Fatal("disabled capture changed ordinary provider behavior")
	}
	capture.Close()
}
