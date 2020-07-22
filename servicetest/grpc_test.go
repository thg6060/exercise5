package test

import "testing"

func TestServerParallel(t *testing.T) {
    // This Run will not return until the parallel tests finish.
        t.Run("Test1", TestRunServer)
    // <tear-down code>
}

func TestClientParallel(t *testing.T) {
    // This Run will not return until the parallel tests finish.
        t.Run("Test1", TestRunClient)
    // <tear-down code>
}