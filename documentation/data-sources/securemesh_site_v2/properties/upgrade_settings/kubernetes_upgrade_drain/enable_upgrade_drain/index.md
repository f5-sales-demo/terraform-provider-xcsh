---
page_title: "upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain"
subcategory: ""
description: "Specify batch upgrade settings for worker nodes within a site."
xcsh_docs: {"aliases": ["upgrade settings kubernetes upgrade drain enable upgrade drain"], "body_bytes": 5656, "body_sha256": "sha256:e2eff2a0fee079d6c9812525014bd77965dbc71016b98d73b78792bb0669c673", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain:enable_vega_upgrade_mode"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain", "path": "documentation/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/enable_upgrade_drain/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0021210223223221-1202110301313103-0001033232221120-3133022310111011-1230301200223222-3001001331031032-3021203121012132-2203311323331123", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain", "enable_upgrade_drain"], "schema_version": 1, "sections": [{"aliases": ["upgrade settings kubernetes upgrade drain enable upgrade drain disable vega upgrade mode"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain", "enable_upgrade_drain", "disable_vega_upgrade_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["upgrade settings kubernetes upgrade drain enable upgrade drain drain max unavailable node count"], "anchor": "schema-upgrade_settings--kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count", "description": "Exclusive with", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain", "enable_upgrade_drain", "drain_max_unavailable_node_count"], "syntax": "attribute", "type": "number"}, {"aliases": ["upgrade settings kubernetes upgrade drain enable upgrade drain drain max unavailable node percentage"], "anchor": "schema-upgrade_settings--kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage", "description": "Maximum percentage of nodes unavailable during upgrade draining.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain", "enable_upgrade_drain", "drain_max_unavailable_node_percentage"], "syntax": "attribute", "type": "number"}, {"aliases": ["duration", "upgrade settings kubernetes upgrade drain enable upgrade drain drain node timeout"], "anchor": "schema-upgrade_settings--kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout", "description": "Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is recommended to use the", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain", "enable_upgrade_drain", "drain_node_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["upgrade settings kubernetes upgrade drain enable upgrade drain enable vega upgrade mode"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain:enable_vega_upgrade_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain", "enable_upgrade_drain", "enable_vega_upgrade_mode"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/enable_upgrade_drain/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Specify batch upgrade settings for worker nodes within a site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [upgrade_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/upgrade_settings/)
- [upgrade_settings.kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain

<a id="section"></a>

Type: `"single"`. Computed.

Specify batch upgrade settings for worker nodes within a site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

## Direct properties

- [disable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/): complete subsection reference.

<a id="schema-upgrade_settings--kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count"></a>

### drain_max_unavailable_node_count property

Type: `"number"`. Computed.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="schema-upgrade_settings--kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage"></a>

### drain_max_unavailable_node_percentage property

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="schema-upgrade_settings--kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout"></a>

### drain_node_timeout property

Type: `"number"`. Computed.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/enable_upgrade_drain/enable_vega_upgrade_mode/): complete subsection reference.

## Next pages

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/enable_upgrade_drain/enable_vega_upgrade_mode/)
- [upgrade_settings.kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
