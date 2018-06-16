package unitparser

type UnitPrefix uint8

const (
    PREFIX_NONE UnitPrefix =  iota
    PREFIX_YOTTA
    PREFIX_ZETTA
    PREFIX_EXA
    PREFIX_PETA
    PREFIX_TERA
    PREFIX_GIGA
    PREFIX_MEGA
    PREFIX_KILO
    PREFIX_HECTO
    PREFIX_DEKA
    PREFIX_DECI
    PREFIX_CENTI
    PREFIX_MILLI
    PREFIX_MICRO
    PREFIX_NANO
    PREFIX_PICO
    PREFIX_FEMTO
    PREFIX_ATTO
    PREFIX_ZEPTO
    PREFIX_YOCTO
)

var (
    unitPrefixVariants = map[UnitPrefix][]string {
        PREFIX_ATTO: []string{"a", "а"},
        PREFIX_CENTI: []string{"c", "с"},
        PREFIX_DECI: []string{"d", "д"},
        PREFIX_DEKA: []string{"da", "да"},
        PREFIX_EXA: []string{"E", "Э"},
        PREFIX_FEMTO: []string{"f", "ф"},
        PREFIX_GIGA: []string{"G", "Г"},
        PREFIX_HECTO: []string{"h", "г"},
        PREFIX_KILO: []string{"k", "к", "K"},
        PREFIX_MEGA: []string{"M", "М"},
        PREFIX_MICRO: []string{"μ", "мк",},
        PREFIX_MILLI: []string{"m", "м"},
        PREFIX_NANO: []string{"n", "н"},
        PREFIX_PETA: []string{"P", "П"},
        PREFIX_PICO: []string{"p", "п"},
        PREFIX_TERA: []string{"T", "Т"},
        PREFIX_YOCTO: []string{"y", "и"},
        PREFIX_YOTTA: []string{"Y", "И"},
        PREFIX_ZEPTO: []string{"z", "з"},
        PREFIX_ZETTA: []string{"Z", "З"},
    }
)

func UnitPrefixValues() []UnitPrefix {
    keys := make([]UnitPrefix, 0, len(unitPrefixVariants))
    for k := range unitPrefixVariants {
        keys = append(keys, k)
    }
    return keys
}

func UnitPrefixAllStringVariants() []string {
    allVariants := make([]string, 0, len(unitPrefixVariants) * 2 + 1)[:0]
    for _, variants := range unitPrefixVariants {
        allVariants = append(allVariants, variants...)
    }
    return allVariants
}

func UnitPrefixFromString(s string) UnitPrefix {
    for prefix, variants := range unitPrefixVariants {
        for _, variant := range variants {
            if variant == s {
                return prefix
            }
        }
    }
    return PREFIX_NONE
}

func (up UnitPrefix) String() string {
    return up.StringVariants()[0]
}

func (up UnitPrefix) StringVariants() []string {
    if variants, ok := unitPrefixVariants[up]; ok {
        return variants
    } else {
        return []string{""}
    }
}

func (up UnitPrefix) GetMultiplier() float64 {
    switch {
    case PREFIX_ATTO:
        return 1E-18
    case PREFIX_CENTI:
        return 1E-2
    case PREFIX_DECI:
        return 1E-1
    case PREFIX_DEKA:
        return 10.0
    case PREFIX_EXA:
        return 1E+18
    case PREFIX_FEMTO:
        return 1E-15
    case PREFIX_GIGA:
        return 1E+9
    case PREFIX_HECTO:
        return 1E+2
    case PREFIX_KILO:
        return 1E+3
    case PREFIX_MEGA:
        return 1E+6
    case PREFIX_MICRO:
        return 1E-6
    case PREFIX_MILLI:
        return 1E-3
    case PREFIX_NANO:
        return 1E-9
    case PREFIX_PETA:
        return 1E+15
    case PREFIX_PICO:
        return 1E-12
    case PREFIX_TERA:
        return 1E+12
    case PREFIX_YOCTO:
        return 1E-24
    case PREFIX_YOTTA:
        return 1E+24
    case PREFIX_ZEPTO:
        return 1E-21
    case PREFIX_ZETTA:
        return 1E+21
    default:
        return 1
    }
}
