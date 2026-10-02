---
page_title: "items.object.spec.gc_spec.infra.hw_info.chassis"
subcategory: ""
description: "Chassis Details. Chassis information."
xcsh_docs: {"aliases": ["items object spec gc spec infra hw info chassis"], "body_bytes": 2820, "body_sha256": "sha256:56285b13ed6556e60f977443f1d07d352241773691a196b91869d92bdfd29a94", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/chassis/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3231131221120100-0312001230112331-2133321231023311-2032110132322312-2023113002303001-2301122130322210-0133110111321223-2303230121332010", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis"], "schema_version": 1, "sections": [{"aliases": ["asset tag"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--chassis--asset_tag", "description": "Information from /sys/class/dmi/ID/chassis_asset_tag.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis", "asset_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["serial"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--chassis--serial", "description": "Information from /sys/class/dmi/ID/chassis_serial.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis", "serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["type"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--chassis--type", "description": "Information from /sys/class/dmi/ID/chassis_type.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis", "type"], "syntax": "attribute", "type": "number"}, {"aliases": ["vendor"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--chassis--vendor", "description": "Information from /sys/class/dmi/ID/chassis_vendor.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["version"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--chassis--version", "description": "Information from /sys/class/dmi/ID/chassis_version.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/chassis/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Chassis Details. Chassis information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info.chassis

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/)
- items.object.spec.gc_spec.infra.hw_info.chassis

<a id="section"></a>

Type: `"single"`. Computed.

Chassis Details. Chassis information.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--chassis--asset_tag"></a>

### asset_tag property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_asset\_tag.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--chassis--serial"></a>

### serial property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_serial.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--chassis--type"></a>

### type property

Type: `"number"`. Computed.

Information from /sys/class/dmi/ID/chassis\_type.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--chassis--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_vendor.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--chassis--version"></a>

### version property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_version.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
