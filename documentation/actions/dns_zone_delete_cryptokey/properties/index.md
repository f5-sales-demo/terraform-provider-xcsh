---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_zone_delete_cryptokey."
xcsh_docs: {"aliases": ["dns zone delete cryptokey"], "body_bytes": 1611, "body_sha256": "sha256:f8432140fd1a777241b28af41e575298ea7526f8dd55abcb2e9a67ffdda55e06", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:collection", "completeness": "complete", "id": "xcsh-docs:actions:dns_zone_delete_cryptokey:reference", "parent_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:fundamentals", "path": "documentation/actions/dns_zone_delete_cryptokey/properties/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_delete_cryptokey", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "actions", "registry_anchor": "canonical-1103032100322021-2001303320000110-2331221323002003-2030210003223133-1201102310333011-1330111231102102-1301022103211223-2321303012132222", "registry_path": "docs/guides/actions--dns_zone_delete_cryptokey--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["key id"], "anchor": "schema-key_id", "description": "Key ID. Unique identifier for this resource", "document_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["key_id"], "syntax": "attribute", "type": "number"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace is always system for dns_zone.", "document_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["zone name"], "anchor": "schema-zone_name", "description": "Zone Name. Human-readable name for the resource", "document_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["zone_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_delete_cryptokey/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_dns_zone_delete_cryptokey.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_zone_delete_cryptokey](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/)
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
| `key_id` | [key_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/properties/#schema-key_id) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/properties/#schema-namespace) |
| `zone_name` | [zone_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/properties/#schema-zone_name) |
