package unitparser

import (
    "testing"
)

func TestOhmParsing(t *testing.T) {
    unit, ok := ParseString(" 1 кОм ")

    if !ok {
        t.Error("Parsing failed")
    } else if unit.String() != "1 kΩ" {
        t.Errorf("%s != 1 kΩ", unit)
    }
}

func TestAmpParsing(t *testing.T) {
    unit, ok := ParseString(" 10 мкА")

    if !ok {
        t.Error("Parsing failed")
    } else if unit.String() != "10 μA" {
        t.Errorf("%s != 10 μA", unit)
    }
}

func TestVoltParsing(t *testing.T) {
    unit, ok := ParseString("230  В")

    if !ok {
        t.Error("Parsing failed")
    } else if unit.String() != "230 V" {
        t.Errorf("%s != 1 V", unit)
    }
}

func TestHzParsing(t *testing.T) {
    unit, ok := ParseString("3 ГГц")

    if !ok {
        t.Error("Parsing failed")
    } else if unit.String() != "3 GHz" {
        t.Errorf("%s != 3 GHz", unit)
    }
}

func TestHenryParsing(t *testing.T) {
    unit, ok := ParseString(" 13 Гн ")

    if !ok {
        t.Error("Parsing failed")
    } else if unit.String() != "13 H" {
        t.Errorf("%s != 13 H", unit)
    }
}

func TestFaradParsing(t *testing.T) {
    unit, ok := ParseString(" 4700пФ ")

    if !ok {
        t.Error("Parsing failed")
    } else if unit.String() != "4700 pF" {
        t.Errorf("%s != 4700 pF", unit)
    }
}
