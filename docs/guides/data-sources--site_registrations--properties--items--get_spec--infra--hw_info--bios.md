---
page_title: "items.get_spec.infra.hw_info.bios"
subcategory: ""
description: "items.get_spec.infra.hw_info.bios for xcsh_site_registrations."
xcsh_docs: {"aliases": [], "body_bytes": 1615, "body_sha256": "sha256:af9f20fd01ac5a08b947b41e3b0987e659dad89d4a4852f9f304b823acfd25ef", "canonical_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:bios", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:bios", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info", "path": "docs/guides/data-sources--site_registrations--properties--items--get_spec--infra--hw_info--bios.md", "provider_name": "site_registrations", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "bios"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/bios/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.get_spec.infra.hw_info.bios for xcsh_site_registrations.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# items.get_spec.infra.hw_info.bios

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md)
- [Property reference](data-sources--site_registrations--reference.md)
- [items](data-sources--site_registrations--properties--items.md)
- [items.get_spec](data-sources--site_registrations--properties--items--get_spec.md)
- [items.get_spec.infra](data-sources--site_registrations--properties--items--get_spec--infra.md)
- [items.get_spec.infra.hw_info](data-sources--site_registrations--properties--items--get_spec--infra--hw_info.md)
- items.get_spec.infra.hw_info.bios

<a id="section"></a>

Type: `"single"`. Computed.

Bios Data. BIOS information.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--bios--date"></a>

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

<a id="schema-items--get_spec--infra--hw_info--bios--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_vendor.

<a id="schema-items--get_spec--infra--hw_info--bios--version"></a>

### version property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_version.

## Next pages

- [items.get_spec.infra.hw_info](data-sources--site_registrations--properties--items--get_spec--infra--hw_info.md)
- [xcsh_site_registrations](../data-sources/site_registrations.md)
