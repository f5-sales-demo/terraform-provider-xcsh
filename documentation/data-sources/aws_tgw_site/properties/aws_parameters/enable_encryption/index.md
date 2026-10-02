---
page_title: "aws_parameters.enable_encryption"
subcategory: ""
description: "Information related to disk encryption."
xcsh_docs: {"aliases": ["aws parameters enable encryption"], "body_bytes": 1914, "body_sha256": "sha256:474fa06e236245429d888fe9f208a0860d7cc95afc9506ed63bb05fa4ba3e420", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:enable_encryption", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters", "path": "documentation/data-sources/aws_tgw_site/properties/aws_parameters/enable_encryption/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3010302002120100-2103002021320330-3110033102312003-2010320201333323-3230020212131301-2330231333302020-1201123130210012-0300222330022213", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "enable_encryption"], "schema_version": 1, "sections": [{"aliases": ["kms key id"], "anchor": "schema-aws_parameters--enable_encryption--kms_key_id", "description": "AWS KMS Key to be used to encrypt the disk attached to the VM.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:enable_encryption", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "enable_encryption", "kms_key_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/aws_parameters/enable_encryption/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Information related to disk encryption.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.enable_encryption

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/)
- aws_parameters.enable_encryption

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

<a id="schema-aws_parameters--enable_encryption--kms_key_id"></a>

### kms_key_id property

Type: `"string"`. Computed.

AWS KMS Key to be used to encrypt the disk attached to the VM.

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

- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
