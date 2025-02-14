package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type ConditionalAccessGrantControl int

const (
    BLOCK_CONDITIONALACCESSGRANTCONTROL ConditionalAccessGrantControl = iota
    MFA_CONDITIONALACCESSGRANTCONTROL
    COMPLIANTDEVICE_CONDITIONALACCESSGRANTCONTROL
    DOMAINJOINEDDEVICE_CONDITIONALACCESSGRANTCONTROL
    APPROVEDAPPLICATION_CONDITIONALACCESSGRANTCONTROL
    COMPLIANTAPPLICATION_CONDITIONALACCESSGRANTCONTROL
    PASSWORDCHANGE_CONDITIONALACCESSGRANTCONTROL
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSGRANTCONTROL
)

func (i ConditionalAccessGrantControl) String() string {
    return []string{"block", "mfa", "compliantDevice", "domainJoinedDevice", "approvedApplication", "compliantApplication", "passwordChange", "unknownFutureValue"}[i]
}
func ParseConditionalAccessGrantControl(v string) (interface{}, error) {
    result := BLOCK_CONDITIONALACCESSGRANTCONTROL
    switch v {
        case "block":
            result = BLOCK_CONDITIONALACCESSGRANTCONTROL
        case "mfa":
            result = MFA_CONDITIONALACCESSGRANTCONTROL
        case "compliantDevice":
            result = COMPLIANTDEVICE_CONDITIONALACCESSGRANTCONTROL
        case "domainJoinedDevice":
            result = DOMAINJOINEDDEVICE_CONDITIONALACCESSGRANTCONTROL
        case "approvedApplication":
            result = APPROVEDAPPLICATION_CONDITIONALACCESSGRANTCONTROL
        case "compliantApplication":
            result = COMPLIANTAPPLICATION_CONDITIONALACCESSGRANTCONTROL
        case "passwordChange":
            result = PASSWORDCHANGE_CONDITIONALACCESSGRANTCONTROL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONDITIONALACCESSGRANTCONTROL
        default:
            return 0, errors.New("Unknown ConditionalAccessGrantControl value: " + v)
    }
    return &result, nil
}
func SerializeConditionalAccessGrantControl(values []ConditionalAccessGrantControl) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
