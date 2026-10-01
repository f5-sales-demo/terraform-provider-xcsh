---
page_title: "items.get_spec.infra.hw_info.storage"
subcategory: ""
description: "items.get_spec.infra.hw_info.storage for xcsh_site_registrations."
xcsh_docs: {"aliases": [], "body_bytes": 2537, "body_sha256": "sha256:94b95ece7867803ca9f627f067deb9aaaf134c232377453099c9bbe197e1c1ad", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:storage", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/storage/index.md", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "storage"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/storage/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.get_spec.infra.hw_info.storage for xcsh_site_registrations.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.storage

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.storage

<a id="section"></a>

Type: `"list"`. Computed.

Storage. List of storage devices in server.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--storage--driver"></a>

### driver property

Type: `"string"`. Computed.

Driver. Driver of device.

<a id="schema-items--get_spec--infra--hw_info--storage--model"></a>

### model property

Type: `"string"`. Computed.

Model. Model of device.

<a id="schema-items--get_spec--infra--hw_info--storage--name"></a>

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

<a id="schema-items--get_spec--infra--hw_info--storage--serial"></a>

### serial property

Type: `"string"`. Computed.

Serial Number. Serial of device.

<a id="schema-items--get_spec--infra--hw_info--storage--size_gb"></a>

### size_gb property

Type: `"number"`. Computed.

Size(GB). Device size in GB.

<a id="schema-items--get_spec--infra--hw_info--storage--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. Vendor of device.

## Next pages

- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
