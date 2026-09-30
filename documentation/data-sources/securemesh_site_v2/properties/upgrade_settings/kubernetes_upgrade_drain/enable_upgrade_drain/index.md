---
page_title: "upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain"
subcategory: ""
description: "upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 5557, "body_sha256": "sha256:d5f3752c1a00308cee6aec5bb707cc50700356d2936f1d9f8b6096f735ef7c49", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain:enable_vega_upgrade_mode"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain", "path": "documentation/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/enable_upgrade_drain/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain", "enable_upgrade_drain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/enable_upgrade_drain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
