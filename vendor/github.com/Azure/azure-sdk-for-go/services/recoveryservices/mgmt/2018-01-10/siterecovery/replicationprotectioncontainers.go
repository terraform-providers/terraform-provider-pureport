package siterecovery

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ReplicationProtectionContainersClient is the client for the ReplicationProtectionContainers methods of the
// Siterecovery service.
type ReplicationProtectionContainersClient struct {
	BaseClient
}

// NewReplicationProtectionContainersClient creates an instance of the ReplicationProtectionContainersClient client.
func NewReplicationProtectionContainersClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationProtectionContainersClient {
	return NewReplicationProtectionContainersClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationProtectionContainersClientWithBaseURI creates an instance of the ReplicationProtectionContainersClient
// client.
func NewReplicationProtectionContainersClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationProtectionContainersClient {
	return ReplicationProtectionContainersClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Create operation to create a protection container.
// Parameters:
// fabricName - unique fabric ARM name.
// protectionContainerName - unique protection container ARM name.
// creationInput - creation input.
func (client ReplicationProtectionContainersClient) Create(ctx context.Context, fabricName string, protectionContainerName string, creationInput CreateProtectionContainerInput) (result ReplicationProtectionContainersCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationProtectionContainersClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, fabricName, protectionContainerName, creationInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ReplicationProtectionContainersClient) CreatePreparer(ctx context.Context, fabricName string, protectionContainerName string, creationInput CreateProtectionContainerInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}", pathParameters),
		autorest.WithJSON(creationInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) CreateSender(req *http.Request) (future ReplicationProtectionContainersCreateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) CreateResponder(resp *http.Response) (result ProtectionContainer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete operation to remove a protection container.
// Parameters:
// fabricName - unique fabric ARM name.
// protectionContainerName - unique protection container ARM name.
func (client ReplicationProtectionContainersClient) Delete(ctx context.Context, fabricName string, protectionContainerName string) (result ReplicationProtectionContainersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationProtectionContainersClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, fabricName, protectionContainerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ReplicationProtectionContainersClient) DeletePreparer(ctx context.Context, fabricName string, protectionContainerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/remove", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) DeleteSender(req *http.Request) (future ReplicationProtectionContainersDeleteFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DiscoverProtectableItem the operation to a add a protectable item to a protection container(Add physical server.)
// Parameters:
// fabricName - the name of the fabric.
// protectionContainerName - the name of the protection container.
// discoverProtectableItemRequest - the request object to add a protectable item.
func (client ReplicationProtectionContainersClient) DiscoverProtectableItem(ctx context.Context, fabricName string, protectionContainerName string, discoverProtectableItemRequest DiscoverProtectableItemRequest) (result ReplicationProtectionContainersDiscoverProtectableItemFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationProtectionContainersClient.DiscoverProtectableItem")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DiscoverProtectableItemPreparer(ctx, fabricName, protectionContainerName, discoverProtectableItemRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "DiscoverProtectableItem", nil, "Failure preparing request")
		return
	}

	result, err = client.DiscoverProtectableItemSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "DiscoverProtectableItem", result.Response(), "Failure sending request")
		return
	}

	return
}

// DiscoverProtectableItemPreparer prepares the DiscoverProtectableItem request.
func (client ReplicationProtectionContainersClient) DiscoverProtectableItemPreparer(ctx context.Context, fabricName string, protectionContainerName string, discoverProtectableItemRequest DiscoverProtectableItemRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/discoverProtectableItem", pathParameters),
		autorest.WithJSON(discoverProtectableItemRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DiscoverProtectableItemSender sends the DiscoverProtectableItem request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) DiscoverProtectableItemSender(req *http.Request) (future ReplicationProtectionContainersDiscoverProtectableItemFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DiscoverProtectableItemResponder handles the response to the DiscoverProtectableItem request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) DiscoverProtectableItemResponder(resp *http.Response) (result ProtectionContainer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the details of a protection container.
// Parameters:
// fabricName - fabric name.
// protectionContainerName - protection container name.
func (client ReplicationProtectionContainersClient) Get(ctx context.Context, fabricName string, protectionContainerName string) (result ProtectionContainer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationProtectionContainersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, fabricName, protectionContainerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationProtectionContainersClient) GetPreparer(ctx context.Context, fabricName string, protectionContainerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) GetResponder(resp *http.Response) (result ProtectionContainer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the protection containers in a vault.
func (client ReplicationProtectionContainersClient) List(ctx context.Context) (result ProtectionContainerCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationProtectionContainersClient.List")
		defer func() {
			sc := -1
			if result.pcc.Response.Response != nil {
				sc = result.pcc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pcc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "List", resp, "Failure sending request")
		return
	}

	result.pcc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationProtectionContainersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) ListResponder(resp *http.Response) (result ProtectionContainerCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ReplicationProtectionContainersClient) listNextResults(ctx context.Context, lastResults ProtectionContainerCollection) (result ProtectionContainerCollection, err error) {
	req, err := lastResults.protectionContainerCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationProtectionContainersClient) ListComplete(ctx context.Context) (result ProtectionContainerCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationProtectionContainersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByReplicationFabrics lists the protection containers in the specified fabric.
// Parameters:
// fabricName - fabric name.
func (client ReplicationProtectionContainersClient) ListByReplicationFabrics(ctx context.Context, fabricName string) (result ProtectionContainerCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationProtectionContainersClient.ListByReplicationFabrics")
		defer func() {
			sc := -1
			if result.pcc.Response.Response != nil {
				sc = result.pcc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByReplicationFabricsNextResults
	req, err := client.ListByReplicationFabricsPreparer(ctx, fabricName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "ListByReplicationFabrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.pcc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "ListByReplicationFabrics", resp, "Failure sending request")
		return
	}

	result.pcc, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "ListByReplicationFabrics", resp, "Failure responding to request")
	}

	return
}

// ListByReplicationFabricsPreparer prepares the ListByReplicationFabrics request.
func (client ReplicationProtectionContainersClient) ListByReplicationFabricsPreparer(ctx context.Context, fabricName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReplicationFabricsSender sends the ListByReplicationFabrics request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) ListByReplicationFabricsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByReplicationFabricsResponder handles the response to the ListByReplicationFabrics request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) ListByReplicationFabricsResponder(resp *http.Response) (result ProtectionContainerCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReplicationFabricsNextResults retrieves the next set of results, if any.
func (client ReplicationProtectionContainersClient) listByReplicationFabricsNextResults(ctx context.Context, lastResults ProtectionContainerCollection) (result ProtectionContainerCollection, err error) {
	req, err := lastResults.protectionContainerCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "listByReplicationFabricsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "listByReplicationFabricsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "listByReplicationFabricsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReplicationFabricsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationProtectionContainersClient) ListByReplicationFabricsComplete(ctx context.Context, fabricName string) (result ProtectionContainerCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationProtectionContainersClient.ListByReplicationFabrics")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByReplicationFabrics(ctx, fabricName)
	return
}

// SwitchProtection operation to switch protection from one container to another or one replication provider to
// another.
// Parameters:
// fabricName - unique fabric name.
// protectionContainerName - protection container name.
// switchInput - switch protection input.
func (client ReplicationProtectionContainersClient) SwitchProtection(ctx context.Context, fabricName string, protectionContainerName string, switchInput SwitchProtectionInput) (result ReplicationProtectionContainersSwitchProtectionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationProtectionContainersClient.SwitchProtection")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SwitchProtectionPreparer(ctx, fabricName, protectionContainerName, switchInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "SwitchProtection", nil, "Failure preparing request")
		return
	}

	result, err = client.SwitchProtectionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainersClient", "SwitchProtection", result.Response(), "Failure sending request")
		return
	}

	return
}

// SwitchProtectionPreparer prepares the SwitchProtection request.
func (client ReplicationProtectionContainersClient) SwitchProtectionPreparer(ctx context.Context, fabricName string, protectionContainerName string, switchInput SwitchProtectionInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/switchprotection", pathParameters),
		autorest.WithJSON(switchInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SwitchProtectionSender sends the SwitchProtection request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainersClient) SwitchProtectionSender(req *http.Request) (future ReplicationProtectionContainersSwitchProtectionFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// SwitchProtectionResponder handles the response to the SwitchProtection request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainersClient) SwitchProtectionResponder(resp *http.Response) (result ProtectionContainer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
