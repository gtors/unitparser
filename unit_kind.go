package unitparser

type UnitKind  uint8

const (
    UNIT_AMPERE UnitKind = iota
    UNIT_FARAD
    UNIT_HENRY
    UNIT_HERTZ
    UNIT_OHM
    UNIT_VOLT
)

func (uk UnitKind) String() string {
    switch(uk) {
    case UNIT_AMPERE:
        return "A"
    case UNIT_FARAD:
        return "F"
    case UNIT_HENRY:
        return "H"
    case UNIT_HERTZ:
        return "Hz"
    case UNIT_OHM:
        return "Ω"
    case UNIT_VOLT:
        return "V"
    }
    return ""
}
