---
page_title: "local_vrf.slo_config"
subcategory: ""
description: "Site local network configuration."
xcsh_docs: {"aliases": ["local vrf slo config"], "body_bytes": 6894, "body_sha256": "sha256:20cff1de1c8105348babac0df43c5ae408ce6c8f3a6803880ccdebedb7a6da83", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:no_static_routes", "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:no_v6_static_routes", "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes", "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf", "path": "documentation/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_vrf", "slo_config"], "schema_version": 1, "sections": [{"aliases": ["local vrf slo config labels"], "anchor": "schema-local_vrf--slo_config--labels", "description": "Add Labels for this network, these labels can be used in firewall policy.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["local vrf slo config nameserver"], "anchor": "schema-local_vrf--slo_config--nameserver", "description": "Optional IPv4 DNS server to be used for name resolution.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "nameserver"], "syntax": "attribute", "type": "string"}, {"aliases": ["local vrf slo config no static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:no_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "no_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["local vrf slo config no v6 static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:no_v6_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "no_v6_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["local vrf slo config secondary nameserver"], "anchor": "schema-local_vrf--slo_config--secondary_nameserver", "description": "Optional Secondary IPv4 DNS server to be used for name resolution.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "secondary_nameserver"], "syntax": "attribute", "type": "string"}, {"aliases": ["local vrf slo config static routes"], "anchor": "section", "description": "Configuration parameter for static routes.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_vrf", "slo_config", "static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["local vrf slo config static v6 routes"], "anchor": "section", "description": "List of IPv6 static routes.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_vrf", "slo_config", "static_v6_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["local vrf slo config vip"], "anchor": "schema-local_vrf--slo_config--vip", "description": "Optional common virtual V4 IP across all nodes to be used as automatic VIP.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "vip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Site local network configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.slo_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [local_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/)
- local_vrf.slo_config

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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

## Direct properties

<a id="schema-local_vrf--slo_config--labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Add Labels for this network, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    }
  },
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

<a id="schema-local_vrf--slo_config--nameserver"></a>

### nameserver property

Type: `"string"`. Computed.

Optional IPv4 DNS server to be used for name resolution.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/no_static_routes/): complete subsection reference.

- [no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/no_v6_static_routes/): complete subsection reference.

<a id="schema-local_vrf--slo_config--secondary_nameserver"></a>

### secondary_nameserver property

Type: `"string"`. Computed.

Optional Secondary IPv4 DNS server to be used for name resolution.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_routes/): complete subsection reference.

- [static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/): complete subsection reference.

<a id="schema-local_vrf--slo_config--vip"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [local_vrf.slo_config.no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/no_static_routes/)
- [local_vrf.slo_config.no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/no_v6_static_routes/)
- [local_vrf.slo_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_routes/)
- [local_vrf.slo_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/)
- [local_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
