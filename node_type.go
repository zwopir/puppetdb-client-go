package puppetdb

/*
Data structure representative of the return wire format from nodes query
end-points.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/query/v3/nodes.html#get-v3nodes
*/
type Node struct {
	Name             string `json:"name"`
	Deactivated      string `json:"deactivated"`
	CatalogTimestamp string `json:"catalog_timestamp"`
	FactsTimestamp   string `json:"facts_timestamp"`
	ReportTimestamp  string `json:"report_timestamp"`
}
