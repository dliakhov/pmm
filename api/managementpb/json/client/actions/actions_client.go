// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new actions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for actions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CancelAction(params *CancelActionParams) (*CancelActionOK, error)

	GetAction(params *GetActionParams) (*GetActionOK, error)

	StartMongoDBExplainAction(params *StartMongoDBExplainActionParams) (*StartMongoDBExplainActionOK, error)

	StartMySQLExplainAction(params *StartMySQLExplainActionParams) (*StartMySQLExplainActionOK, error)

	StartMySQLExplainJSONAction(params *StartMySQLExplainJSONActionParams) (*StartMySQLExplainJSONActionOK, error)

	StartMySQLExplainTraditionalJSONAction(params *StartMySQLExplainTraditionalJSONActionParams) (*StartMySQLExplainTraditionalJSONActionOK, error)

	StartMySQLShowCreateTableAction(params *StartMySQLShowCreateTableActionParams) (*StartMySQLShowCreateTableActionOK, error)

	StartMySQLShowIndexAction(params *StartMySQLShowIndexActionParams) (*StartMySQLShowIndexActionOK, error)

	StartMySQLShowTableStatusAction(params *StartMySQLShowTableStatusActionParams) (*StartMySQLShowTableStatusActionOK, error)

<<<<<<< HEAD
	StartPTMySQLSummaryAction(params *StartPTMySQLSummaryActionParams) (*StartPTMySQLSummaryActionOK, error)
=======
	StartPTMongoDBSummaryAction(params *StartPTMongoDBSummaryActionParams) (*StartPTMongoDBSummaryActionOK, error)
>>>>>>> PMM-2.0

	StartPTSummaryAction(params *StartPTSummaryActionParams) (*StartPTSummaryActionOK, error)

	StartPostgreSQLShowCreateTableAction(params *StartPostgreSQLShowCreateTableActionParams) (*StartPostgreSQLShowCreateTableActionOK, error)

	StartPostgreSQLShowIndexAction(params *StartPostgreSQLShowIndexActionParams) (*StartPostgreSQLShowIndexActionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CancelAction cancels action stops an action
*/
func (a *Client) CancelAction(params *CancelActionParams) (*CancelActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CancelAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/Cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CancelActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CancelActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CancelActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAction gets action gets an result of given action
*/
func (a *Client) GetAction(params *GetActionParams) (*GetActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/Get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartMongoDBExplainAction starts mongo DB explain action starts mongo DB e x p l a i n action
*/
func (a *Client) StartMongoDBExplainAction(params *StartMongoDBExplainActionParams) (*StartMongoDBExplainActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMongoDBExplainActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMongoDBExplainAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartMongoDBExplain",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMongoDBExplainActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartMongoDBExplainActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartMongoDBExplainActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartMySQLExplainAction starts my SQL explain action starts my SQL e x p l a i n action with traditional output
*/
func (a *Client) StartMySQLExplainAction(params *StartMySQLExplainActionParams) (*StartMySQLExplainActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMySQLExplainActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMySQLExplainAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartMySQLExplain",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMySQLExplainActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartMySQLExplainActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartMySQLExplainActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartMySQLExplainJSONAction starts my SQL explain JSON action starts my SQL e x p l a i n action with JSON output
*/
func (a *Client) StartMySQLExplainJSONAction(params *StartMySQLExplainJSONActionParams) (*StartMySQLExplainJSONActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMySQLExplainJSONActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMySQLExplainJSONAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartMySQLExplainJSON",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMySQLExplainJSONActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartMySQLExplainJSONActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartMySQLExplainJSONActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartMySQLExplainTraditionalJSONAction starts my SQL explain traditional JSON action starts my SQL e x p l a i n action with traditional JSON output
*/
func (a *Client) StartMySQLExplainTraditionalJSONAction(params *StartMySQLExplainTraditionalJSONActionParams) (*StartMySQLExplainTraditionalJSONActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMySQLExplainTraditionalJSONActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMySQLExplainTraditionalJSONAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartMySQLExplainTraditionalJSON",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMySQLExplainTraditionalJSONActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartMySQLExplainTraditionalJSONActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartMySQLExplainTraditionalJSONActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartMySQLShowCreateTableAction starts my SQL show create table action starts my SQL s h o w c r e a t e t a b l e action
*/
func (a *Client) StartMySQLShowCreateTableAction(params *StartMySQLShowCreateTableActionParams) (*StartMySQLShowCreateTableActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMySQLShowCreateTableActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMySQLShowCreateTableAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartMySQLShowCreateTable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMySQLShowCreateTableActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartMySQLShowCreateTableActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartMySQLShowCreateTableActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartMySQLShowIndexAction starts my SQL show index action starts my SQL s h o w i n d e x action
*/
func (a *Client) StartMySQLShowIndexAction(params *StartMySQLShowIndexActionParams) (*StartMySQLShowIndexActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMySQLShowIndexActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMySQLShowIndexAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartMySQLShowIndex",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMySQLShowIndexActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartMySQLShowIndexActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartMySQLShowIndexActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartMySQLShowTableStatusAction starts my SQL show table status action starts my SQL s h o w t a b l e s t a t u s action
*/
func (a *Client) StartMySQLShowTableStatusAction(params *StartMySQLShowTableStatusActionParams) (*StartMySQLShowTableStatusActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMySQLShowTableStatusActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMySQLShowTableStatusAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartMySQLShowTableStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMySQLShowTableStatusActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartMySQLShowTableStatusActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartMySQLShowTableStatusActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
<<<<<<< HEAD
  StartPTMySQLSummaryAction starts PT my SQL summary action starts pt mysql summary action
*/
func (a *Client) StartPTMySQLSummaryAction(params *StartPTMySQLSummaryActionParams) (*StartPTMySQLSummaryActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartPTMySQLSummaryActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartPTMySQLSummaryAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartPTMySQLSummary",
=======
  StartPTMongoDBSummaryAction starts PT mongo DB summary action starts pt mongodb summary action
*/
func (a *Client) StartPTMongoDBSummaryAction(params *StartPTMongoDBSummaryActionParams) (*StartPTMongoDBSummaryActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartPTMongoDBSummaryActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartPTMongoDBSummaryAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartPTMongoDBSummary",
>>>>>>> PMM-2.0
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
<<<<<<< HEAD
		Reader:             &StartPTMySQLSummaryActionReader{formats: a.formats},
=======
		Reader:             &StartPTMongoDBSummaryActionReader{formats: a.formats},
>>>>>>> PMM-2.0
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	success, ok := result.(*StartPTMySQLSummaryActionOK)
=======
	success, ok := result.(*StartPTMongoDBSummaryActionOK)
>>>>>>> PMM-2.0
	if ok {
		return success, nil
	}
	// unexpected success response
<<<<<<< HEAD
	unexpectedSuccess := result.(*StartPTMySQLSummaryActionDefault)
=======
	unexpectedSuccess := result.(*StartPTMongoDBSummaryActionDefault)
>>>>>>> PMM-2.0
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartPTSummaryAction starts PT summary action starts pt summary action
*/
func (a *Client) StartPTSummaryAction(params *StartPTSummaryActionParams) (*StartPTSummaryActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartPTSummaryActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartPTSummaryAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartPTSummary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartPTSummaryActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartPTSummaryActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartPTSummaryActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartPostgreSQLShowCreateTableAction starts postgre SQL show create table action starts postgre SQL s h o w c r e a t e t a b l e action
*/
func (a *Client) StartPostgreSQLShowCreateTableAction(params *StartPostgreSQLShowCreateTableActionParams) (*StartPostgreSQLShowCreateTableActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartPostgreSQLShowCreateTableActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartPostgreSQLShowCreateTableAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartPostgreSQLShowCreateTable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartPostgreSQLShowCreateTableActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartPostgreSQLShowCreateTableActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartPostgreSQLShowCreateTableActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartPostgreSQLShowIndexAction starts postgre SQL show index action starts postgre SQL s h o w i n d e x action
*/
func (a *Client) StartPostgreSQLShowIndexAction(params *StartPostgreSQLShowIndexActionParams) (*StartPostgreSQLShowIndexActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartPostgreSQLShowIndexActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartPostgreSQLShowIndexAction",
		Method:             "POST",
		PathPattern:        "/v1/management/Actions/StartPostgreSQLShowIndex",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartPostgreSQLShowIndexActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartPostgreSQLShowIndexActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartPostgreSQLShowIndexActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
