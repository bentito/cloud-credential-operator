package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type AccessPackageAssignmentState int

const (
    DELIVERING_ACCESSPACKAGEASSIGNMENTSTATE AccessPackageAssignmentState = iota
    PARTIALLYDELIVERED_ACCESSPACKAGEASSIGNMENTSTATE
    DELIVERED_ACCESSPACKAGEASSIGNMENTSTATE
    EXPIRED_ACCESSPACKAGEASSIGNMENTSTATE
    DELIVERYFAILED_ACCESSPACKAGEASSIGNMENTSTATE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGEASSIGNMENTSTATE
)

func (i AccessPackageAssignmentState) String() string {
    return []string{"delivering", "partiallyDelivered", "delivered", "expired", "deliveryFailed", "unknownFutureValue"}[i]
}
func ParseAccessPackageAssignmentState(v string) (interface{}, error) {
    result := DELIVERING_ACCESSPACKAGEASSIGNMENTSTATE
    switch v {
        case "delivering":
            result = DELIVERING_ACCESSPACKAGEASSIGNMENTSTATE
        case "partiallyDelivered":
            result = PARTIALLYDELIVERED_ACCESSPACKAGEASSIGNMENTSTATE
        case "delivered":
            result = DELIVERED_ACCESSPACKAGEASSIGNMENTSTATE
        case "expired":
            result = EXPIRED_ACCESSPACKAGEASSIGNMENTSTATE
        case "deliveryFailed":
            result = DELIVERYFAILED_ACCESSPACKAGEASSIGNMENTSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGEASSIGNMENTSTATE
        default:
            return 0, errors.New("Unknown AccessPackageAssignmentState value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageAssignmentState(values []AccessPackageAssignmentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
