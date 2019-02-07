package ports

import (
    "testing"
)

func TestIsOpen(t *testing.T) {
    if(!IsOpen(BEGIN_RANGE)) {
        t.Error("The port should be open.")
    }
}

func TestListUsedPorts(t *testing.T) {
    slice := ListUsedPorts()
    if(len(slice) > 0) {
        t.Error("There should be no used ports.")
    }
}

func TestNextAvailPort(t *testing.T) {
    port := NextAvailPort()
    if(port != BEGIN_RANGE) {
        t.Error("First of range should be available.")
    }
}
