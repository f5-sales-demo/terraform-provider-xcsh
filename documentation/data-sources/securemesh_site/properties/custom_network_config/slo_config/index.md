---
page_title: "custom_network_config.slo_config"
subcategory: ""
description: "custom_network_config.slo_config for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 6433, "body_sha256": "sha256:06a56284f48ddca429412a5f514fa48cbed96572a30f887266aec4c1f63ffc46", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:dc_cluster_group", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:no_dc_cluster_group", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:no_static_routes", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:no_v6_static_routes", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:static_routes", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:static_v6_routes"], "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config", "path": "documentation/data-sources/securemesh_site/properties/custom_network_config/slo_config/index.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["custom_network_config", "slo_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/slo_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.slo_config for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.slo_config

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/)
- custom_network_config.slo_config

<a id="section"></a>

Type: `"single"`. Computed.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

## Direct properties

- [dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/dc_cluster_group/): complete subsection reference.

<a id="schema-custom_network_config--slo_config--labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Add Labels for this network, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-custom_network_config--slo_config--nameserver"></a>

### nameserver property

Type: `"string"`. Computed.

Optional DNS V4 server IP to be used for name resolution.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/no_dc_cluster_group/): complete subsection reference.

- [no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/no_static_routes/): complete subsection reference.

- [no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/no_v6_static_routes/): complete subsection reference.

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_routes/): complete subsection reference.

- [static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/): complete subsection reference.

<a id="schema-custom_network_config--slo_config--vip"></a>

### vip property

Type: `"string"`. Computed.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [custom_network_config.slo_config.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/dc_cluster_group/)
- [custom_network_config.slo_config.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/no_dc_cluster_group/)
- [custom_network_config.slo_config.no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/no_static_routes/)
- [custom_network_config.slo_config.no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/no_v6_static_routes/)
- [custom_network_config.slo_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_routes/)
- [custom_network_config.slo_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
