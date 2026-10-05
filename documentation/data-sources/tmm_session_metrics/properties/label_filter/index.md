---
page_title: "label_filter"
subcategory: ""
description: "List of label filter expressions of the form 'label key' QueryOp 'value'. Response will only contain data that matches all the conditions specified in the label_filter."
xcsh_docs: {"aliases": ["label filter"], "body_bytes": 2467, "body_sha256": "sha256:87e3d1a8b0ad859d5a8d6ca9d100e390ee5736b4684fdccef47c075c4df6000f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "path": "documentation/data-sources/tmm_session_metrics/properties/label_filter/index.md", "product": "distributed-cloud", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1010000131313013-2120202023111032-2233100230133022-1111200001201310-1222203030223230-3101003310111320-3031311033113202-1132000000301312", "registry_path": "docs/guides/data-sources--tmm_session_metrics--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["label_filter"], "schema_version": 1, "sections": [{"aliases": ["label filter label"], "anchor": "schema-label_filter--label", "description": "Metrics used to construct the session metrics are tagged with these labels and therefore the metrics can be sliced and diced based on one or more of these labels. Indicates the field not being set Identifies the workspace where the service is deployed Identifies the virtual server. Possible values are", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["label_filter", "label"], "syntax": "attribute", "type": "string"}, {"aliases": ["label filter op"], "anchor": "schema-label_filter--op", "description": "The operator to use when filtering metrics based on label values. Possible values are `EQ`, `NEQ`. Defaults to `EQ`.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["label_filter", "op"], "syntax": "attribute", "type": "string"}, {"aliases": ["label filter value"], "anchor": "schema-label_filter--value", "description": "Value. Value of the label.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["label_filter", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/label_filter/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of label filter expressions of the form 'label key' QueryOp 'value'. Response will only contain data that matches all the conditions specified in the label_filter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/)
- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
