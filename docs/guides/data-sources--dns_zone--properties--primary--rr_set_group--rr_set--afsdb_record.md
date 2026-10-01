---
page_title: "primary.rr_set_group.rr_set.afsdb_record"
subcategory: "DNS"
description: "primary.rr_set_group.rr_set.afsdb_record for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 2857, "body_sha256": "sha256:408cf88e2b85b0fcd4a7c2dd9707bbb08145b1538057645e08f860cefc1ef00e", "canonical_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record:values"], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set", "path": "docs/guides/data-sources--dns_zone--properties--primary--rr_set_group--rr_set--afsdb_record.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "afsdb_record"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.rr_set_group.rr_set.afsdb_record for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.afsdb_record

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md)
- [Property reference](data-sources--dns_zone--reference.md)
- [primary](data-sources--dns_zone--properties--primary.md)
- [primary.rr_set_group](data-sources--dns_zone--properties--primary--rr_set_group.md)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--properties--primary--rr_set_group--rr_set.md)
- primary.rr_set_group.rr_set.afsdb_record

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for afsdb record.

Upstream description:

DNS AFSDB Record.

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

<a id="schema-primary--rr_set_group--rr_set--afsdb_record--name"></a>

### name property

Type: `"string"`. Computed.

AFSDB Record name, please provide only the specific subdomain or record name without the base
domain.

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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--afsdb_record--values.md): complete subsection reference.

## Next pages

- [primary.rr_set_group.rr_set.afsdb_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--afsdb_record--values.md)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--properties--primary--rr_set_group--rr_set.md)
- [xcsh_dns_zone](../data-sources/dns_zone.md)
