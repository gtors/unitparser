package unitparser

import (
    "regexp"
    "strconv"
    "strings"
)

const (
    // TODO: float numbers
    digitSubpattern = `(-?\d+)`
    prefixSubpattern = `(да|д|г|с|к|м|М|мк|Г|н|Т|п|П|ф|Э|а|З|з|И|и|Y|Z|E|P|T|G|M|k|h|da|d|c|m|μ|n|p|f|a|z|y)?`
    commonPatternStart = `^` + digitSubpattern  + `\s*` + prefixSubpattern
)

var (
    parsers = make([]Parser, 0, 6)[:0]
    ohmPattern = regexp.MustCompile(commonPatternStart + `(?:Ом|Ohm|Ω)$`)
    ampPattern = regexp.MustCompile(commonPatternStart + `(?:А|A)$`)
    faradPattern = regexp.MustCompile(commonPatternStart + `(?:Ф|F)$`)
    henryPattern = regexp.MustCompile(commonPatternStart + `(?:Гн|H)$`)
    hzPattern = regexp.MustCompile(commonPatternStart + `(?:Гц|Hz)$`)
    voltPattern = regexp.MustCompile(commonPatternStart + `(?:В|V)$`)
)

type Parser func(s string) (*UnitQuantity, bool)

func init() {
    parsers = append(parsers,
        parseOhm,
        parseAmp,
        parseVolt,
        parseFarad,
        parseHenry,
        parseHertz)
}

func parseOhm(s string) (*UnitQuantity, bool) {
    return commonParser(s, ohmPattern, UNIT_OHM)
}

func parseAmp(s string) (*UnitQuantity, bool) {
    return commonParser(s, ampPattern, UNIT_AMPERE)
}

func parseVolt(s string) (*UnitQuantity, bool) {
    return commonParser(s, voltPattern, UNIT_VOLT)
}

func parseFarad(s string) (*UnitQuantity, bool) {
    return commonParser(s, faradPattern, UNIT_FARAD)
}

func parseHertz(s string) (*UnitQuantity, bool) {
    return commonParser(s, hzPattern, UNIT_HERTZ)
}

func parseHenry(s string) (*UnitQuantity, bool) {
    return commonParser(s, henryPattern, UNIT_HENRY)
}

func commonParser(s string, pattern *regexp.Regexp, kind UnitKind) (*UnitQuantity, bool) {
    if match := pattern.FindStringSubmatch(s); len(match) != 0{
        strVal, strPrefix := match[1], match[2]
        intVal, err := strconv.ParseInt(strVal, 10, 64)

        if err != nil {
            return nil, false
        }

        return &UnitQuantity {
            Prefix: UnitPrefixFromString(strPrefix),
            Kind: kind,
            Value: intVal,
        }, true
    }

    return nil, false
}

func ParseString(s string) (*UnitQuantity, bool) {
    s = strings.TrimSpace(s)
    for _, p := range(parsers) {
        if q, ok := p(s); ok {
            return q, true
        }
    }
    return nil, false
}
