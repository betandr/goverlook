package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

var mazeJsonStr = `
{
	"cells": [
		[{
			"northRoute": false,
			"westRoute": false
		}, {
			"northRoute": true,
			"westRoute": false
		}],
		[{
			"northRoute": false,
			"westRoute": false
		}, {
			"northRoute": true,
			"westRoute": true
		}]
	]
}
`

func TestMainLoadCreatesNonZeroOutput(t *testing.T) {

	var buf bytes.Buffer
	out := bufio.NewWriter(&buf)
	in := strings.NewReader(mazeJsonStr)

	run(map[string]string{}, out, in, true)
	out.Flush()

	if len(buf.Bytes()) <= 0 {
		t.Errorf("output size error: %d bytes", len(buf.Bytes()))
	}
}

func TestMainLoadWithBadJSONDoesNotPanic(t *testing.T) {

	var buf bytes.Buffer
	out := bufio.NewWriter(&buf)
	in := strings.NewReader("NOT-JSON")

	run(map[string]string{}, out, in, true)
}

func TestMainGenerateCreatesNonZeroOutput(t *testing.T) {

	var buf bytes.Buffer
	out := bufio.NewWriter(&buf)
	in := bufio.NewReader(os.Stdin)

	run(map[string]string{}, out, in, false)
	out.Flush()

	if len(buf.Bytes()) <= 0 {
		t.Errorf("output size error: %d bytes", len(buf.Bytes()))
	}
}

func TestMainWithWidthCreatesNonZeroOutput(t *testing.T) {

	var buf bytes.Buffer
	out := bufio.NewWriter(&buf)

	in := bufio.NewReader(os.Stdin)

	args := map[string]string{"width": "10", "height": "10"}

	run(args, out, in, false)
	out.Flush()

	if len(buf.Bytes()) <= 0 {
		t.Errorf("output size error: %d bytes", len(buf.Bytes()))
	}
}
func TestMainWithJSONContainsCellsField(t *testing.T) {
	var buf bytes.Buffer
	out := bufio.NewWriter(&buf)
	in := bufio.NewReader(os.Stdin)

	args := map[string]string{"out": "json"}
	run(args, out, in, false)
	out.Flush()

	var outputMap map[string]json.RawMessage
	err := json.Unmarshal(buf.Bytes(), &outputMap)
	if err != nil {
		t.Errorf(err.Error())
	}

	if _, ok := outputMap["cells"]; !ok {
		t.Errorf("maze output json does not contain 'cells' field")
	}
}
