---
page_title: "items.get_spec.infra.hw_info.bios"
subcategory: ""
description: "Bios Data. BIOS information."
xcsh_docs: {"aliases": ["items get spec infra hw info bios"], "body_bytes": 2115, "body_sha256": "sha256:1b63d4e1bb78dd7ccd1f4d0d03a3fd932359665e9ee53af939494c5fbda0c375", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:bios", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/bios/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1001212023032321-0220330131020201-0231012221322031-3003031031332302-1030333311233310-3331012211020112-2020220103031322-0101000121102033", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "bios"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info bios date"], "anchor": "schema-items--get_spec--infra--hw_info--bios--date", "description": "Information from /sys/class/dmi/ID/bios_date.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:bios", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "bios", "date"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info bios vendor"], "anchor": "schema-items--get_spec--infra--hw_info--bios--vendor", "description": "Information from /sys/class/dmi/ID/bios_vendor.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:bios", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "bios", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info bios version"], "anchor": "schema-items--get_spec--infra--hw_info--bios--version", "description": "Information from /sys/class/dmi/ID/bios_version.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:bios", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "bios", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/bios/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Bios Data. BIOS information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.bios

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
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

## Next pages

- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
