---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_zone_cryptokeys."
xcsh_docs: {"aliases": [], "body_bytes": 2890, "body_sha256": "sha256:2e685b8932ed6e7aba9640d426783e26c8e4df0a8ddca0f10fdf4e5dcbf28298", "child_ids": ["xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys"], "collection_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone_cryptokeys:reference", "parent_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:fundamentals", "path": "documentation/data-sources/dns_zone_cryptokeys/properties/index.md", "provider_name": "dns_zone_cryptokeys", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone_cryptokeys/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_dns_zone_cryptokeys.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/)
- Property reference

## Direct properties

- [keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/keys/): complete subsection reference.

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
| `keys` | [keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/keys/#section) |
| `keys.active` | [keys.active](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/keys/#schema-keys--active) |
| `keys.algorithm` | [keys.algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/keys/#schema-keys--algorithm) |
| `keys.dnskey` | [keys.dnskey](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/keys/#schema-keys--dnskey) |
| `keys.key_id` | [keys.key_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/keys/#schema-keys--key_id) |
| `keys.key_type` | [keys.key_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/keys/#schema-keys--key_type) |
| `keys.published` | [keys.published](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/keys/#schema-keys--published) |
| `keys.type` | [keys.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/keys/#schema-keys--type) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/#schema-namespace) |
| `zone_name` | [zone_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/#schema-zone_name) |

## Next pages

- [keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/keys/)
- [xcsh_dns_zone_cryptokeys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/)
