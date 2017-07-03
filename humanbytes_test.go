package humanbytes

import (
	"testing"
)

func TestParseKiloBytes(t *testing.T) {
	b, err := ParseBytes("1KB")
	if b != 1024 && err == nil {
		t.Error("1KB should be 1024 Bytes")
	}
}

func TestParseKiloBytesSpaces(t *testing.T) {
	b, err := ParseBytes("1 KB")
	if b != 1024 && err == nil {
		t.Error("1 KB should be 1024 Bytes")
	}
}

func TestParseMegaBytes(t *testing.T) {
	b, err := ParseBytes("1MB")
	if b != 1024*1024 && err == nil {
		t.Error("1MB should be 1024 Bytes")
	}
}

func TestParseMegaBytesSpaces(t *testing.T) {
	b, err := ParseBytes("1 MB")
	if b != 1048576 && err == nil {
		t.Error("1 MB should be 1048576 Bytes")
	}
}

func TestSprintKiloBytes(t *testing.T) {
	s, err := Sprint(1024, "KB")
	if s != "1 KB" && err == nil {
		t.Error("1024 byteas should display as '1 KB'")
		t.Errorf("Received: [%s]", s)
	}
}

func TestSprintMegaBytes(t *testing.T) {
	s, err := Sprint(1024, "MB")
	if s != "0.0009765625 MB" && err == nil {
		t.Error("1024 byteas should display as '0.000000 MB'")
		t.Errorf("Received: [%s]", s)
	}
}

func TestSprintKiloBytesPoint(t *testing.T) {
	s, err := Sprint(1536, "KB")
	if s != "1.5 KB" && err == nil {
		t.Error("1536 byteas should display as '1.5 KB'")
		t.Errorf("Received: [%s]", s)
	}
}

func TestKBtoMB(t *testing.T) {
	mb, err := Convert(1048576, "MB")
	if mb != 1.0 && err == nil {
		t.Error("1048576 byteas should convert to 1.0 MB")
		t.Errorf("Received: [%v]", mb)
	}
}

func TestKBtoMBPoint(t *testing.T) {
	mb, err := Convert(1572864, "MB")
	if mb != 1.5 && err == nil {
		t.Error("1572864 byteas should convert to 1.5 MB")
		t.Errorf("Received: [%v]", mb)
	}
}

func TestConvertKBtoBytes(t *testing.T) {
	mb, err := ConvertToBytes(1, "KB")
	if mb != 1024 || err != nil {
		t.Error("1KB should convert to 1024 bytes")
		t.Errorf("Received: [%v]", mb)
	}
}

// Error handling tests

func TestParseBytesInvalidUnit(t *testing.T) {
	b, err := ParseBytes("1BK")
	if err != nil && b != 0 {
		t.Error("1BK should have caused an error")
	}
}

func TestParseBytesInvalidNumber(t *testing.T) {
	b, err := ParseBytes("x KB")
	if err != nil && b != 0 {
		t.Error("x KB should have caused an error")
	}
}

func TestParseBytesInvalidNumber2(t *testing.T) {
	b, err := ParseBytes(".6 KB")
	if err != nil && b != 0 {
		t.Error(".6 KB should have caused an error")
	}
}

func TestConvertKBtoBytesInvalidUnit(t *testing.T) {
	mb, err := ConvertToBytes(1, "kb")
	if mb != 0 || err == nil {
		t.Error("1 kb should cause an error")
		t.Errorf("Received: [%v]", mb)
	}
}

func TestConvertInvalidUnit(t *testing.T) {
	mb, err := Convert(1, "BM")
	if mb != 0 || err == nil {
		t.Error("The unit BM shoudl have caused an error")
		t.Errorf("Received: [%v]", mb)
	}
}

func TestSprintInvalidUnit(t *testing.T) {
	s, err := Sprint(1024, "BM")
	if s != "0.0009765625 MB" && err == nil {
		t.Error("The unit BM shoudl have caused an error")
		t.Errorf("Received: [%s]", s)
	}
}
