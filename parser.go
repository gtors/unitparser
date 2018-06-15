package unitparser

import (
    "regexp"
    "strconv"
    "strings"
)

var (
    parsers = make([]Parser, 0, 100)[:0]
)

type Parser func(s string) (*UnitQuantity, bool)

func init() {

    common := `^(-?\d+)\s*(` + strings.Join(UnitPrefixAllStringVariants(), "|") + `)?`

    for _, kind := range UnitKindValues() {
        kindSubpattern := `(?:` + strings.Join(kind.StringVariants(), "|") + `)$`
        pattern := regexp.MustCompile(common + kindSubpattern)
        _kind := kind

        parsers = append(parsers, func (s string) (*UnitQuantity, bool) {
            return commonParser(s, pattern, _kind)
        })
    }
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
