---
page_title: "segment_vrf.segment_config"
subcategory: ""
description: "segment_vrf.segment_config for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 3900, "body_sha256": "sha256:7048d917f31ebceccc7b0d2b867140f697866b1511985332bdab0b95884677f5", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:no_static_routes", "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:no_v6_static_routes", "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes", "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf", "path": "docs/guides/data-sources--securemesh_site_v2--properties--segment_vrf--segment_config.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["segment_vrf", "segment_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "segment_vrf.segment_config for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_vrf.segment_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [segment_vrf](data-sources--securemesh_site_v2--properties--segment_vrf.md)
- segment_vrf.segment_config

<a id="section"></a>

Type: `"single"`. Computed.

Segment Network Configuration. Segment Network Configuration.

Upstream description:

Segment Network Configuration.

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

<a id="schema-segment_vrf--segment_config--nameserver"></a>

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

- [no_static_routes](data-sources--securemesh_site_v2--properties--segment_vrf--segment_config--no_static_routes.md): complete subsection reference.

- [no_v6_static_routes](data-sources--securemesh_site_v2--properties--segment_vrf--segment_config--no_v6_static_routes.md): complete subsection reference.

<a id="schema-segment_vrf--segment_config--secondary_nameserver"></a>

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

- [static_routes](data-sources--securemesh_site_v2--properties--segment_vrf--segment_config--static_routes.md): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site_v2--properties--segment_vrf--segment_config--static_v6_routes.md): complete subsection reference.

## Next pages

- [segment_vrf.segment_config.no_static_routes](data-sources--securemesh_site_v2--properties--segment_vrf--segment_config--no_static_routes.md)
- [segment_vrf.segment_config.no_v6_static_routes](data-sources--securemesh_site_v2--properties--segment_vrf--segment_config--no_v6_static_routes.md)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--properties--segment_vrf--segment_config--static_routes.md)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--properties--segment_vrf--segment_config--static_v6_routes.md)
- [segment_vrf](data-sources--securemesh_site_v2--properties--segment_vrf.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
