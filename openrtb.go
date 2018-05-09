package rtb

type OpenRTB struct {
	Version       string      `json:"ver"`        // ver string Version of the Layer-3 OpenRTB specification (e.g., “3.0”).
	DomainSpec    string      `json:"domainspec"` // domainspec string; recommended Identifier of the Layer-4 domain model contained within “domain” objects in the Advertising Common Object Model, “AdCOM” (e.g., “1.0”).
	DomainVersion string      `json:"domainver"`  // domainver string; recommended Specification version of the Layer-4 domain model referenced in the “domainspec” attribute.
	Request       Request     `json:"request"`    // request object; required * Bid request container. * Required only for request payloads.
	Response      unk         `json:"response"`   // response object; required * Bid response container. * Required only for response payloads.
	Ext           interface{} `json:"ext"`        // ext object Optional exchange or demand source specific extensions.
}
