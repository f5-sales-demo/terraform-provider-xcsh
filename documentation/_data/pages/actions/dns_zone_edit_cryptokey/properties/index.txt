---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_zone_edit_cryptokey."
xcsh_docs: {"aliases": ["dns zone edit cryptokey"], "body_bytes": 1970, "body_sha256": "sha256:a429005e868b7f07b5ef5495fad135d2c58bf25bbf51fff949574b963bb4a777", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:collection", "completeness": "complete", "id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "parent_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:fundamentals", "path": "documentation/actions/dns_zone_edit_cryptokey/properties/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_edit_cryptokey", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "registry_anchor": "canonical-2030131320303020-2330130021202001-0203001123031021-3120120210311232-1132000311113113-2200003332022233-3230212003331112-2202112202232310", "registry_path": "docs/guides/actions--dns_zone_edit_cryptokey--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["active"], "anchor": "schema-active", "description": "Active. Indicates if the resource is active", "document_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["active"], "syntax": "attribute", "type": "bool"}, {"aliases": ["key id"], "anchor": "schema-key_id", "description": "Key ID. Unique identifier for this resource", "document_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["key_id"], "syntax": "attribute", "type": "number"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace is always system for dns_zone.", "document_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["zone name"], "anchor": "schema-zone_name", "description": "Zone Name. Human-readable name for the resource", "document_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["zone_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_edit_cryptokey/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_dns_zone_edit_cryptokey.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [xcsh_dns_zone_edit_cryptokey](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_edit_cryptokey/)
