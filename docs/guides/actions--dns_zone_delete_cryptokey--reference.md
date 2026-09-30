---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_zone_delete_cryptokey."
xcsh_docs: {"aliases": [], "body_bytes": 1364, "body_sha256": "sha256:bd07cf3edee440860fb50fc0dc0caffa8010dfa338982ff23fd5461059219615", "canonical_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:reference", "child_ids": [], "collection_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:collection", "completeness": "complete", "id": "xcsh-docs:actions:dns_zone_delete_cryptokey:reference", "parent_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:fundamentals", "path": "docs/guides/actions--dns_zone_delete_cryptokey--reference.md", "provider_name": "dns_zone_delete_cryptokey", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_delete_cryptokey/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_dns_zone_delete_cryptokey.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_dns_zone_delete_cryptokey](../actions/dns_zone_delete_cryptokey.md)
- Property reference

## Direct properties

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
| `key_id` | [key_id](actions--dns_zone_delete_cryptokey--reference.md#schema-key_id) |
| `namespace` | [namespace](actions--dns_zone_delete_cryptokey--reference.md#schema-namespace) |
| `zone_name` | [zone_name](actions--dns_zone_delete_cryptokey--reference.md#schema-zone_name) |

## Next pages

- [xcsh_dns_zone_delete_cryptokey](../actions/dns_zone_delete_cryptokey.md)
