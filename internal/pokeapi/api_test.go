package pokeapi

// import (
// 	"bytes"
// 	"io"
// 	"os"
// 	"strings"
// 	"testing"
// )

// func TestAPIGet(t *testing.T) {

// 	// Store output in a buffer
// 	r, w, _ := os.Pipe()
// 	os.Stdout = w // redirect stdout to buffer

// 	config := &config{
// 		Next:     pokeAPIUrl,
// 		Previous: "",
// 	}
// 	err := commandMap(config)
// 	if err != nil {
// 		t.Fatalf("commandMap(config.Next) returned err: %v", err)
// 	}

// 	var buf bytes.Buffer
// 	io.Copy(&buf, r)
// 	w.Close()

// 	expectedOutputContains := "mt-coronet-2f"
// 	if !strings.Contains(buf.String(), expectedOutputContains) {
// 		t.Fatalf("buffer did not contain %s", expectedOutputContains)
// 	}

// }
