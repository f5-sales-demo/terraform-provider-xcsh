---
page_title: "custom_network_config.interface_list.interfaces.tunnel_interface"
subcategory: ""
description: "custom_network_config.interface_list.interfaces.tunnel_interface for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 6221, "body_sha256": "sha256:334566b155ce6b3068be2fa1621a302f1d83cf920eb06feaebddf6d78691412e", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:site_local_inside_network", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:site_local_network", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:tunnel"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces", "path": "docs/guides/data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "tunnel_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces.tunnel_interface for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.interface_list.interfaces.tunnel_interface

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_network_config](data-sources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.interface_list](data-sources--voltstack_site--properties--custom_network_config--interface_list.md)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces.md)
- custom_network_config.interface_list.interfaces.tunnel_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for tunnel interface.

Upstream description:

Tunnel Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-node_choice": "[\"node\"]"
}
```

## Direct properties

<a id="schema-custom_network_config--interface_list--interfaces--tunnel_interface--mtu"></a>

### mtu property

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="schema-custom_network_config--interface_list--interfaces--tunnel_interface--node"></a>

### node property

Type: `"string"`. Computed.

Exclusive with \[\] Configuration will apply to a given device on the given node.

Upstream description:

Exclusive with \[\] Configuration will apply to a given device on the given node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-custom_network_config--interface_list--interfaces--tunnel_interface--priority"></a>

### priority property

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface--site_local_inside_network.md): complete subsection reference.

- [site_local_network](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface--site_local_network.md): complete subsection reference.

- [static_ip](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface--static_ip.md): complete subsection reference.

- [tunnel](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface--tunnel.md): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.tunnel_interface.site_local_inside_network](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface--site_local_inside_network.md)
- [custom_network_config.interface_list.interfaces.tunnel_interface.site_local_network](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface--site_local_network.md)
- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface--static_ip.md)
- [custom_network_config.interface_list.interfaces.tunnel_interface.tunnel](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface--tunnel.md)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
