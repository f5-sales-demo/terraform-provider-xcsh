---
page_title: "items.object.spec.gc_spec.infra.hw_info.chassis"
subcategory: ""
description: "Chassis Details. Chassis information."
xcsh_docs: {"aliases": ["items object spec gc spec infra hw info chassis"], "body_bytes": 2724, "body_sha256": "sha256:6cf24a6ab10da991de11a1f5326d6b5bee6ab49d9febcb773e0a3062ae700680", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/chassis/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0111111022012310-1013313211331223-3312103121012233-3230331331100022-2102211302330002-1320130021203303-1213131223310203-2131012012112012", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis"], "schema_version": 1, "sections": [{"aliases": ["items object spec gc spec infra hw info chassis asset tag"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--chassis--asset_tag", "description": "Information from /sys/class/dmi/ID/chassis_asset_tag.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis", "asset_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra hw info chassis serial"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--chassis--serial", "description": "Information from /sys/class/dmi/ID/chassis_serial.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis", "serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra hw info chassis type"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--chassis--type", "description": "Information from /sys/class/dmi/ID/chassis_type.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis", "type"], "syntax": "attribute", "type": "number"}, {"aliases": ["items object spec gc spec infra hw info chassis vendor"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--chassis--vendor", "description": "Information from /sys/class/dmi/ID/chassis_vendor.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra hw info chassis version"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--chassis--version", "description": "Information from /sys/class/dmi/ID/chassis_version.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/chassis/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Chassis Details. Chassis information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info.chassis

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/)
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

- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
