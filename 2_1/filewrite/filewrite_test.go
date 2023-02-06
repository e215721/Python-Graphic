package filewrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("LuongBao")
    want := "Hi, LuongBao. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}
