---
page_title: "label_filter"
subcategory: ""
description: "label_filter for xcsh_tmm_session_metrics."
xcsh_docs: {"aliases": [], "body_bytes": 2368, "body_sha256": "sha256:415da0949c840c975e474049cbc746a831cc049ecb432e2af7362e2dd2b2bc87", "child_ids": [], "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "path": "documentation/data-sources/tmm_session_metrics/properties/label_filter/index.md", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["label_filter"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/label_filter/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "label_filter for xcsh_tmm_session_metrics.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
