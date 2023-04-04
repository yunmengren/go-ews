package goews

import (
	"github.com/yunmengren/go-ews/elements"
	"github.com/yunmengren/go-ews/operations"
)

// Find information about the FindItem EWS operation.
// The FindItem operation searches for items that are located in a user's mailbox. This operation provides many ways to filter and format how search results are returned to the caller.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/finditem-operation
func (c *client) GetItem(eItem *elements.GetItem) (*elements.GetItemResponse, error) {
	oem, err := operations.NewEnvelopeMarshal(eItem)
	if err != nil {
		return nil, err
	}
	eResponse := &elements.GetItemResponse{}
	err = c.DoRequest(oem, eResponse)
	if err != nil {
		return nil, err
	}
	return eResponse, nil
}
