---
page_title: "items.get_spec.infra.hw_info.storage"
subcategory: ""
description: "Storage. List of storage devices in server."
xcsh_docs: {"aliases": ["items get spec infra hw info storage"], "body_bytes": 2296, "body_sha256": "sha256:9640c616f2b50770d96374ec0e600cd64f3551da19feb8bb07b82ae9e969b4b4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:storage", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/storage/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2322332212010010-0120121232200212-3223133321113021-1310210213110021-0131122121102212-3321020331231022-1233123301313231-3230113002312220", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "storage"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info storage driver"], "anchor": "schema-items--get_spec--infra--hw_info--storage--driver", "description": "Driver. Driver of device.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "driver"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info storage model"], "anchor": "schema-items--get_spec--infra--hw_info--storage--model", "description": "Model. Model of device.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "model"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info storage name"], "anchor": "schema-items--get_spec--infra--hw_info--storage--name", "description": "Name. Name of device, eg. Nvme0n1.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info storage serial"], "anchor": "schema-items--get_spec--infra--hw_info--storage--serial", "description": "Serial Number. Serial of device.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info storage size gb"], "anchor": "schema-items--get_spec--infra--hw_info--storage--size_gb", "description": "Size(GB). Device size in GB.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "size_gb"], "syntax": "attribute", "type": "number"}, {"aliases": ["items get spec infra hw info storage vendor"], "anchor": "schema-items--get_spec--infra--hw_info--storage--vendor", "description": "Vendor. Vendor of device.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage", "vendor"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/storage/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Storage. List of storage devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.storage

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/)
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
