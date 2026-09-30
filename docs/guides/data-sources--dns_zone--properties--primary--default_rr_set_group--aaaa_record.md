---
page_title: "primary.default_rr_set_group.aaaa_record"
subcategory: "DNS"
description: "primary.default_rr_set_group.aaaa_record for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 3853, "body_sha256": "sha256:e75cf758aac13441a17dddd205fcb0e3040cb33279c74b45158dbcc5a8655274", "canonical_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:aaaa_record", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:aaaa_record", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group", "path": "docs/guides/data-sources--dns_zone--properties--primary--default_rr_set_group--aaaa_record.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "default_rr_set_group", "aaaa_record"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/default_rr_set_group/aaaa_record/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.default_rr_set_group.aaaa_record for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# primary.default_rr_set_group.aaaa_record

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md)
- [Property reference](data-sources--dns_zone--reference.md)
- [primary](data-sources--dns_zone--properties--primary.md)
- [primary.default_rr_set_group](data-sources--dns_zone--properties--primary--default_rr_set_group.md)
- primary.default_rr_set_group.aaaa_record

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for aaaa record.

Upstream description:

RecordSet for AAAA Records.

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

<a id="schema-primary--default_rr_set_group--aaaa_record--name"></a>

### name property

Type: `"string"`. Computed.

AAAA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="schema-primary--default_rr_set_group--aaaa_record--values"></a>

### values property

Type: `["list", "string"]`. Computed.

IPv6 Addresses. A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Upstream description:

A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [primary.default_rr_set_group](data-sources--dns_zone--properties--primary--default_rr_set_group.md)
- [xcsh_dns_zone](../data-sources/dns_zone.md)
