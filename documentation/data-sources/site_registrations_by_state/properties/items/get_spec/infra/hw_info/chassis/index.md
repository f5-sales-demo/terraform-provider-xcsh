---
page_title: "items.get_spec.infra.hw_info.chassis"
subcategory: ""
description: "Chassis Details. Chassis information."
xcsh_docs: {"aliases": ["items get spec infra hw info chassis"], "body_bytes": 2043, "body_sha256": "sha256:64b2dde80b07e25b507f06f60356c17074c1aceccd60814b4c3900eaa10bee3c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:chassis", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/chassis/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2010322021010012-0322303300203232-0200220212222113-1232321232010110-2210323310223202-2313220002300032-2303013122211331-1112322022122133", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "chassis"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info chassis asset tag"], "anchor": "schema-items--get_spec--infra--hw_info--chassis--asset_tag", "description": "Information from /sys/class/dmi/ID/chassis_asset_tag.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "chassis", "asset_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info chassis serial"], "anchor": "schema-items--get_spec--infra--hw_info--chassis--serial", "description": "Information from /sys/class/dmi/ID/chassis_serial.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "chassis", "serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info chassis type"], "anchor": "schema-items--get_spec--infra--hw_info--chassis--type", "description": "Information from /sys/class/dmi/ID/chassis_type.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "chassis", "type"], "syntax": "attribute", "type": "number"}, {"aliases": ["items get spec infra hw info chassis vendor"], "anchor": "schema-items--get_spec--infra--hw_info--chassis--vendor", "description": "Information from /sys/class/dmi/ID/chassis_vendor.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "chassis", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info chassis version"], "anchor": "schema-items--get_spec--infra--hw_info--chassis--version", "description": "Information from /sys/class/dmi/ID/chassis_version.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "chassis", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/chassis/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Chassis Details. Chassis information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.chassis

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.chassis

<a id="section"></a>

Type: `"single"`. Computed.

Chassis Details. Chassis information.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--chassis--asset_tag"></a>

### asset_tag property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_asset\_tag.

<a id="schema-items--get_spec--infra--hw_info--chassis--serial"></a>

### serial property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_serial.

<a id="schema-items--get_spec--infra--hw_info--chassis--type"></a>

### type property

Type: `"number"`. Computed.

Information from /sys/class/dmi/ID/chassis\_type.

<a id="schema-items--get_spec--infra--hw_info--chassis--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_vendor.

<a id="schema-items--get_spec--infra--hw_info--chassis--version"></a>

### version property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_version.
