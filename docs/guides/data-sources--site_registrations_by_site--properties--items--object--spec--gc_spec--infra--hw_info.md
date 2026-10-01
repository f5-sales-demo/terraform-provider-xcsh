---
page_title: "items.object.spec.gc_spec.infra.hw_info"
subcategory: ""
description: "items.object.spec.gc_spec.infra.hw_info for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 5284, "body_sha256": "sha256:ec4b5e2e45cd76fbf105a1e6b2bf7f9fa32714749fd2c5fede323f0dc8bba813", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:bios", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:board", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:chassis", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:cpu", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:gpu", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:kernel", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:memory", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:network", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:os", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:product", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:storage", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:usb"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "path": "docs/guides/data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.infra.hw_info for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- [items](data-sources--site_registrations_by_site--properties--items.md)
- [items.object](data-sources--site_registrations_by_site--properties--items--object.md)
- [items.object.spec](data-sources--site_registrations_by_site--properties--items--object--spec.md)
- [items.object.spec.gc_spec](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec.md)
- [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra.md)
- items.object.spec.gc_spec.infra.hw_info

<a id="section"></a>

Type: `"single"`. Computed.

OsInfo holds information about host OS and HW.

## Direct properties

- [bios](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--bios.md): complete subsection reference.

- [board](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--board.md): complete subsection reference.

- [chassis](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--chassis.md): complete subsection reference.

- [cpu](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--cpu.md): complete subsection reference.

- [gpu](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--gpu.md): complete subsection reference.

- [kernel](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--kernel.md): complete subsection reference.

- [memory](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--memory.md): complete subsection reference.

- [network](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--network.md): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--numa_nodes"></a>

### numa_nodes property

Type: `"number"`. Computed.

Non-uniform memory access (NUMA) nodes count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

- [os](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--os.md): complete subsection reference.

- [product](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--product.md): complete subsection reference.

- [storage](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--storage.md): complete subsection reference.

- [usb](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--usb.md): complete subsection reference.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info.bios](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--bios.md)
- [items.object.spec.gc_spec.infra.hw_info.board](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--board.md)
- [items.object.spec.gc_spec.infra.hw_info.chassis](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--chassis.md)
- [items.object.spec.gc_spec.infra.hw_info.cpu](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--cpu.md)
- [items.object.spec.gc_spec.infra.hw_info.gpu](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--gpu.md)
- [items.object.spec.gc_spec.infra.hw_info.kernel](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--kernel.md)
- [items.object.spec.gc_spec.infra.hw_info.memory](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--memory.md)
- [items.object.spec.gc_spec.infra.hw_info.network](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--network.md)
- [items.object.spec.gc_spec.infra.hw_info.os](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--os.md)
- [items.object.spec.gc_spec.infra.hw_info.product](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--product.md)
- [items.object.spec.gc_spec.infra.hw_info.storage](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--storage.md)
- [items.object.spec.gc_spec.infra.hw_info.usb](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--hw_info--usb.md)
- [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra.md)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
