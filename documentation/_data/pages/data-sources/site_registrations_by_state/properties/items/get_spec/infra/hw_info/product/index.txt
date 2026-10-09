---
page_title: "items.get_spec.infra.hw_info.product"
subcategory: ""
description: "Product Information. Product information."
xcsh_docs: {"aliases": ["items get spec infra hw info product"], "body_bytes": 2291, "body_sha256": "sha256:b05ac7539a780297208ea48968a7d80d71e8a8226c472a50bbbc54045eb4f1d5", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:product", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/product/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2112103311331133-0033002130121030-1233121233310203-3100323312013313-2321213001202221-1313121320130013-3300003103323023-2012010322113133", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "product"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info product name"], "anchor": "schema-items--get_spec--infra--hw_info--product--name", "description": "Name. Product name, eg. For AWS m5a.xlarge. Info taken from /sys/class/dmi/ID/product_name.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:product", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "product", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info product serial"], "anchor": "schema-items--get_spec--infra--hw_info--product--serial", "description": "Serial number, eg. For AWS 00000000-0000-4000-8000-23460f645a1f. Info taken from /sys/class/dmi/ID/product_serial.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:product", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "product", "serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info product vendor"], "anchor": "schema-items--get_spec--infra--hw_info--product--vendor", "description": "Vendor. Vendor name, eg. For AWS Amazon EC2. Info taken from /sys/class/dmi/ID/product_vendor.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:product", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "product", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info product version"], "anchor": "schema-items--get_spec--infra--hw_info--product--version", "description": "Version name. Info taken from /sys/class/dmi/ID/product_version.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:product", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "product", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/product/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Product Information. Product information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.product

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.product

<a id="section"></a>

Type: `"single"`. Computed.

Product Information. Product information.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--product--name"></a>

### name property

Type: `"string"`. Computed.

Name. Product name, eg. For AWS m5a.xlarge. Info taken from /sys/class/dmi/ID/product\_name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--get_spec--infra--hw_info--product--serial"></a>

### serial property

Type: `"string"`. Computed.

Serial number, eg. For AWS 00000000-0000-4000-8000-23460f645a1f. Info taken from
/sys/class/dmi/ID/product\_serial.

<a id="schema-items--get_spec--infra--hw_info--product--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. Vendor name, eg. For AWS Amazon EC2. Info taken from /sys/class/dmi/ID/product\_vendor.

<a id="schema-items--get_spec--infra--hw_info--product--version"></a>

### version property

Type: `"string"`. Computed.

Version name. Info taken from /sys/class/dmi/ID/product\_version.
