/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VaultCreate struct {
	Name string `json:"name"`
	Size float32 `json:"size"`
	Service string `json:"service,omitempty"`
	Credential VaultCreateCredential `json:"credential,omitempty"`
	Snapshot string `json:"snapshot,omitempty"`
	Tag map[string]interface{} `json:"tag,omitempty"`
}
