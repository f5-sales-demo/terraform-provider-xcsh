---
page_title: "primary.rr_set_group.rr_set.mx_record.values"
subcategory: "DNS"
description: "primary.rr_set_group.rr_set.mx_record.values for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 3695, "body_sha256": "sha256:f85cade18ab03f063b5ca8c6806d5d09f4a5677f0cdd0e423c9bd73511284153", "canonical_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:mx_record:values", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:mx_record:values", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:mx_record", "path": "docs/guides/data-sources--dns_zone--properties--primary--rr_set_group--rr_set--mx_record--values.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "mx_record", "values"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/mx_record/values/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.rr_set_group.rr_set.mx_record.values for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.mx_record.values

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md)
- [Property reference](data-sources--dns_zone--reference.md)
- [primary](data-sources--dns_zone--properties--primary.md)
- [primary.rr_set_group](data-sources--dns_zone--properties--primary--rr_set_group.md)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--properties--primary--rr_set_group--rr_set.md)
- [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--mx_record.md)
- primary.rr_set_group.rr_set.mx_record.values

<a id="section"></a>

Type: `"list"`. Computed.

MX Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

## Direct properties

<a id="schema-primary--rr_set_group--rr_set--mx_record--values--domain"></a>

### domain property

Type: `"string"`. Computed.

Mail exchanger domain name, please provide the full hostname, for.

Upstream description:

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="schema-primary--rr_set_group--rr_set--mx_record--values--priority"></a>

### priority property

Type: `"number"`. Computed.

Priority. Mail exchanger priority code.

Upstream description:

Mail exchanger priority code.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--mx_record.md)
- [xcsh_dns_zone](../data-sources/dns_zone.md)
