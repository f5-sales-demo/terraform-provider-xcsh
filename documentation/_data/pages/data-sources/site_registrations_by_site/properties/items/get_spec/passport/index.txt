---
page_title: "items.get_spec.passport"
subcategory: ""
description: "Passport stores information about identification and node configuration provided by CE during registration. It can be manually updated by user during approval."
xcsh_docs: {"aliases": ["items get spec passport"], "body_bytes": 4468, "body_sha256": "sha256:f0941886db6c5c495cabfe1e1009e78641f5a87fb1e9d77bb2a841c9c948981e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport:default_os_version", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport:default_sw_version"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec", "path": "documentation/data-sources/site_registrations_by_site/properties/items/get_spec/passport/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0303021112311131-2322121310302311-3131313113230232-0333130233220221-3211303321110212-2131312010312201-2112133312200130-2122332302103302", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "passport"], "schema_version": 1, "sections": [{"aliases": ["items get spec passport cluster name"], "anchor": "schema-items--get_spec--passport--cluster_name", "description": "Cluster Name. Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "passport", "cluster_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec passport cluster size"], "anchor": "schema-items--get_spec--passport--cluster_size", "description": "Defines how many master nodes is in the cluster, only 1 or 3 is allowed 1 - cluster have single master, without HA 3 - cluster have 3 masters, with HA, all nodes should be allowed at same time, cluster won't start until ALL nodes are ADMITTED 0 - same as 1 This value can't be changed after..", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "passport", "cluster_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["items get spec passport cluster type"], "anchor": "schema-items--get_spec--passport--cluster_type", "description": "Cluster Type. Cluster or grouping configuration", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "passport", "cluster_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec passport default os version"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport:default_os_version", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "passport", "default_os_version"], "syntax": "attribute", "type": "object"}, {"aliases": ["items get spec passport default sw version"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport:default_sw_version", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "passport", "default_sw_version"], "syntax": "attribute", "type": "object"}, {"aliases": ["items get spec passport latitude"], "anchor": "schema-items--get_spec--passport--latitude", "description": "Latitude. Geographic location of this site.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "passport", "latitude"], "syntax": "attribute", "type": "number"}, {"aliases": ["items get spec passport longitude"], "anchor": "schema-items--get_spec--passport--longitude", "description": "Longitude. Geographic location of this site.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "passport", "longitude"], "syntax": "attribute", "type": "number"}, {"aliases": ["items get spec passport operating system version"], "anchor": "schema-items--get_spec--passport--operating_system_version", "description": "Exclusive with Operating System Version is optional parameter, which allows to specify target SW version for particular site e.g. 7.2009.10.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport", "enum_extraction_complete": true, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "passport", "operating_system_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec passport private network name"], "anchor": "schema-items--get_spec--passport--private_network_name", "description": "Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink, CloudLink and L3VPN.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport", "enum_extraction_complete": true, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "passport", "private_network_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec passport volterra software version"], "anchor": "schema-items--get_spec--passport--volterra_software_version", "description": "Exclusive with F5XC Software Version is optional parameter, which allows to specify target SW version for particular site e.g. Crt-20210329-1002.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:passport", "enum_extraction_complete": true, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "passport", "volterra_software_version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/passport/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Passport stores information about identification and node configuration provided by CE during registration. It can be manually updated by user during approval.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.passport

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- items.get_spec.passport

<a id="section"></a>

Type: `"single"`. Computed.

Passport stores information about identification and node configuration provided by CE during
registration. It can be manually updated by user during approval.

## Direct properties

<a id="schema-items--get_spec--passport--cluster_name"></a>

### cluster_name property

Type: `"string"`. Computed.

Cluster Name. Human-readable name for the resource

<a id="schema-items--get_spec--passport--cluster_size"></a>

### cluster_size property

Type: `"number"`. Computed.

Defines how many master nodes is in the cluster, only 1 or 3 is allowed 1 - cluster have single
master, without HA 3 - cluster have 3 masters, with HA, all nodes should be allowed at same time,
cluster won't start until ALL nodes are ADMITTED 0 - same as 1 This value can't be changed after..

<a id="schema-items--get_spec--passport--cluster_type"></a>

### cluster_type property

Type: `"string"`. Computed.

Cluster Type. Cluster or grouping configuration

- [default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/passport/default_os_version/): complete subsection reference.

- [default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/passport/default_sw_version/): complete subsection reference.

<a id="schema-items--get_spec--passport--latitude"></a>

### latitude property

Type: `"number"`. Computed.

Latitude. Geographic location of this site.

<a id="schema-items--get_spec--passport--longitude"></a>

### longitude property

Type: `"number"`. Computed.

Longitude. Geographic location of this site.

<a id="schema-items--get_spec--passport--operating_system_version"></a>

### operating_system_version property

Type: `"string"`. Computed.

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

<a id="schema-items--get_spec--passport--private_network_name"></a>

### private_network_name property

Type: `"string"`. Computed.

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="schema-items--get_spec--passport--volterra_software_version"></a>

### volterra_software_version property

Type: `"string"`. Computed.

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

## Next pages

- [items.get_spec.passport.default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/passport/default_os_version/)
- [items.get_spec.passport.default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/passport/default_sw_version/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
