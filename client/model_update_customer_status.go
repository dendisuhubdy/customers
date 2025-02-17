/*
 * Customers API
 *
 * Customers ...
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateCustomerStatus struct for UpdateCustomerStatus
type UpdateCustomerStatus struct {
	// Free form comment about the customer status update
	Comment string `json:"comment,omitempty"`
	// State of the customer
	Status string `json:"status"`
}
