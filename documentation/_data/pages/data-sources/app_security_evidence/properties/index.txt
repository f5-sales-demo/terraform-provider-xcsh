---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_app_security_evidence."
xcsh_docs: {"aliases": ["app security evidence"], "body_bytes": 6949, "body_sha256": "sha256:99c6851291bec63276ad01b51bdccbf23b8617a934f9ca7b847d2f7516d0a26a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_security_evidence:properties:aggs", "xcsh-docs:data-sources:app_security_evidence:properties:last_sort_values", "xcsh-docs:data-sources:app_security_evidence:properties:sort_values"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_security_evidence:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_security_evidence:reference", "parent_id": "xcsh-docs:data-sources:app_security_evidence:fundamentals", "path": "documentation/data-sources/app_security_evidence/properties/index.md", "product": "distributed-cloud", "provider_name": "app_security_evidence", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0112003311231232-2332331022230311-2301110100313003-2332031131201310-1301022100202113-0011223203211232-1031003322010031-0310200130000020", "registry_path": "docs/guides/data-sources--app_security_evidence--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["aggs"], "anchor": "section", "description": "Aggregations provide summary/analytics data over the security evidence response. If the number of security evidence that matched the query is large and cannot be returned in a single response message, user can GET helpful insights/summary using aggregations. The aggregations are key'ed by..", "document_id": "xcsh-docs:data-sources:app_security_evidence:properties:aggs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aggs"], "syntax": "attribute", "type": "object"}, {"aliases": ["end time"], "anchor": "schema-end_time", "description": "Fetch security evidence whose timestamp <= end_time format: unix_timestamp|RFC 3339 Optional: If not specified, then the end_time will be evaluated to start_time+10m If start_time is not specified, then the end_time will be evaluated to <current time>.", "document_id": "xcsh-docs:data-sources:app_security_evidence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["end_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["evidences"], "anchor": "schema-evidences", "description": "List of security evidences that matched the query. Contains no more than 500 messages.", "document_id": "xcsh-docs:data-sources:app_security_evidence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["evidences"], "syntax": "attribute", "type": "list"}, {"aliases": ["last sort values"], "anchor": "section", "description": "These are timestamp and doc_id values returned by elastic search in the search request. Client is expected to set these values in a subsequent request to GET the next page of results.", "document_id": "xcsh-docs:data-sources:app_security_evidence:properties:last_sort_values", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["last_sort_values"], "syntax": "attribute", "type": "object"}, {"aliases": ["limit"], "anchor": "schema-limit", "description": "Limits the number of security evidence returned in the response Optional: If not specified, first or last 500 security evidence that matches the query (depending on the sort order) will be returned in the response. The maximum value for limit is 500.", "document_id": "xcsh-docs:data-sources:app_security_evidence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limit"], "syntax": "attribute", "type": "number"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace fetch security evidence for a given namespace.", "document_id": "xcsh-docs:data-sources:app_security_evidence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["query"], "anchor": "schema-query", "description": "Query is used to specify the list of matchers syntax for query := {} <matcher> := <field_name><operator>'<value>' <field_name> := string One or more of these fields in the security evidence may be specified in the query. Domain - domain endpoint - endpoint evidence_id - evidence ID..", "document_id": "xcsh-docs:data-sources:app_security_evidence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["query"], "syntax": "attribute", "type": "string"}, {"aliases": ["search after"], "anchor": "schema-search_after", "description": "Search After is used to retrieve large number of log messages (or all log messages) that matches the query. If search_after is set to true, the sort_values in the response can be used in the API to fetch the next batch of logs. The number of messages in each batch is determined by the limit field.", "document_id": "xcsh-docs:data-sources:app_security_evidence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["search_after"], "syntax": "attribute", "type": "bool"}, {"aliases": ["sort"], "anchor": "schema-sort", "description": "Sort algorithm Sort in descending order Sort in ascending order. Possible values are `DESCENDING`, `ASCENDING`. Defaults to `DESCENDING`.", "document_id": "xcsh-docs:data-sources:app_security_evidence:reference", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ASCENDING", "DESCENDING"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sort"], "syntax": "attribute", "type": "string"}, {"aliases": ["sort by"], "anchor": "schema-sort_by", "description": "Optional: default is sort by last_event_time.", "document_id": "xcsh-docs:data-sources:app_security_evidence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sort_by"], "syntax": "attribute", "type": "string"}, {"aliases": ["sort values"], "anchor": "section", "description": "These are timestamp and doc_id values returned by elastic search in the search request. Client is expected to set these values in a subsequent request to GET the next page of results.", "document_id": "xcsh-docs:data-sources:app_security_evidence:properties:sort_values", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sort_values"], "syntax": "attribute", "type": "object"}, {"aliases": ["start time"], "anchor": "schema-start_time", "description": "Fetch security evidence whose timestamp >= start_time format: unix_timestamp|RFC 3339 Optional: If not specified, then the start_time will be evaluated to end_time-10m If end_time is not specified, then the start_time will be evaluated to <current time>-10m.", "document_id": "xcsh-docs:data-sources:app_security_evidence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["start_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["total hits"], "anchor": "schema-total_hits", "description": "Total number of security events that matched the query.", "document_id": "xcsh-docs:data-sources:app_security_evidence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["total_hits"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_security_evidence/properties/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Property reference for xcsh_app_security_evidence.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_app_security_evidence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/)
- Property reference

## Direct properties

- [aggs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/aggs/): complete subsection reference.

<a id="schema-end_time"></a>

### end_time property

Type: `"string"`. Optional.

Fetch security evidence whose timestamp &lt;= end\_time format: unix\_timestamp|RFC 3339 Optional:
If not specified, then the end\_time will be evaluated to start\_time+10m If start\_time is not
specified, then the end\_time will be evaluated to &lt;current time&gt;.

<a id="schema-evidences"></a>

### evidences property

Type: `["list", "string"]`. Computed.

List of security evidences that matched the query. Contains no more than 500 messages.

- [last_sort_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/last_sort_values/): complete subsection reference.

<a id="schema-limit"></a>

### limit property

Type: `"number"`. Optional.

Limits the number of security evidence returned in the response Optional: If not specified, first or
last 500 security evidence that matches the query (depending on the sort order) will be returned in
the response. The maximum value for limit is 500.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace fetch security evidence for a given namespace.

<a id="schema-query"></a>

### query property

Type: `"string"`. Optional.

Query is used to specify the list of matchers syntax for query := \{\[&lt;matcher&gt;\]\}
&lt;matcher&gt; := &lt;field\_name&gt;&lt;operator&gt;'&lt;value&gt;' &lt;field\_name&gt; := string
One or more of these fields in the security evidence may be specified in the query. Domain - domain
endpoint - endpoint evidence\_id - evidence ID..

<a id="schema-search_after"></a>

### search_after property

Type: `"bool"`. Optional.

Search After is used to retrieve large number of log messages (or all log messages) that matches the
query. If search\_after is set to true, the sort\_values in the response can be used in the API to
fetch the next batch of logs. The number of messages in each batch is determined by the limit field.

<a id="schema-sort"></a>

### sort property

Type: `"string"`. Optional.

\[Enum: DESCENDING|ASCENDING\] Sort algorithm Sort in descending order Sort in ascending order.
Possible values are \`DESCENDING\`, \`ASCENDING\`. Defaults to \`DESCENDING\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ASCENDING","DESCENDING"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DESCENDING",
    "ASCENDING"),
}
```

<a id="schema-sort_by"></a>

### sort_by property

Type: `"string"`. Optional.

Optional: default is sort by last\_event\_time.

- [sort_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/sort_values/): complete subsection reference.

<a id="schema-start_time"></a>

### start_time property

Type: `"string"`. Optional.

Fetch security evidence whose timestamp &gt;= start\_time format: unix\_timestamp|RFC 3339 Optional:
If not specified, then the start\_time will be evaluated to end\_time-10m If end\_time is not
specified, then the start\_time will be evaluated to &lt;current time&gt;-10m.

<a id="schema-total_hits"></a>

### total_hits property

Type: `"string"`. Computed.

Total number of security events that matched the query.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `aggs` | [aggs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/aggs/#section) |
| `end_time` | [end_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/#schema-end_time) |
| `evidences` | [evidences](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/#schema-evidences) |
| `last_sort_values` | [last_sort_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/last_sort_values/#section) |
| `last_sort_values.last_doc_id` | [last_sort_values.last_doc_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/last_sort_values/#schema-last_sort_values--last_doc_id) |
| `last_sort_values.last_timestamp` | [last_sort_values.last_timestamp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/last_sort_values/#schema-last_sort_values--last_timestamp) |
| `limit` | [limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/#schema-limit) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/#schema-namespace) |
| `query` | [query](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/#schema-query) |
| `search_after` | [search_after](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/#schema-search_after) |
| `sort` | [sort](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/#schema-sort) |
| `sort_by` | [sort_by](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/#schema-sort_by) |
| `sort_values` | [sort_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/sort_values/#section) |
| `sort_values.last_doc_id` | [sort_values.last_doc_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/sort_values/#schema-sort_values--last_doc_id) |
| `sort_values.last_timestamp` | [sort_values.last_timestamp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/sort_values/#schema-sort_values--last_timestamp) |
| `start_time` | [start_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/#schema-start_time) |
| `total_hits` | [total_hits](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/#schema-total_hits) |
