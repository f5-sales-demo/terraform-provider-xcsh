---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_app_security_evidence."
xcsh_docs: {"aliases": [], "body_bytes": 7220, "body_sha256": "sha256:f70ee883b3f46863e026214546dafa29c6d9547a68b8cbc539010e0e01011185", "child_ids": ["xcsh-docs:data-sources:app_security_evidence:properties:aggs", "xcsh-docs:data-sources:app_security_evidence:properties:last_sort_values", "xcsh-docs:data-sources:app_security_evidence:properties:sort_values"], "collection_id": "xcsh-docs:data-sources:app_security_evidence:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_security_evidence:reference", "parent_id": "xcsh-docs:data-sources:app_security_evidence:fundamentals", "path": "documentation/data-sources/app_security_evidence/properties/index.md", "provider_name": "app_security_evidence", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_security_evidence/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_app_security_evidence.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [aggs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/aggs/)
- [last_sort_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/last_sort_values/)
- [sort_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/sort_values/)
- [xcsh_app_security_evidence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/)
