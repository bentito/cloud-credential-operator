package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type ConnectionStatus int

const (
    UNKNOWN_CONNECTIONSTATUS ConnectionStatus = iota
    ATTEMPTED_CONNECTIONSTATUS
    SUCCEEDED_CONNECTIONSTATUS
    BLOCKED_CONNECTIONSTATUS
    FAILED_CONNECTIONSTATUS
    UNKNOWNFUTUREVALUE_CONNECTIONSTATUS
)

func (i ConnectionStatus) String() string {
    return []string{"unknown", "attempted", "succeeded", "blocked", "failed", "unknownFutureValue"}[i]
}
func ParseConnectionStatus(v string) (interface{}, error) {
    result := UNKNOWN_CONNECTIONSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_CONNECTIONSTATUS
        case "attempted":
            result = ATTEMPTED_CONNECTIONSTATUS
        case "succeeded":
            result = SUCCEEDED_CONNECTIONSTATUS
        case "blocked":
            result = BLOCKED_CONNECTIONSTATUS
        case "failed":
            result = FAILED_CONNECTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONNECTIONSTATUS
        default:
            return 0, errors.New("Unknown ConnectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeConnectionStatus(values []ConnectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
