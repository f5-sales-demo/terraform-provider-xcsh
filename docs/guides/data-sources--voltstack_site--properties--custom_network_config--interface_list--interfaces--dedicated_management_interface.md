---
page_title: "custom_network_config.interface_list.interfaces.dedicated_management_interface"
subcategory: ""
description: "custom_network_config.interface_list.interfaces.dedicated_management_interface for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 5051, "body_sha256": "sha256:862ddab87c3b455e81ea83bc8e9102fddf63aaa1a143cf634422301257f01fd2", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dedicated_management_interface", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dedicated_management_interface:cluster"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dedicated_management_interface", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces", "path": "docs/guides/data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_management_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces.dedicated_management_interface for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.dedicated_management_interface

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_network_config](data-sources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.interface_list](data-sources--voltstack_site--properties--custom_network_config--interface_list.md)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces.md)
- custom_network_config.interface_list.interfaces.dedicated_management_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for dedicated management interface.

Upstream description:

Dedicated Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]"
}
```

## Direct properties

- [cluster](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface--cluster.md): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--device"></a>

### device property

Type: `"string"`. Computed.

Name of the device for which interface is configured.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--mtu"></a>

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

<a id="schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--node"></a>

### node property

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

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

## Next pages

- [custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface--cluster.md)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
