---
page_title: "data"
subcategory: ""
description: "Data contains time-series TMM Session data."
xcsh_docs: {"aliases": ["data"], "body_bytes": 4081, "body_sha256": "sha256:8f3eecd63d89c22efad963f106fb330e5e4dc008ace629d93d6c20f09eb67e5e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "path": "documentation/data-sources/tmm_session_metrics/properties/data/index.md", "product": "distributed-cloud", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1200130131121221-1313332001213031-2020130202233021-1132311033211230-0000333213330132-1010302302021011-2133133221301123-2223021122031010", "registry_path": "docs/guides/data-sources--tmm_session_metrics--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["data"], "schema_version": 1, "sections": [{"aliases": ["data metric"], "anchor": "section", "description": "Metric. List of metrics.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["data", "metric"], "syntax": "attribute", "type": "object"}, {"aliases": ["data type"], "anchor": "schema-data--type", "description": "X-displayName: TMM Session Metric Type' FieldSelector specifies the metrics that can be queried for virtual servers. Indicates field not being set x-unit: 'count' Total number of active sessions x-unit: 'count' Total number of allowed sessions x-unit: 'count' Total number of denied Session.. Possible values are", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "type"], "syntax": "attribute", "type": "string"}, {"aliases": ["data unit"], "anchor": "schema-data--unit", "description": "UnitType is enumeration of units for scalar fields. Possible values are `UNIT_MILLISECONDS`, `UNIT_SECONDS`, `UNIT_MINUTES`, `UNIT_HOURS`, `UNIT_DAYS`, `UNIT_BYTES`, `UNIT_KBYTES`, `UNIT_MBYTES`, `UNIT_GBYTES`, `UNIT_TBYTES`, `UNIT_KIBIBYTES`, `UNIT_MIBIBYTES`, `UNIT_GIBIBYTES`, `UNIT_TEBIBYTES`,", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "unit"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/data/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data contains time-series TMM Session data.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data

Breadcrumbs:

- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/)
- data

<a id="section"></a>

Type: `"list"`. Computed.

Data contains time-series TMM Session data.

## Direct properties

- [metric](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/): complete subsection reference.

<a id="schema-data--type"></a>

### type property

Type: `"string"`. Computed.

\[Enum:
METRIC\_TYPE\_NONE|METRIC\_TYPE\_ACTIVE|METRIC\_TYPE\_ALLOWED|METRIC\_TYPE\_DENIED|METRIC\_TYPE\_TOTAL|METRIC\_TYPE\_LOGOUT|METRIC\_TYPE\_ESTABLISHED\_TIMEOUT|METRIC\_TYPE\_EVALUATION\_TIMEOUT|METRIC\_TYPE\_ADMIN\_TERMINATED\]
X-displayName: TMM Session Metric Type' FieldSelector specifies the metrics that can be queried for
virtual servers. Indicates field not being set x-unit: 'count' Total number of active sessions
x-unit: 'count' Total number of allowed sessions x-unit: 'count' Total number of denied Session..
Possible values are \`METRIC\_TYPE\_NONE\`, \`METRIC\_TYPE\_ACTIVE\`, \`METRIC\_TYPE\_ALLOWED\`,
\`METRIC\_TYPE\_DENIED\`, \`METRIC\_TYPE\_TOTAL\`, \`METRIC\_TYPE\_LOGOUT\`,
\`METRIC\_TYPE\_ESTABLISHED\_TIMEOUT\`, \`METRIC\_TYPE\_EVALUATION\_TIMEOUT\`,
\`METRIC\_TYPE\_ADMIN\_TERMINATED\`. Defaults to \`METRIC\_TYPE\_NONE\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("METRIC_TYPE_NONE",
    "METRIC_TYPE_ACTIVE",
    "METRIC_TYPE_ALLOWED",
    "METRIC_TYPE_DENIED",
    "METRIC_TYPE_TOTAL",
    "METRIC_TYPE_LOGOUT",
    "METRIC_TYPE_ESTABLISHED_TIMEOUT",
    "METRIC_TYPE_EVALUATION_TIMEOUT",
    "METRIC_TYPE_ADMIN_TERMINATED"),
}
```

<a id="schema-data--unit"></a>

### unit property

Type: `"string"`. Computed.

\[Enum:
UNIT\_MILLISECONDS|UNIT\_SECONDS|UNIT\_MINUTES|UNIT\_HOURS|UNIT\_DAYS|UNIT\_BYTES|UNIT\_KBYTES|UNIT\_MBYTES|UNIT\_GBYTES|UNIT\_TBYTES|UNIT\_KIBIBYTES|UNIT\_MIBIBYTES|UNIT\_GIBIBYTES|UNIT\_TEBIBYTES|UNIT\_BITS\_PER\_SECOND|UNIT\_BYTES\_PER\_SECOND|UNIT\_KBITS\_PER\_SECOND|UNIT\_KBYTES\_PER\_SECOND|UNIT\_MBITS\_PER\_SECOND|UNIT\_MBYTES\_PER\_SECOND|UNIT\_CONNECTIONS\_PER\_SECOND|UNIT\_ERRORS\_PER\_SECOND|UNIT\_PACKETS\_PER\_SECOND|UNIT\_REQUESTS\_PER\_SECOND|UNIT\_PACKETS|UNIT\_PERCENTAGE|UNIT\_COUNT\]
UnitType is enumeration of units for scalar fields. Possible values are \`UNIT\_MILLISECONDS\`,
\`UNIT\_SECONDS\`, \`UNIT\_MINUTES\`, \`UNIT\_HOURS\`, \`UNIT\_DAYS\`, \`UNIT\_BYTES\`,
\`UNIT\_KBYTES\`, \`UNIT\_MBYTES\`, \`UNIT\_GBYTES\`, \`UNIT\_TBYTES\`, \`UNIT\_KIBIBYTES\`,
\`UNIT\_MIBIBYTES\`, \`UNIT\_GIBIBYTES\`, \`UNIT\_TEBIBYTES\`, \`UNIT\_BITS\_PER\_SECOND\`,
\`UNIT\_BYTES\_PER\_SECOND\`, \`UNIT\_KBITS\_PER\_SECOND\`, \`UNIT\_KBYTES\_PER\_SECOND\`,
\`UNIT\_MBITS\_PER\_SECOND\`, \`UNIT\_MBYTES\_PER\_SECOND\`, \`UNIT\_CONNECTIONS\_PER\_SECOND\`,
\`UNIT\_ERRORS\_PER\_SECOND\`, \`UNIT\_PACKETS\_PER\_SECOND\`, \`UNIT\_REQUESTS\_PER\_SECOND\`,
\`UNIT\_PACKETS\`, \`UNIT\_PERCENTAGE\`, \`UNIT\_COUNT\`. Defaults to \`UNIT\_MILLISECONDS\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNIT_MILLISECONDS",
    "UNIT_SECONDS",
    "UNIT_MINUTES",
    "UNIT_HOURS",
    "UNIT_DAYS",
    "UNIT_BYTES",
    "UNIT_KBYTES",
    "UNIT_MBYTES",
    "UNIT_GBYTES",
    "UNIT_TBYTES",
    "UNIT_KIBIBYTES",
    "UNIT_MIBIBYTES",
    "UNIT_GIBIBYTES",
    "UNIT_TEBIBYTES",
    "UNIT_BITS_PER_SECOND",
    "UNIT_BYTES_PER_SECOND",
    "UNIT_KBITS_PER_SECOND",
    "UNIT_KBYTES_PER_SECOND",
    "UNIT_MBITS_PER_SECOND",
    "UNIT_MBYTES_PER_SECOND",
    "UNIT_CONNECTIONS_PER_SECOND",
    "UNIT_ERRORS_PER_SECOND",
    "UNIT_PACKETS_PER_SECOND",
    "UNIT_REQUESTS_PER_SECOND",
    "UNIT_PACKETS",
    "UNIT_PERCENTAGE",
    "UNIT_COUNT"),
}
```
