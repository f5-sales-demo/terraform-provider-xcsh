---
page_title: "primary.default_rr_set_group.cds_record.values.sha256_digest"
subcategory: "DNS"
description: "Configuration parameter for sha256 digest."
xcsh_docs: {"aliases": ["primary default rr set group cds record values sha256 digest"], "body_bytes": 3134, "body_sha256": "sha256:050189e341118b4b40759828f940813db14e62876e184b0ad23fd8103376f2bb", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:cds_record:values:sha256_digest", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:cds_record:values", "path": "documentation/resources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha256_digest/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3022132203233302-2131230303203032-1021201100330012-1220002130003110-3130311103021331-1110200021122003-2112203020002200-1033210013210121", "registry_path": "docs/guides/resources--dns_zone--reference--group-001.md", "relationships": [{"anchor": "schema-primary--default_rr_set_group--cds_record--values--sha256_digest--digest", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.cds_record.values.sha256_digest:RequiredObjectAttributes:digest", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:cds_record:values:sha256_digest", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "cds_record", "values", "sha256_digest"], "schema_version": 1, "sections": [{"aliases": ["primary default rr set group cds record values sha256 digest digest"], "anchor": "schema-primary--default_rr_set_group--cds_record--values--sha256_digest--digest", "description": "The 'digest' is the DS key and the actual contents of the DS record.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:cds_record:values:sha256_digest", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "cds_record", "values", "sha256_digest", "digest"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha256_digest/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration parameter for sha256 digest.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.cds_record.values.sha256_digest

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/)
- [primary.default_rr_set_group.cds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/cds_record/)
- [primary.default_rr_set_group.cds_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/)
- primary.default_rr_set_group.cds_record.values.sha256_digest

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha256_digest {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--default_rr_set_group--cds_record--values--sha256_digest--digest"></a>

### digest property

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 64
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
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

## Next pages

- [primary.default_rr_set_group.cds_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
