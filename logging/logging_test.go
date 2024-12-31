package logging

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestLogger_Info(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := NewLogger("TEST", INFO, buffer)
	logger.Info("This is an info message")

	output := buffer.String()
	if !strings.Contains(output, "INFO") {
		t.Errorf("Expected output to contain 'INFO', got: %s", output)
	}
	if !strings.Contains(output, "This is an info message") {
		t.Errorf("Expected output to contain the message, got: %s", output)
	}
}

func TestLogger_Debug_Filtering(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := NewLogger("TEST", INFO, buffer)
	logger.Debug("This debug message should not appear")

	output := buffer.String()
	if output != "" {
		t.Errorf("Expected no output for DEBUG level below minLevel INFO, got: %s", output)
	}
}

func TestLogger_MultipleLevels(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := NewLogger("TEST", DEBUG, buffer)

	logger.Info("Info message")
	logger.Debug("Debug message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	output := buffer.String()
	if !strings.Contains(output, "INFO") {
		t.Errorf("Expected output to contain 'INFO', got: %s", output)
	}
	if !strings.Contains(output, "DEBUG") {
		t.Errorf("Expected output to contain 'DEBUG', got: %s", output)
	}
	if !strings.Contains(output, "WARN") {
		t.Errorf("Expected output to contain 'WARN', got: %s", output)
	}
	if !strings.Contains(output, "ERROR") {
		t.Errorf("Expected output to contain 'ERROR', got: %s", output)
	}
}

func TestLogger_OutputFormatting(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := NewLogger("TEST", INFO, buffer)
	logger.Info("Formatted output test")

	output := buffer.String()
	if !strings.Contains(output, "[INFO]") {
		t.Errorf("Expected output to contain '[INFO]', got: %s", output)
	}
	if !strings.Contains(output, "TEST") {
		t.Errorf("Expected output to contain prefix 'TEST', got: %s", output)
	}
	if !strings.Contains(output, "Formatted output test") {
		t.Errorf("Expected output to contain message, got: %s", output)
	}
	if !strings.Contains(output, time.Now().Format("2006-01-02")) {
		t.Errorf("Expected output to contain current date, got: %s", output)
	}
}

func TestLogger_NilOutput(t *testing.T) {
	logger := NewLogger("TEST", INFO, nil)
	buffer := &bytes.Buffer{}
	logger.Info("Testing with nil output")
	output := buffer.String()
	if output != "" {
		t.Errorf("Expected no output when output is nil, got: %s", output)
	}
}

func TestLogger_WarnLevel(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := NewLogger("TEST", WARN, buffer)

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warn message")
	logger.Error("Error message")

	output := buffer.String()
	if strings.Contains(output, "DEBUG") {
		t.Errorf("DEBUG messages should not appear at WARN level, got: %s", output)
	}
	if strings.Contains(output, "INFO") {
		t.Errorf("INFO messages should not appear at WARN level, got: %s", output)
	}
	if !strings.Contains(output, "WARN") {
		t.Errorf("Expected output to contain 'WARN', got: %s", output)
	}
	if !strings.Contains(output, "ERROR") {
		t.Errorf("Expected output to contain 'ERROR', got: %s", output)
	}
}

func TestLogger_ErrorOnly(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := NewLogger("TEST", ERROR, buffer)

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warn message")
	logger.Error("Error message")

	output := buffer.String()
	if output != "" && !strings.Contains(output, "ERROR") {
		t.Errorf("Expected only 'ERROR' messages, got: %s", output)
	}
}
