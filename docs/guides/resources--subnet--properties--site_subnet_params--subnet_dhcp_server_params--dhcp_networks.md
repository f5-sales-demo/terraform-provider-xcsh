---
page_title: "site_subnet_params.subnet_dhcp_server_params.dhcp_networks"
subcategory: ""
description: "site_subnet_params.subnet_dhcp_server_params.dhcp_networks for xcsh_subnet."
xcsh_docs: {"aliases": [], "body_bytes": 2598, "body_sha256": "sha256:029b8e752bbd5a0a82c06a3f80877031ff4fe54dbd0d09e257ac1fc90f8a4a4b", "canonical_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params:dhcp_networks", "child_ids": [], "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params:dhcp_networks", "parent_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params", "path": "docs/guides/resources--subnet--properties--site_subnet_params--subnet_dhcp_server_params--dhcp_networks.md", "provider_name": "subnet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_subnet_params", "subnet_dhcp_server_params", "dhcp_networks"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/dhcp_networks/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_subnet_params.subnet_dhcp_server_params.dhcp_networks for xcsh_subnet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_subnet_params.subnet_dhcp_server_params.dhcp_networks

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md)
- [Property reference](resources--subnet--reference.md)
- [site_subnet_params](resources--subnet--properties--site_subnet_params.md)
- [site_subnet_params.subnet_dhcp_server_params](resources--subnet--properties--site_subnet_params--subnet_dhcp_server_params.md)
- site_subnet_params.subnet_dhcp_server_params.dhcp_networks

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP server can allocate IP addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-site_subnet_params--subnet_dhcp_server_params--dhcp_networks--network_prefix"></a>

### network_prefix property

Type: `"string"`. Optional.

Exclusive with \[\] Network prefix for subnet.

Upstream description:

Exclusive with \[\] Network prefix for subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

## Next pages

- [site_subnet_params.subnet_dhcp_server_params](resources--subnet--properties--site_subnet_params--subnet_dhcp_server_params.md)
- [xcsh_subnet](../resources/subnet.md)
