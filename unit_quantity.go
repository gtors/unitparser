package unitparser

import (
    "fmt"
)

type UnitQuantity struct {
    Prefix UnitPrefix
    Kind UnitKind
    Value int64
}

func (u *UnitQuantity) String() string {
    return fmt.Sprintf("%d %s%s", u.Value, u.Prefix, u.Kind)
}
