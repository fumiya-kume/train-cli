package main

import (
    "bytes"
    "testing"
)

func TestPrintTrain(t *testing.T) {
    // Capture the output of printTrain
    var buf bytes.Buffer
    old := out
    out = &buf
    defer func() { out = old }()

    printTrain()

    // Expected ASCII art of the modern train
    expected := "        ____\n  __,-~~/~    `---.\n_/_,---(      ,    )\n __/        <    /   )\n  ~~-~~/`---'~`--~\n"

    if buf.String() != expected {
        t.Errorf("Expected %q but got %q", expected, buf.String())
    }
}
