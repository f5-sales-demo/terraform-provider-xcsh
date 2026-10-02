---
page_title: "aws_parameters.enable_encryption"
subcategory: ""
description: "Information related to disk encryption."
xcsh_docs: {"aliases": ["aws parameters enable encryption"], "body_bytes": 2174, "body_sha256": "sha256:87c90c6e8e7cbad2ad864a7bbd81477c03d18dd81ea05d307f0375a63d799ed0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_encryption", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/enable_encryption/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2133303310323133-0210232321223011-2220133231121310-2230230312210302-1301312210121223-3011313310300231-3011001123132333-1121203312013021", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [{"anchor": "schema-aws_parameters--enable_encryption--kms_key_id", "enforcement": "provider-schema", "group": "aws_parameters.enable_encryption:RequiredObjectAttributes:kms_key_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_encryption", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "enable_encryption"], "schema_version": 1, "sections": [{"aliases": ["kms key id"], "anchor": "schema-aws_parameters--enable_encryption--kms_key_id", "description": "AWS KMS Key to be used to encrypt the disk attached to the VM.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_encryption", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "enable_encryption", "kms_key_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/enable_encryption/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Information related to disk encryption.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.enable_encryption

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- aws_parameters.enable_encryption

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("kms_key_id")}
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

<a id="schema-aws_parameters--enable_encryption--kms_key_id"></a>

### kms_key_id property

Type: `"string"`. Optional.

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

- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
