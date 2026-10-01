---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_zone_cryptokeys."
xcsh_docs: {"aliases": [], "body_bytes": 2271, "body_sha256": "sha256:87d93ada34137af03ae0f587deea7572ecf61de261ca2caa0aa7dc0343598b2e", "canonical_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:reference", "child_ids": ["xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys"], "collection_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone_cryptokeys:reference", "parent_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:fundamentals", "path": "docs/guides/data-sources--dns_zone_cryptokeys--reference.md", "provider_name": "dns_zone_cryptokeys", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone_cryptokeys/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_dns_zone_cryptokeys.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md)
- Property reference

## Direct properties

- [keys](data-sources--dns_zone_cryptokeys--properties--keys.md): complete subsection reference.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Optional.

Namespace is always system for dns\_zone.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-zone_name"></a>

### zone_name property

Type: `"string"`. Optional.

Zone Name. Human-readable name for the resource

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `keys` | [keys](data-sources--dns_zone_cryptokeys--properties--keys.md#section) |
| `keys.active` | [keys.active](data-sources--dns_zone_cryptokeys--properties--keys.md#schema-keys--active) |
| `keys.algorithm` | [keys.algorithm](data-sources--dns_zone_cryptokeys--properties--keys.md#schema-keys--algorithm) |
| `keys.dnskey` | [keys.dnskey](data-sources--dns_zone_cryptokeys--properties--keys.md#schema-keys--dnskey) |
| `keys.key_id` | [keys.key_id](data-sources--dns_zone_cryptokeys--properties--keys.md#schema-keys--key_id) |
| `keys.key_type` | [keys.key_type](data-sources--dns_zone_cryptokeys--properties--keys.md#schema-keys--key_type) |
| `keys.published` | [keys.published](data-sources--dns_zone_cryptokeys--properties--keys.md#schema-keys--published) |
| `keys.type` | [keys.type](data-sources--dns_zone_cryptokeys--properties--keys.md#schema-keys--type) |
| `namespace` | [namespace](data-sources--dns_zone_cryptokeys--reference.md#schema-namespace) |
| `zone_name` | [zone_name](data-sources--dns_zone_cryptokeys--reference.md#schema-zone_name) |

## Next pages

- [keys](data-sources--dns_zone_cryptokeys--properties--keys.md)
- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md)
