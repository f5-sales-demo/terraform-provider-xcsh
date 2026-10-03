---
page_title: "items.get_spec.infra.hw_info.os"
subcategory: ""
description: "OS. Details of Operating System."
xcsh_docs: {"aliases": ["items get spec infra hw info os"], "body_bytes": 2341, "body_sha256": "sha256:1ca172dc14d170b22c4488c81bcf5244c89ad4c246b4c42478e556565d0b7f12", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:os", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/os/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3023000112321132-0233130212230003-0023300101223200-0310212213103222-3130021312211322-1320312300111321-0311030113101122-0031311110132233", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "os"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info os architecture"], "anchor": "schema-items--get_spec--infra--hw_info--os--architecture", "description": "Architecture. Architecture of OS.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "os", "architecture"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info os name"], "anchor": "schema-items--get_spec--infra--hw_info--os--name", "description": "Name. Name of OS.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "os", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info os release"], "anchor": "schema-items--get_spec--infra--hw_info--os--release", "description": "Release. Release of the OS.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "os", "release"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info os vendor"], "anchor": "schema-items--get_spec--infra--hw_info--os--vendor", "description": "Vendor. Vendor of OS.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "os", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info os version"], "anchor": "schema-items--get_spec--infra--hw_info--os--version", "description": "Version. Version of OS.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "os", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/os/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "OS. Details of Operating System.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.os

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.os

<a id="section"></a>

Type: `"single"`. Computed.

OS. Details of Operating System.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--os--architecture"></a>

### architecture property

Type: `"string"`. Computed.

Architecture. Architecture of OS.

<a id="schema-items--get_spec--infra--hw_info--os--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of OS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--get_spec--infra--hw_info--os--release"></a>

### release property

Type: `"string"`. Computed.

Release. Release of the OS.

<a id="schema-items--get_spec--infra--hw_info--os--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. Vendor of OS.

<a id="schema-items--get_spec--infra--hw_info--os--version"></a>

### version property

Type: `"string"`. Computed.

Version. Version of OS.

## Next pages

- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
