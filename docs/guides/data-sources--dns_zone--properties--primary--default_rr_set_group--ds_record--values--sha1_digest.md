---
page_title: "primary.default_rr_set_group.ds_record.values.sha1_digest"
subcategory: "DNS"
description: "primary.default_rr_set_group.ds_record.values.sha1_digest for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 2232, "body_sha256": "sha256:9c1c99cef3585156a696cdbf0b7bd11f82faf4adc338ded5d82148d4cb862ead", "canonical_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha1_digest", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha1_digest", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:ds_record:values", "path": "docs/guides/data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values--sha1_digest.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "default_rr_set_group", "ds_record", "values", "sha1_digest"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha1_digest/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.default_rr_set_group.ds_record.values.sha1_digest for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# primary.default_rr_set_group.ds_record.values.sha1_digest

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md)
- [Property reference](data-sources--dns_zone--reference.md)
- [primary](data-sources--dns_zone--properties--primary.md)
- [primary.default_rr_set_group](data-sources--dns_zone--properties--primary--default_rr_set_group.md)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record.md)
- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values.md)
- primary.default_rr_set_group.ds_record.values.sha1_digest

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 digest.

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

<a id="schema-primary--default_rr_set_group--ds_record--values--sha1_digest--digest"></a>

### digest property

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

## Next pages

- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values.md)
- [xcsh_dns_zone](../data-sources/dns_zone.md)
