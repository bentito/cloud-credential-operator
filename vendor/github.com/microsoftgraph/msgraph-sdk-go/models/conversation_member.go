package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConversationMember provides operations to manage the collection of agreementAcceptance entities.
type ConversationMember struct {
    Entity
    // The display name of the user.
    displayName *string
    // The roles for that user. This property only contains additional qualifiers when relevant - for example, if the member has owner privileges, the roles property contains owner as one of the values. Similarly, if the member is a guest, the roles property contains guest as one of the values. A basic member should not have any values specified in the roles property.
    roles []string
    // The timestamp denoting how far back a conversation's history is shared with the conversation member. This property is settable only for members of a chat.
    visibleHistoryStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewConversationMember instantiates a new conversationMember and sets the default values.
func NewConversationMember()(*ConversationMember) {
    m := &ConversationMember{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.conversationMember";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateConversationMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConversationMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.aadUserConversationMember":
                        return NewAadUserConversationMember(), nil
                }
            }
        }
    }
    return NewConversationMember(), nil
}
// GetDisplayName gets the displayName property value. The display name of the user.
func (m *ConversationMember) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConversationMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["roles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoles)
    res["visibleHistoryStartDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetVisibleHistoryStartDateTime)
    return res
}
// GetRoles gets the roles property value. The roles for that user. This property only contains additional qualifiers when relevant - for example, if the member has owner privileges, the roles property contains owner as one of the values. Similarly, if the member is a guest, the roles property contains guest as one of the values. A basic member should not have any values specified in the roles property.
func (m *ConversationMember) GetRoles()([]string) {
    return m.roles
}
// GetVisibleHistoryStartDateTime gets the visibleHistoryStartDateTime property value. The timestamp denoting how far back a conversation's history is shared with the conversation member. This property is settable only for members of a chat.
func (m *ConversationMember) GetVisibleHistoryStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.visibleHistoryStartDateTime
}
// Serialize serializes information the current object
func (m *ConversationMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        err = writer.WriteCollectionOfStringValues("roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("visibleHistoryStartDateTime", m.GetVisibleHistoryStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the user.
func (m *ConversationMember) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetRoles sets the roles property value. The roles for that user. This property only contains additional qualifiers when relevant - for example, if the member has owner privileges, the roles property contains owner as one of the values. Similarly, if the member is a guest, the roles property contains guest as one of the values. A basic member should not have any values specified in the roles property.
func (m *ConversationMember) SetRoles(value []string)() {
    m.roles = value
}
// SetVisibleHistoryStartDateTime sets the visibleHistoryStartDateTime property value. The timestamp denoting how far back a conversation's history is shared with the conversation member. This property is settable only for members of a chat.
func (m *ConversationMember) SetVisibleHistoryStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.visibleHistoryStartDateTime = value
}
