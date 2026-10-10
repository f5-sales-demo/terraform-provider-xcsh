---
page_title: "items.get_spec.infra.hw_info.bios"
subcategory: ""
description: "Bios Data. BIOS information."
xcsh_docs: {"aliases": ["items get spec infra hw info bios"], "body_bytes": 1904, "body_sha256": "sha256:a3159356c59bc4de1285e280d6f849c5f580d885acc4c7bcd69c04f3d6644423", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:bios", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/bios/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3001203323000323-2222212331201210-1002200302300132-3323010212233333-3020100232130321-0221331120202231-3023032031321331-3231223201122120", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "bios"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info bios date"], "anchor": "schema-items--get_spec--infra--hw_info--bios--date", "description": "Information from /sys/class/dmi/ID/bios_date.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:bios", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "bios", "date"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info bios vendor"], "anchor": "schema-items--get_spec--infra--hw_info--bios--vendor", "description": "Information from /sys/class/dmi/ID/bios_vendor.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:bios", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "bios", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info bios version"], "anchor": "schema-items--get_spec--infra--hw_info--bios--version", "description": "Information from /sys/class/dmi/ID/bios_version.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:bios", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "bios", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/bios/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Bios Data. BIOS information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.bios

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.bios

<a id="section"></a>

Type: `"single"`. Computed.

Bios Data. BIOS information.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--bios--date"></a>

### date property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_date.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(10, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
    ""),
}
```

<a id="schema-items--get_spec--infra--hw_info--bios--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_vendor.

<a id="schema-items--get_spec--infra--hw_info--bios--version"></a>

### version property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_version.
