---
page_title: "items.object.spec.gc_spec.infra.hw_info.os"
subcategory: ""
description: "items.object.spec.gc_spec.infra.hw_info.os for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 2293, "body_sha256": "sha256:17d1e412d14f36ef660997735f3429849d4a41835217c01d41eeee2bdacaa668", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:os", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:os", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info", "path": "docs/guides/data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--os.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "os"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/os/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.infra.hw_info.os for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# items.object.spec.gc_spec.infra.hw_info.os

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- [items](data-sources--site_registrations_by_site--properties--items.md)
- [items.object](data-sources--site_registrations_by_site--properties--items--object.md)
- [items.object.spec](data-sources--site_registrations_by_site--properties--items--object--spec.md)
- [items.object.spec.gc_spec](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec.md)
- [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra.md)
- [items.object.spec.gc_spec.infra.hw_info](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info.md)
- items.object.spec.gc_spec.infra.hw_info.os

<a id="section"></a>

Type: `"single"`. Computed.

OS. Details of Operating System.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--os--architecture"></a>

### architecture property

Type: `"string"`. Computed.

Architecture. Architecture of OS.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--os--name"></a>

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

<a id="schema-items--object--spec--gc_spec--infra--hw_info--os--release"></a>

### release property

Type: `"string"`. Computed.

Release. Release of the OS.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--os--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. Vendor of OS.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--os--version"></a>

### version property

Type: `"string"`. Computed.

Version. Version of OS.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info.md)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
