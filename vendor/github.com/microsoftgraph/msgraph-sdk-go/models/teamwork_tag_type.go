package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type TeamworkTagType int

const (
    STANDARD_TEAMWORKTAGTYPE TeamworkTagType = iota
    UNKNOWNFUTUREVALUE_TEAMWORKTAGTYPE
)

func (i TeamworkTagType) String() string {
    return []string{"standard", "unknownFutureValue"}[i]
}
func ParseTeamworkTagType(v string) (interface{}, error) {
    result := STANDARD_TEAMWORKTAGTYPE
    switch v {
        case "standard":
            result = STANDARD_TEAMWORKTAGTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMWORKTAGTYPE
        default:
            return 0, errors.New("Unknown TeamworkTagType value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkTagType(values []TeamworkTagType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
