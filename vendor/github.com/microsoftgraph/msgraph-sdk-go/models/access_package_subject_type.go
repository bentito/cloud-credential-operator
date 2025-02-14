package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type AccessPackageSubjectType int

const (
    NOTSPECIFIED_ACCESSPACKAGESUBJECTTYPE AccessPackageSubjectType = iota
    USER_ACCESSPACKAGESUBJECTTYPE
    SERVICEPRINCIPAL_ACCESSPACKAGESUBJECTTYPE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGESUBJECTTYPE
)

func (i AccessPackageSubjectType) String() string {
    return []string{"notSpecified", "user", "servicePrincipal", "unknownFutureValue"}[i]
}
func ParseAccessPackageSubjectType(v string) (interface{}, error) {
    result := NOTSPECIFIED_ACCESSPACKAGESUBJECTTYPE
    switch v {
        case "notSpecified":
            result = NOTSPECIFIED_ACCESSPACKAGESUBJECTTYPE
        case "user":
            result = USER_ACCESSPACKAGESUBJECTTYPE
        case "servicePrincipal":
            result = SERVICEPRINCIPAL_ACCESSPACKAGESUBJECTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGESUBJECTTYPE
        default:
            return 0, errors.New("Unknown AccessPackageSubjectType value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageSubjectType(values []AccessPackageSubjectType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
