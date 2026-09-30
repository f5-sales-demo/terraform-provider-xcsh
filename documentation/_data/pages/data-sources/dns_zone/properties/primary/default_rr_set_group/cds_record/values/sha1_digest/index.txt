---
page_title: "primary.default_rr_set_group.cds_record.values.sha1_digest"
subcategory: "DNS"
description: "primary.default_rr_set_group.cds_record.values.sha1_digest for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 2642, "body_sha256": "sha256:410bcfbb5a8211e4714fa6b5ec372d0d147a0510111e3949e934c27e0f8c5a5b", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:cds_record:values:sha1_digest", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:cds_record:values", "path": "documentation/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha1_digest/index.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["primary", "default_rr_set_group", "cds_record", "values", "sha1_digest"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha1_digest/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.default_rr_set_group.cds_record.values.sha1_digest for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# primary.default_rr_set_group.cds_record.values.sha1_digest

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/)
- [primary.default_rr_set_group.cds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/)
- [primary.default_rr_set_group.cds_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/)
- primary.default_rr_set_group.cds_record.values.sha1_digest

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

<a id="schema-primary--default_rr_set_group--cds_record--values--sha1_digest--digest"></a>

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

- [primary.default_rr_set_group.cds_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
