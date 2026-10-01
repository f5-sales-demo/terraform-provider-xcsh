---
page_title: "primary.rr_set_group.rr_set.ds_record.values.sha1_digest"
subcategory: "DNS"
description: "primary.rr_set_group.rr_set.ds_record.values.sha1_digest for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 2412, "body_sha256": "sha256:ddd03fc958b7ae4c1abbc84691166a4b632dfedadaa44b0cd554c2cbddaa767b", "canonical_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:ds_record:values:sha1_digest", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:ds_record:values:sha1_digest", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:ds_record:values", "path": "docs/guides/data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values--sha1_digest.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "ds_record", "values", "sha1_digest"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/values/sha1_digest/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.rr_set_group.rr_set.ds_record.values.sha1_digest for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.ds_record.values.sha1_digest

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md)
- [Property reference](data-sources--dns_zone--reference.md)
- [primary](data-sources--dns_zone--properties--primary.md)
- [primary.rr_set_group](data-sources--dns_zone--properties--primary--rr_set_group.md)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--properties--primary--rr_set_group--rr_set.md)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record.md)
- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values.md)
- primary.rr_set_group.rr_set.ds_record.values.sha1_digest

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

<a id="schema-primary--rr_set_group--rr_set--ds_record--values--sha1_digest--digest"></a>

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

- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values.md)
- [xcsh_dns_zone](../data-sources/dns_zone.md)
