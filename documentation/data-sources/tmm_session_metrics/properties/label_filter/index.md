---
page_title: "label_filter"
subcategory: ""
description: "List of label filter expressions of the form 'label key' QueryOp 'value'. Response will only contain data that matches all the conditions specified in the label_filter."
xcsh_docs: {"aliases": ["label filter"], "body_bytes": 2823, "body_sha256": "sha256:72cb26216607ffafc1cd5cde70992be8aa510d30501ee60a3872b53db0b4350f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "path": "documentation/data-sources/tmm_session_metrics/properties/label_filter/index.md", "product": "distributed-cloud", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1010000131313013-2120202023111032-2233100230133022-1111200001201310-1222203030223230-3101003310111320-3031311033113202-1132000000301312", "registry_path": "docs/guides/data-sources--tmm_session_metrics--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["label_filter"], "schema_version": 1, "sections": [{"aliases": ["label filter label"], "anchor": "schema-label_filter--label", "description": "Metrics used to construct the session metrics are tagged with these labels and therefore the metrics can be sliced and diced based on one or more of these labels. Indicates the field not being set Identifies the workspace where the service is deployed Identifies the virtual server. Possible values are", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["METRIC_LABEL_NAMESPACE", "METRIC_LABEL_NONE", "METRIC_LABEL_VIRTUAL_SERVER"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["label_filter", "label"], "syntax": "attribute", "type": "string"}, {"aliases": ["label filter op"], "anchor": "schema-label_filter--op", "description": "The operator to use when filtering metrics based on label values. Possible values are `EQ`, `NEQ`. Defaults to `EQ`.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["EQ", "NEQ"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["label_filter", "op"], "syntax": "attribute", "type": "string"}, {"aliases": ["label filter value"], "anchor": "schema-label_filter--value", "description": "Value. Value of the label.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["label_filter", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/label_filter/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of label filter expressions of the form 'label key' QueryOp 'value'. Response will only contain data that matches all the conditions specified in the label_filter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# label_filter

Breadcrumbs:

- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/)
- label_filter

<a id="section"></a>

Type: `"list"`. Optional.

List of label filter expressions of the form 'label key' QueryOp 'value'. Response will only contain
data that matches all the conditions specified in the label\_filter.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

## Direct properties

<a id="schema-label_filter--label"></a>

### label property

Type: `"string"`. Optional.

\[Enum: METRIC\_LABEL\_NONE|METRIC\_LABEL\_NAMESPACE|METRIC\_LABEL\_VIRTUAL\_SERVER\] Metrics used
to construct the session metrics are tagged with these labels and therefore the metrics can be
sliced and diced based on one or more of these labels. Indicates the field not being set Identifies
the workspace where the service is deployed Identifies the virtual server. Possible values are
\`METRIC\_LABEL\_NONE\`, \`METRIC\_LABEL\_NAMESPACE\`, \`METRIC\_LABEL\_VIRTUAL\_SERVER\`. Defaults
to \`METRIC\_LABEL\_NONE\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["METRIC_LABEL_NAMESPACE","METRIC_LABEL_NONE","METRIC_LABEL_VIRTUAL_SERVER"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("METRIC_LABEL_NONE",
    "METRIC_LABEL_NAMESPACE",
    "METRIC_LABEL_VIRTUAL_SERVER"),
}
```

<a id="schema-label_filter--op"></a>

### op property

Type: `"string"`. Optional.

\[Enum: EQ|NEQ\] The operator to use when filtering metrics based on label values. Possible values
are \`EQ\`, \`NEQ\`. Defaults to \`EQ\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EQ","NEQ"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("EQ",
    "NEQ"),
}
```

<a id="schema-label_filter--value"></a>

### value property

Type: `"string"`. Optional.

Value. Value of the label.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```
