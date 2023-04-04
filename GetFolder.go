package goews

import (
	"github.com/yunmengren/go-ews/elements"
	"github.com/yunmengren/go-ews/operations"
)

// The GetFolder operation gets folders from the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfolder-operation
func (c *client) GetFolder(eItem *elements.GetFolder) (*elements.GetFolderResponse, error) {
	oem, err := operations.NewEnvelopeMarshal(eItem)
	if err != nil {
		return nil, err
	}
	eResponse := &elements.GetFolderResponse{}
	err = c.DoRequest(oem, eResponse)
	if err != nil {
		return nil, err
	}
	return eResponse, nil
}
