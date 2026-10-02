---
page_title: "items.get_spec.infra.hw_info.board"
subcategory: ""
description: "Board Details. Board information."
xcsh_docs: {"aliases": ["items get spec infra hw info board"], "body_bytes": 2482, "body_sha256": "sha256:8beaa71f6c08bf2ddc084f2cee8f7b70f8a877fbd362d919c8bf496f7e752901", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:board", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/board/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2122203023121032-3302323200211131-3300023032010113-3010213213112031-1130130322322210-1331122132110300-3231330021133101-2001033130112031", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "board"], "schema_version": 1, "sections": [{"aliases": ["asset tag"], "anchor": "schema-items--get_spec--infra--hw_info--board--asset_tag", "description": "Information from /sys/class/dmi/ID/board_asset_tag.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "board", "asset_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-items--get_spec--infra--hw_info--board--name", "description": "Information from /sys/class/dmi/ID/board_name.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "board", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["serial"], "anchor": "schema-items--get_spec--infra--hw_info--board--serial", "description": "Information from /sys/class/dmi/ID/board_serial.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "board", "serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["vendor"], "anchor": "schema-items--get_spec--infra--hw_info--board--vendor", "description": "Information from /sys/class/dmi/ID/board_vendor.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "board", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["version"], "anchor": "schema-items--get_spec--infra--hw_info--board--version", "description": "Information from /sys/class/dmi/ID/board_version.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "board", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/board/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Board Details. Board information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.board

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.board

<a id="section"></a>

Type: `"single"`. Computed.

Board Details. Board information.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--board--asset_tag"></a>

### asset_tag property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_asset\_tag.

<a id="schema-items--get_spec--infra--hw_info--board--name"></a>

### name property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--get_spec--infra--hw_info--board--serial"></a>

### serial property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_serial.

<a id="schema-items--get_spec--infra--hw_info--board--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_vendor.

<a id="schema-items--get_spec--infra--hw_info--board--version"></a>

### version property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_version.

## Next pages

- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
