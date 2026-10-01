---
page_title: "primary.default_rr_set_group.alias_record"
subcategory: "DNS"
description: "primary.default_rr_set_group.alias_record for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 2173, "body_sha256": "sha256:2bb7ea5d26129c7e6d9fe6a835eacef14b865daa6cedf6821e2c40bea5acc99d", "canonical_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:alias_record", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:alias_record", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group", "path": "docs/guides/resources--dns_zone--properties--primary--default_rr_set_group--alias_record.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "default_rr_set_group", "alias_record"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_rr_set_group/alias_record/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.default_rr_set_group.alias_record for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.alias_record

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md)
- [Property reference](resources--dns_zone--reference.md)
- [primary](resources--dns_zone--properties--primary.md)
- [primary.default_rr_set_group](resources--dns_zone--properties--primary--default_rr_set_group.md)
- primary.default_rr_set_group.alias_record

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for alias record.

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
alias_record {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--default_rr_set_group--alias_record--value"></a>

### value property

Type: `"string"`. Optional.

Domain. A valid domain name, for example: example.com.

Upstream description:

A valid domain name, for example: example.com.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

## Next pages

- [primary.default_rr_set_group](resources--dns_zone--properties--primary--default_rr_set_group.md)
- [xcsh_dns_zone](../resources/dns_zone.md)
