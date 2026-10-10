---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_zone_edit_cryptokey."
xcsh_docs: {"aliases": ["dns zone edit cryptokey"], "body_bytes": 1861, "body_sha256": "sha256:5bf22fd0dac1321e5adca5bd4664c806bc0983821d8ed150f16b4d5e8cad3397", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:collection", "completeness": "complete", "id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "parent_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:fundamentals", "path": "documentation/actions/dns_zone_edit_cryptokey/properties/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_edit_cryptokey", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "actions", "registry_anchor": "canonical-2030131320303020-2330130021202001-0203001123031021-3120120210311232-1132000311113113-2200003332022233-3230212003331112-2202112202232310", "registry_path": "docs/guides/actions--dns_zone_edit_cryptokey--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["active"], "anchor": "schema-active", "description": "Active. Indicates if the resource is active", "document_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["active"], "syntax": "attribute", "type": "bool"}, {"aliases": ["key id"], "anchor": "schema-key_id", "description": "Key ID. Unique identifier for this resource", "document_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["key_id"], "syntax": "attribute", "type": "number"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace is always system for dns_zone.", "document_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["zone name"], "anchor": "schema-zone_name", "description": "Zone Name. Human-readable name for the resource", "document_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["zone_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_edit_cryptokey/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_dns_zone_edit_cryptokey.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_zone_edit_cryptokey](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_edit_cryptokey/)
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
| `active` | [active](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_edit_cryptokey/properties/#schema-active) |
| `key_id` | [key_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_edit_cryptokey/properties/#schema-key_id) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_edit_cryptokey/properties/#schema-namespace) |
| `zone_name` | [zone_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_edit_cryptokey/properties/#schema-zone_name) |
