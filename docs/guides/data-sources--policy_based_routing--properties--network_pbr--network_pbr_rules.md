---
page_title: "network_pbr.network_pbr_rules"
subcategory: ""
description: "network_pbr.network_pbr_rules for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 5688, "body_sha256": "sha256:06d2d3d298b55f6899d25340acff4209fc708ed36a5f588484883361c863521e", "canonical_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_traffic", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_udp_traffic", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:any", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:applications", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:forwarding_class_list", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:metadata", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:prefix_list", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules:protocol_port_range"], "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:network_pbr_rules", "parent_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr", "path": "docs/guides/data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["network_pbr", "network_pbr_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "network_pbr.network_pbr_rules for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# network_pbr.network_pbr_rules

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
- [Property reference](data-sources--policy_based_routing--reference.md)
- [network_pbr](data-sources--policy_based_routing--properties--network_pbr.md)
- network_pbr.network_pbr_rules

<a id="section"></a>

Type: `"list"`. Computed.

L3/L4 Destination Routing Rules. Network(L3/L4) routing policy rule.

Upstream description:

Network(L3/L4) routing policy rule.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [all_tcp_traffic](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--all_tcp_traffic.md): complete subsection reference.

- [all_traffic](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--all_traffic.md): complete subsection reference.

- [all_udp_traffic](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--all_udp_traffic.md): complete subsection reference.

- [any](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--any.md): complete subsection reference.

- [applications](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--applications.md): complete subsection reference.

<a id="schema-network_pbr--network_pbr_rules--dns_name"></a>

### dns_name property

Type: `"string"`. Computed.

Exclusive with \[any ip\_prefix\_set prefix\_list\] Resolve hostname to GET the IP.

Upstream description:

Exclusive with \[any ip\_prefix\_set prefix\_list\] Resolve hostname to GET the IP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [forwarding_class_list](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--forwarding_class_list.md): complete subsection reference.

- [ip_prefix_set](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set.md): complete subsection reference.

- [metadata](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--metadata.md): complete subsection reference.

- [prefix_list](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--prefix_list.md): complete subsection reference.

- [protocol_port_range](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--protocol_port_range.md): complete subsection reference.

## Next pages

- [network_pbr.network_pbr_rules.all_tcp_traffic](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--all_tcp_traffic.md)
- [network_pbr.network_pbr_rules.all_traffic](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--all_traffic.md)
- [network_pbr.network_pbr_rules.all_udp_traffic](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--all_udp_traffic.md)
- [network_pbr.network_pbr_rules.any](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--any.md)
- [network_pbr.network_pbr_rules.applications](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--applications.md)
- [network_pbr.network_pbr_rules.forwarding_class_list](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--forwarding_class_list.md)
- [network_pbr.network_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--ip_prefix_set.md)
- [network_pbr.network_pbr_rules.metadata](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--metadata.md)
- [network_pbr.network_pbr_rules.prefix_list](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--prefix_list.md)
- [network_pbr.network_pbr_rules.protocol_port_range](data-sources--policy_based_routing--properties--network_pbr--network_pbr_rules--protocol_port_range.md)
- [network_pbr](data-sources--policy_based_routing--properties--network_pbr.md)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
