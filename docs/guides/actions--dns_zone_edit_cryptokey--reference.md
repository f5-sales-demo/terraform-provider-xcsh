---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_zone_edit_cryptokey."
xcsh_docs: {"aliases": [], "body_bytes": 1656, "body_sha256": "sha256:5cea1c1648424968b0081db96d6fbb52eea8f01f21d42912efd9e0d7fcf94df2", "canonical_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "child_ids": [], "collection_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:collection", "completeness": "complete", "id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "parent_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:fundamentals", "path": "docs/guides/actions--dns_zone_edit_cryptokey--reference.md", "provider_name": "dns_zone_edit_cryptokey", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_edit_cryptokey/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_dns_zone_edit_cryptokey.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_zone_edit_cryptokey](../actions/dns_zone_edit_cryptokey.md)
- Property reference

## Direct properties

<a id="schema-active"></a>

### active property

Type: `"bool"`. Optional.

Active. Indicates if the resource is active

<a id="schema-key_id"></a>

### key_id property

Type: `"number"`. Optional.

Key ID. Unique identifier for this resource

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
| `active` | [active](actions--dns_zone_edit_cryptokey--reference.md#schema-active) |
| `key_id` | [key_id](actions--dns_zone_edit_cryptokey--reference.md#schema-key_id) |
| `namespace` | [namespace](actions--dns_zone_edit_cryptokey--reference.md#schema-namespace) |
| `zone_name` | [zone_name](actions--dns_zone_edit_cryptokey--reference.md#schema-zone_name) |

## Next pages

- [xcsh_dns_zone_edit_cryptokey](../actions/dns_zone_edit_cryptokey.md)
