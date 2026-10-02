---
page_title: "enable_encryption"
subcategory: "Infrastructure"
description: "Information related to disk encryption."
xcsh_docs: {"aliases": ["enable encryption"], "body_bytes": 3256, "body_sha256": "sha256:f5720b4b0824978714fb4966e696ab099b7c7965e6f08847846b7d10977044c1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:enable_encryption", "parent_id": "xcsh-docs:resources:azure_vnet_site:reference", "path": "documentation/resources/azure_vnet_site/properties/enable_encryption/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1000310322001320-2311023113002011-3103023303010213-3012300211031000-1221202202331303-1003131333000322-1213302333210100-2312200331021300", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-003.md", "relationships": [{"anchor": "schema-enable_encryption--disk_encryption_set_id", "enforcement": "provider-schema", "group": "enable_encryption:RequiredObjectAttributes:disk_encryption_set_id,resource_group", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:enable_encryption", "type": "requires"}, {"anchor": "schema-enable_encryption--resource_group", "enforcement": "provider-schema", "group": "enable_encryption:RequiredObjectAttributes:disk_encryption_set_id,resource_group", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:enable_encryption", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_encryption"], "schema_version": 1, "sections": [{"aliases": ["disk encryption set id"], "anchor": "schema-enable_encryption--disk_encryption_set_id", "description": "Azure Disk Encryption Set to be used to encrypt the disk attached to the VM.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:enable_encryption", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_encryption", "disk_encryption_set_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["resource group"], "anchor": "schema-enable_encryption--resource_group", "description": "The resource group in which the Disk Encryption Set is present.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:enable_encryption", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_encryption", "resource_group"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/enable_encryption/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Information related to disk encryption.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_encryption

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- enable_encryption

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("disk_encryption_set_id",
    "resource_group")}
```

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

Terraform syntax:

```terraform
enable_encryption {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-enable_encryption--disk_encryption_set_id"></a>

### disk_encryption_set_id property

Type: `"string"`. Optional.

Azure Disk Encryption Set to be used to encrypt the disk attached to the VM.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-enable_encryption--resource_group"></a>

### resource_group property

Type: `"string"`. Optional.

The resource group in which the Disk Encryption Set is present.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
