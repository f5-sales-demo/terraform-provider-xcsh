---
page_title: "items.object.spec.gc_spec.infra.hw_info.storage"
subcategory: ""
description: "items.object.spec.gc_spec.infra.hw_info.storage for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 2601, "body_sha256": "sha256:879fe6db6e540da6c65a118bd5c71f2227ce88d812ad222d36c2a3918fce1cc7", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:storage", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:storage", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info", "path": "docs/guides/data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--storage.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "storage"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/storage/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.infra.hw_info.storage for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info.storage

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- [items](data-sources--site_registrations_by_site--properties--items.md)
- [items.object](data-sources--site_registrations_by_site--properties--items--object.md)
- [items.object.spec](data-sources--site_registrations_by_site--properties--items--object--spec.md)
- [items.object.spec.gc_spec](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec.md)
- [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra.md)
- [items.object.spec.gc_spec.infra.hw_info](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info.md)
- items.object.spec.gc_spec.infra.hw_info.storage

<a id="section"></a>

Type: `"list"`. Computed.

Storage. List of storage devices in server.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--driver"></a>

### driver property

Type: `"string"`. Computed.

Driver. Driver of device.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--model"></a>

### model property

Type: `"string"`. Computed.

Model. Model of device.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--name"></a>

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

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--serial"></a>

### serial property

Type: `"string"`. Computed.

Serial Number. Serial of device.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--size_gb"></a>

### size_gb property

Type: `"number"`. Computed.

Size(GB). Device size in GB.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. Vendor of device.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info.md)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
