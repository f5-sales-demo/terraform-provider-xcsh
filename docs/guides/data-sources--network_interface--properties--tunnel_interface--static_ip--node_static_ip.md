---
page_title: "tunnel_interface.static_ip.node_static_ip"
subcategory: ""
description: "tunnel_interface.static_ip.node_static_ip for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 3145, "body_sha256": "sha256:27366498ab43b8be15254820ce7acf4c854bb1282fb921f5b600af96b0b37cc3", "canonical_id": "xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip:node_static_ip", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip:node_static_ip", "parent_id": "xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip", "path": "docs/guides/data-sources--network_interface--properties--tunnel_interface--static_ip--node_static_ip.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tunnel_interface", "static_ip", "node_static_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/tunnel_interface/static_ip/node_static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tunnel_interface.static_ip.node_static_ip for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tunnel_interface.static_ip.node_static_ip

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md)
- [Property reference](data-sources--network_interface--reference.md)
- [tunnel_interface](data-sources--network_interface--properties--tunnel_interface.md)
- [tunnel_interface.static_ip](data-sources--network_interface--properties--tunnel_interface--static_ip.md)
- tunnel_interface.static_ip.node_static_ip

<a id="section"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-tunnel_interface--static_ip--node_static_ip--default_gw"></a>

### default_gw property

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-tunnel_interface--static_ip--node_static_ip--dns_server"></a>

### dns_server property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="schema-tunnel_interface--static_ip--node_static_ip--ip_address"></a>

### ip_address property

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

## Next pages

- [tunnel_interface.static_ip](data-sources--network_interface--properties--tunnel_interface--static_ip.md)
- [xcsh_network_interface](../data-sources/network_interface.md)
