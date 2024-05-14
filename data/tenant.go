package data

import "time"

type Tenant struct {
	ETag                     string    `json:"ETag"`
	PartitionKey             string    `json:"PartitionKey"`
	RowKey                   string    `json:"RowKey"`
	Timestamp                time.Time `json:"Timestamp"`
	ExcludeDate              string    `json:"ExcludeDate"`
	ExcludeUser              string    `json:"ExcludeUser"`
	Excluded                 bool      `json:"Excluded"`
	GraphErrorCount          int       `json:"GraphErrorCount"`
	LastGraphError           string    `json:"LastGraphError"`
	LastRefresh              time.Time `json:"LastRefresh"`
	RequiresRefresh          bool      `json:"RequiresRefresh"`
	CustomerId               string    `json:"customerId"`
	DefaultDomainName        string    `json:"defaultDomainName"`
	DelegatedPrivilegeStatus string    `json:"delegatedPrivilegeStatus"`
	DisplayName              string    `json:"displayName"`
	Domains                  string    `json:"domains"`
	HasAutoExtend            bool      `json:"hasAutoExtend"`
	InitialDomainName        string    `json:"initialDomainName"`
	RelationshipCount        int       `json:"relationshipCount"`
	RelationshipEnd          time.Time `json:"relationshipEnd"`
}
