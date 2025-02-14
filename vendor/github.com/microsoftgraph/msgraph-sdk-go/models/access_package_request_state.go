package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type AccessPackageRequestState int

const (
    SUBMITTED_ACCESSPACKAGEREQUESTSTATE AccessPackageRequestState = iota
    PENDINGAPPROVAL_ACCESSPACKAGEREQUESTSTATE
    DELIVERING_ACCESSPACKAGEREQUESTSTATE
    DELIVERED_ACCESSPACKAGEREQUESTSTATE
    DELIVERYFAILED_ACCESSPACKAGEREQUESTSTATE
    DENIED_ACCESSPACKAGEREQUESTSTATE
    SCHEDULED_ACCESSPACKAGEREQUESTSTATE
    CANCELED_ACCESSPACKAGEREQUESTSTATE
    PARTIALLYDELIVERED_ACCESSPACKAGEREQUESTSTATE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGEREQUESTSTATE
)

func (i AccessPackageRequestState) String() string {
    return []string{"submitted", "pendingApproval", "delivering", "delivered", "deliveryFailed", "denied", "scheduled", "canceled", "partiallyDelivered", "unknownFutureValue"}[i]
}
func ParseAccessPackageRequestState(v string) (interface{}, error) {
    result := SUBMITTED_ACCESSPACKAGEREQUESTSTATE
    switch v {
        case "submitted":
            result = SUBMITTED_ACCESSPACKAGEREQUESTSTATE
        case "pendingApproval":
            result = PENDINGAPPROVAL_ACCESSPACKAGEREQUESTSTATE
        case "delivering":
            result = DELIVERING_ACCESSPACKAGEREQUESTSTATE
        case "delivered":
            result = DELIVERED_ACCESSPACKAGEREQUESTSTATE
        case "deliveryFailed":
            result = DELIVERYFAILED_ACCESSPACKAGEREQUESTSTATE
        case "denied":
            result = DENIED_ACCESSPACKAGEREQUESTSTATE
        case "scheduled":
            result = SCHEDULED_ACCESSPACKAGEREQUESTSTATE
        case "canceled":
            result = CANCELED_ACCESSPACKAGEREQUESTSTATE
        case "partiallyDelivered":
            result = PARTIALLYDELIVERED_ACCESSPACKAGEREQUESTSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGEREQUESTSTATE
        default:
            return 0, errors.New("Unknown AccessPackageRequestState value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageRequestState(values []AccessPackageRequestState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
