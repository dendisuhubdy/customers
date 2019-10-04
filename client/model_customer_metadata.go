/*
 * Customers API
 *
 * Customers ...
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CustomerMetadata struct for CustomerMetadata
type CustomerMetadata struct {
	// Map of unique keys associated to values to act as foreign key relationships or arbitrary data associated to a Customer.
	Metadata map[string]string `json:"metadata"`
}
