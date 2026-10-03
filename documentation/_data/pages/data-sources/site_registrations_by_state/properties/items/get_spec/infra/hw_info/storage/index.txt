---
page_title: "items.get_spec.infra.hw_info.storage"
subcategory: ""
description: "Storage. List of storage devices in server."
xcsh_docs: {"aliases": ["items get spec infra hw info storage"], "body_bytes": 2627, "body_sha256": "sha256:b3074a23c597feef5b101d7f082ce33950661ac2a343c619146306eab5e5f59a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:storage", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/storage/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2121200210023133-3230202221331302-0330102230131100-1112231300032211-3313013102203030-3023101123020000-3013113003000002-3121000001131113", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "storage"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info storage driver"], "anchor": "schema-items--get_spec--infra--hw_info--storage--driver", "description": "Driver. Driver of device.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "driver"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info storage model"], "anchor": "schema-items--get_spec--infra--hw_info--storage--model", "description": "Model. Model of device.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "model"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info storage name"], "anchor": "schema-items--get_spec--infra--hw_info--storage--name", "description": "Name. Name of device, eg. Nvme0n1.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info storage serial"], "anchor": "schema-items--get_spec--infra--hw_info--storage--serial", "description": "Serial Number. Serial of device.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info storage size gb"], "anchor": "schema-items--get_spec--infra--hw_info--storage--size_gb", "description": "Size(GB). Device size in GB.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "size_gb"], "syntax": "attribute", "type": "number"}, {"aliases": ["items get spec infra hw info storage vendor"], "anchor": "schema-items--get_spec--infra--hw_info--storage--vendor", "description": "Vendor. Vendor of device.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "vendor"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/storage/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Storage. List of storage devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.storage

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.storage

<a id="section"></a>

Type: `"list"`. Computed.

Storage. List of storage devices in server.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--storage--driver"></a>

### driver property

Type: `"string"`. Computed.

Driver. Driver of device.

<a id="schema-items--get_spec--infra--hw_info--storage--model"></a>

### model property

Type: `"string"`. Computed.

Model. Model of device.

<a id="schema-items--get_spec--infra--hw_info--storage--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of device, eg. Nvme0n1.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--get_spec--infra--hw_info--storage--serial"></a>

### serial property

Type: `"string"`. Computed.

Serial Number. Serial of device.

<a id="schema-items--get_spec--infra--hw_info--storage--size_gb"></a>

### size_gb property

Type: `"number"`. Computed.

Size(GB). Device size in GB.

<a id="schema-items--get_spec--infra--hw_info--storage--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. Vendor of device.

## Next pages

- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
