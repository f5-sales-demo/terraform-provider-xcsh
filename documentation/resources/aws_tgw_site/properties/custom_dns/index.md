---
page_title: "custom_dns"
subcategory: ""
description: "Custom DNS is the configured for specify CE site."
xcsh_docs: {"aliases": ["custom dns"], "body_bytes": 2799, "body_sha256": "sha256:a6d2afd2f3790595d2db13db22b6a9279785bba733f49d5696f683535c2c3e5e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:custom_dns", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "documentation/resources/aws_tgw_site/properties/custom_dns/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-0310130113103331-1321231230201233-0112322012222330-2020103011100211-0311022022233022-2032112020200200-0003021332313331-1212231013022131", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_dns"], "schema_version": 1, "sections": [{"aliases": ["custom dns inside nameserver"], "anchor": "schema-custom_dns--inside_nameserver", "description": "Optional DNS server IP to be used for name resolution in inside network.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:custom_dns", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_dns", "inside_nameserver"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom dns outside nameserver"], "anchor": "schema-custom_dns--outside_nameserver", "description": "Optional DNS server IP to be used for name resolution in outside network.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:custom_dns", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_dns", "outside_nameserver"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/custom_dns/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Custom DNS is the configured for specify CE site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_dns

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- custom_dns

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom DNS is the configured for specify CE site.

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
custom_dns {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_dns--inside_nameserver"></a>

### inside_nameserver property

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in inside network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-29T12:14:42+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="schema-custom_dns--outside_nameserver"></a>

### outside_nameserver property

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in outside network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-29T12:14:42+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```
