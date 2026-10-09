---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_zone_cryptokeys."
xcsh_docs: {"aliases": ["dns zone cryptokeys"], "body_bytes": 2767, "body_sha256": "sha256:bea8ee8747de9b6f62ab0137973ffafe1db5c83f0a6613fc140a19916c275fb8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone_cryptokeys:reference", "parent_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:fundamentals", "path": "documentation/data-sources/dns_zone_cryptokeys/properties/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_cryptokeys", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2123012120131000-0120323221212210-1230120031121100-3312302123120331-1103103020013131-0030222020311203-3311222211103303-2313101012023223", "registry_path": "docs/guides/data-sources--dns_zone_cryptokeys--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["keys"], "anchor": "section", "description": "Configuration parameter for keys", "document_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["keys"], "syntax": "attribute", "type": "object"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace is always system for dns_zone.", "document_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["zone name"], "anchor": "schema-zone_name", "description": "Zone Name. Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["zone_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone_cryptokeys/properties/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Property reference for xcsh_dns_zone_cryptokeys.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
EnumExtractionComplete: false
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
