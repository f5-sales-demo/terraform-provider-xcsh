---
page_title: "items.object.spec.gc_spec.infra.hw_info.bios"
subcategory: ""
description: "items.object.spec.gc_spec.infra.hw_info.bios for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 2152, "body_sha256": "sha256:98ea80cd2988c001cc4f01dcd096a46a7dc132c053322c916babd5faa1836784", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:bios", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:bios", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info", "path": "docs/guides/data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--infra--hw_info--bios.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "bios"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hw_info/bios/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.infra.hw_info.bios for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info.bios

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
- [Property reference](data-sources--site_registrations_by_state--reference.md)
- [items](data-sources--site_registrations_by_state--properties--items.md)
- [items.object](data-sources--site_registrations_by_state--properties--items--object.md)
- [items.object.spec](data-sources--site_registrations_by_state--properties--items--object--spec.md)
- [items.object.spec.gc_spec](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec.md)
- [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--infra.md)
- [items.object.spec.gc_spec.infra.hw_info](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--infra--hw_info.md)
- items.object.spec.gc_spec.infra.hw_info.bios

<a id="section"></a>

Type: `"single"`. Computed.

Bios Data. BIOS information.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--bios--date"></a>

### date property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_date.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(10, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
    ""),
}
```

<a id="schema-items--object--spec--gc_spec--infra--hw_info--bios--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_vendor.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--bios--version"></a>

### version property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_version.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--infra--hw_info.md)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
