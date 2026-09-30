---
page_title: "vnet.new_vnet"
subcategory: "Infrastructure"
description: "vnet.new_vnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 3676, "body_sha256": "sha256:569f5bd8047efc3b4af56f2b90c89bd701fbda3f58433c882ca10d38f28dc9cf", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:new_vnet", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:vnet:new_vnet:autogenerate"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:new_vnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet", "path": "docs/guides/data-sources--azure_vnet_site--properties--vnet--new_vnet.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vnet", "new_vnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/vnet/new_vnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vnet.new_vnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# vnet.new_vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [vnet](data-sources--azure_vnet_site--properties--vnet.md)
- vnet.new_vnet

<a id="section"></a>

Type: `"single"`. Computed.

Azure VNet Parameters. Parameters to create a new Azure VNet.

Upstream description:

Parameters to create a new Azure VNet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-name_choice": "[\"autogenerate\",\"name\"]"
}
```

## Direct properties

- [autogenerate](data-sources--azure_vnet_site--properties--vnet--new_vnet--autogenerate.md): complete subsection reference.

<a id="schema-vnet--new_vnet--name"></a>

### name property

Type: `"string"`. Computed.

Exclusive with \[autogenerate\] Specify the VNet Name.

Upstream description:

Exclusive with \[autogenerate\] Specify the VNet Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
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

<a id="schema-vnet--new_vnet--primary_ipv4"></a>

### primary_ipv4 property

Type: `"string"`. Computed.

IPv4 CIDR block for this VNet. It has to be private address space.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  }
}
```

## Next pages

- [vnet.new_vnet.autogenerate](data-sources--azure_vnet_site--properties--vnet--new_vnet--autogenerate.md)
- [vnet](data-sources--azure_vnet_site--properties--vnet.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
