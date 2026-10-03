---
page_title: "primary.default_rr_set_group.cds_record.values.sha256_digest"
subcategory: "DNS"
description: "Configuration parameter for sha256 digest."
xcsh_docs: {"aliases": ["primary default rr set group cds record values sha256 digest"], "body_bytes": 2749, "body_sha256": "sha256:5d958f894baa9affbb12b074335ba937056b9a08c03ce49f3ab97fefc2d87d9e", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:cds_record:values:sha256_digest", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:cds_record:values", "path": "documentation/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha256_digest/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2231130131103101-0130122012121130-3211003230012003-0130223100003331-3223211002130320-1332012313033213-0203122220232300-3301303321003311", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "cds_record", "values", "sha256_digest"], "schema_version": 1, "sections": [{"aliases": ["primary default rr set group cds record values sha256 digest digest"], "anchor": "schema-primary--default_rr_set_group--cds_record--values--sha256_digest--digest", "description": "The 'digest' is the DS key and the actual contents of the DS record.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:cds_record:values:sha256_digest", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "cds_record", "values", "sha256_digest", "digest"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha256_digest/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for sha256 digest.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.cds_record.values.sha256_digest

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/)
- [primary.default_rr_set_group.cds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/)
- [primary.default_rr_set_group.cds_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/)
- primary.default_rr_set_group.cds_record.values.sha256_digest

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 digest.

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

<a id="schema-primary--default_rr_set_group--cds_record--values--sha256_digest--digest"></a>

### digest property

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [primary.default_rr_set_group.cds_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
