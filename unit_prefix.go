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

func UnitPrefixFromString(s string) UnitPrefix {
    switch (s) {
    case "Г", "G":
        return PREFIX_GIGA
    case "З", "Z":
        return PREFIX_ZETTA
    case "И", "Y":
        return PREFIX_YOTTA
    case "М", "M":
        return PREFIX_MEGA
    case "П", "P":
        return PREFIX_PETA
    case "Т", "T":
        return PREFIX_TERA
    case "Э", "E":
        return PREFIX_EXA
    case "а", "a":
        return PREFIX_ATTO
    case "г", "h":
        return PREFIX_HECTO
    case "д", "d":
        return PREFIX_DECI
    case "да", "da":
        return PREFIX_DEKA
    case "з", "z":
        return PREFIX_ZEPTO
    case "и", "y":
        return PREFIX_YOCTO
    case "к", "k":
        return PREFIX_KILO
    case "м", "m":
        return PREFIX_MILLI
    case "мк", "μ":
        return PREFIX_MICRO
    case "н", "n":
        return PREFIX_NANO
    case "п", "p":
        return PREFIX_PICO
    case "с", "c":
        return PREFIX_CENTI
    case "ф", "f":
        return PREFIX_FEMTO
    default:
        return PREFIX_NONE
    }
}

func (up UnitPrefix) String() string {
    switch (up) {
    case PREFIX_YOTTA:
        return "Y"
    case PREFIX_ZETTA:
        return "Z"
    case PREFIX_EXA:
        return "E"
    case PREFIX_PETA:
        return "P"
    case PREFIX_TERA:
        return "T"
    case PREFIX_GIGA:
        return "G"
    case PREFIX_MEGA:
        return "M"
    case PREFIX_KILO:
        return "k"
    case PREFIX_HECTO:
        return "h"
    case PREFIX_DEKA:
        return "da"
    case PREFIX_DECI:
        return "d"
    case PREFIX_CENTI:
        return "c"
    case PREFIX_MILLI:
        return "m"
    case PREFIX_MICRO:
        return "μ"
    case PREFIX_NANO:
        return "n"
    case PREFIX_PICO:
        return "p"
    case PREFIX_FEMTO:
        return "f"
    case PREFIX_ATTO:
        return "a"
    case PREFIX_ZEPTO:
        return "z"
    case PREFIX_YOCTO:
        return "y"
    }
    return ""
}
