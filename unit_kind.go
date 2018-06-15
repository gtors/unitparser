package unitparser

type UnitKind  uint8

const (
    UNIT_NONE UnitKind = iota
    UNIT_AMPERE
    UNIT_BECQUEREL
    UNIT_CANDELA
    UNIT_CELSIUS
    UNIT_COULOMB
    UNIT_FARAD
    UNIT_GRAY
    UNIT_HENRY
    UNIT_HERTZ
    UNIT_JOULE
    UNIT_KELVIN
    UNIT_LUMEN
    UNIT_LUX
    UNIT_METRE
    UNIT_MOLE
    UNIT_NEWTON
    UNIT_OHM
    UNIT_PASCAL
    UNIT_RADIAN
    UNIT_SECOND
    UNIT_SIEMENS
    UNIT_SIEVERT
    UNIT_STERADIAN
    UNIT_TESLA
    UNIT_VOLT
    UNIT_WATT
    UNIT_WEBBER
)

var (
    unitKindVariants= map[UnitKind][]string {
        UNIT_AMPERE: []string{"A", "А"},
        UNIT_BECQUEREL: []string{"Bq", "Бк"},
        UNIT_CANDELA: []string{"cd","кд"},
        UNIT_CELSIUS: []string{"°C"},
        UNIT_COULOMB: []string{"C", "Кл"},
        UNIT_FARAD: []string{"F", "Ф"},
        UNIT_GRAY: []string{"Gy", "Гр"},
        UNIT_HENRY: []string{"H", "Гн"},
        UNIT_HERTZ: []string{"Hz", "Гц"},
        UNIT_JOULE: []string{"J", "Дж"},
        UNIT_KELVIN: []string{"K", "К"},
        UNIT_LUMEN: []string{"lm", "лм"},
        UNIT_LUX: []string{"lx", "лк"},
        UNIT_METRE: []string{"m","м"},
        UNIT_MOLE: []string{"mol", "моль"},
        UNIT_NEWTON: []string{"N", "Н"},
        UNIT_OHM: []string{"Ω", "Ohm", "Ом"},
        UNIT_PASCAL: []string{"Pa", "Па"},
        UNIT_RADIAN: []string{"rad", "рад"},
        UNIT_SECOND: []string{"s", "с"},
        UNIT_SIEMENS: []string{"S", "См"},
        UNIT_SIEVERT: []string{"Sv", "Зв"},
        UNIT_STERADIAN: []string{"sr", "ср"},
        UNIT_TESLA: []string{"T", "Т"},
        UNIT_VOLT: []string{"V", "В"},
        UNIT_WATT: []string{"W", "Вт"},
        UNIT_WEBBER: []string{"Wb", "Вб"},
    }
)

func UnitKindValues() []UnitKind {
    keys := make([]UnitKind, 0, len(unitKindVariants))
    for k := range unitKindVariants {
        keys = append(keys, k)
    }
    return keys
}

func UnitKindFromString(s string) UnitKind {
    for kind, variants := range unitKindVariants {
        for _, variant := range variants {
            if variant == s {
                return kind
            }
        }
    }
    return UNIT_NONE
}

func (uk UnitKind) String() string {
    return uk.StringVariants()[0]
}

func (uk UnitKind) StringVariants() []string {
    if variants, ok := unitKindVariants[uk]; ok {
        return variants
    } else {
        return []string{""}
    }
}
