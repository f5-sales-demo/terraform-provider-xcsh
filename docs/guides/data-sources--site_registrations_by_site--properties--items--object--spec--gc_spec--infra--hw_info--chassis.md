---
page_title: "items.object.spec.gc_spec.infra.hw_info.chassis"
subcategory: ""
description: "items.object.spec.gc_spec.infra.hw_info.chassis for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 2229, "body_sha256": "sha256:538c7509955f69dac95065c258993ac0342e0219e03719d9f8d19bd55d584435", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info", "path": "docs/guides/data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--chassis.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "chassis"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/chassis/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.infra.hw_info.chassis for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# items.object.spec.gc_spec.infra.hw_info.chassis

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- [items](data-sources--site_registrations_by_site--properties--items.md)
- [items.object](data-sources--site_registrations_by_site--properties--items--object.md)
- [items.object.spec](data-sources--site_registrations_by_site--properties--items--object--spec.md)
- [items.object.spec.gc_spec](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec.md)
- [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra.md)
- [items.object.spec.gc_spec.infra.hw_info](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info.md)
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

- [items.object.spec.gc_spec.infra.hw_info](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info.md)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
