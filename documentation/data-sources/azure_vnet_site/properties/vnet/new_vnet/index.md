---
page_title: "vnet.new_vnet"
subcategory: "Infrastructure"
description: "Parameters to create a new Azure VNet."
xcsh_docs: {"aliases": ["vnet new vnet"], "body_bytes": 4130, "body_sha256": "sha256:58249c321a5e478809eef2d4928febfd2076ef7ed88d7f32e3decba0104c06d6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:vnet:new_vnet:autogenerate"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:new_vnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet", "path": "documentation/data-sources/azure_vnet_site/properties/vnet/new_vnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0202231211233031-0113312331300012-2230131111311213-1030221212230113-2032203121323033-1231313113330100-1132011201220213-2210012112212320", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vnet", "new_vnet"], "schema_version": 1, "sections": [{"aliases": ["autogenerate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:new_vnet:autogenerate", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vnet", "new_vnet", "autogenerate"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-vnet--new_vnet--name", "description": "Exclusive with Specify the VNet Name.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:new_vnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vnet", "new_vnet", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary ipv4"], "anchor": "schema-vnet--new_vnet--primary_ipv4", "description": "IPv4 CIDR block for this VNet. It has to be private address space.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:new_vnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vnet", "new_vnet", "primary_ipv4"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/vnet/new_vnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters to create a new Azure VNet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vnet.new_vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/)
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

- [autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/new_vnet/autogenerate/): complete subsection reference.

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

- [vnet.new_vnet.autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/new_vnet/autogenerate/)
- [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
