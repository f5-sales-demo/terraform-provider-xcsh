---
page_title: "enable_encryption"
subcategory: "Infrastructure"
description: "Information related to disk encryption."
xcsh_docs: {"aliases": ["enable encryption"], "body_bytes": 2503, "body_sha256": "sha256:ea51ba98b6bcfbec1cd9466494ba938e5f5c5e24c6baca162f7797f9b431b034", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:enable_encryption", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:reference", "path": "documentation/data-sources/gcp_vpc_site/properties/enable_encryption/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2230311030212003-2200003030123112-0333131120113200-1312032120021222-3102300300031323-1102000322031212-3102223002032122-2323112010030033", "registry_path": "docs/guides/data-sources--gcp_vpc_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_encryption"], "schema_version": 1, "sections": [{"aliases": ["kms key resource id"], "anchor": "schema-enable_encryption--kms_key_resource_id", "description": "GCP KMS Key to be used to encrypt the disk attached to the VM.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:enable_encryption", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_encryption", "kms_key_resource_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["kms key ring id"], "anchor": "schema-enable_encryption--kms_key_ring_id", "description": "Key ring in which the CMK to be used to encrypt is present.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:enable_encryption", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_encryption", "kms_key_ring_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/enable_encryption/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Information related to disk encryption.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_encryption

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- enable_encryption

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

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

<a id="schema-enable_encryption--kms_key_resource_id"></a>

### kms_key_resource_id property

Type: `"string"`. Computed.

GCP KMS Key to be used to encrypt the disk attached to the VM.

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

<a id="schema-enable_encryption--kms_key_ring_id"></a>

### kms_key_ring_id property

Type: `"string"`. Computed.

Key ring in which the CMK to be used to encrypt is present.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
