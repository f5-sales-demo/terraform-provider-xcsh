---
page_title: "data"
subcategory: ""
description: "data for xcsh_tmm_session_metrics."
xcsh_docs: {"aliases": [], "body_bytes": 4472, "body_sha256": "sha256:43b3c3f6e962561db3ad8f7855ce2d1051d4ec720d1a3be0ffc5b8a0b479eb5a", "child_ids": ["xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric"], "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "path": "documentation/data-sources/tmm_session_metrics/properties/data/index.md", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["data"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/data/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "data for xcsh_tmm_session_metrics.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [data.metric](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/)
- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
