---
page_title: "items.get_spec.infra.hw_info"
subcategory: ""
description: "items.get_spec.infra.hw_info for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 6059, "body_sha256": "sha256:1882dc5f2f03f03de8b5114ddf508d26b418bc70600d680c43512ec83ec64f34", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:bios", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:board", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:chassis", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:cpu", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:gpu", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:kernel", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:memory", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:network", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:os", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:product", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:storage", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:usb"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra", "path": "documentation/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/index.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.get_spec.infra.hw_info for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
- items.get_spec.infra.hw_info

<a id="section"></a>

Type: `"single"`. Computed.

OsInfo holds information about host OS and HW.

## Direct properties

- [bios](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/bios/): complete subsection reference.

- [board](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/board/): complete subsection reference.

- [chassis](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/chassis/): complete subsection reference.

- [cpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/cpu/): complete subsection reference.

- [gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/gpu/): complete subsection reference.

- [kernel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/kernel/): complete subsection reference.

- [memory](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/memory/): complete subsection reference.

- [network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/network/): complete subsection reference.

<a id="schema-items--get_spec--infra--hw_info--numa_nodes"></a>

### numa_nodes property

Type: `"number"`. Computed.

Non-uniform memory access (NUMA) nodes count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/os/): complete subsection reference.

- [product](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/product/): complete subsection reference.

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/storage/): complete subsection reference.

- [usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/usb/): complete subsection reference.

## Next pages

- [items.get_spec.infra.hw_info.bios](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/bios/)
- [items.get_spec.infra.hw_info.board](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/board/)
- [items.get_spec.infra.hw_info.chassis](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/chassis/)
- [items.get_spec.infra.hw_info.cpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/cpu/)
- [items.get_spec.infra.hw_info.gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/gpu/)
- [items.get_spec.infra.hw_info.kernel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/kernel/)
- [items.get_spec.infra.hw_info.memory](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/memory/)
- [items.get_spec.infra.hw_info.network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/network/)
- [items.get_spec.infra.hw_info.os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/os/)
- [items.get_spec.infra.hw_info.product](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/product/)
- [items.get_spec.infra.hw_info.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/storage/)
- [items.get_spec.infra.hw_info.usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/usb/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
