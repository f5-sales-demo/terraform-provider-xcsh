---
page_title: "items.get_spec.infra.hw_info.product"
subcategory: ""
description: "items.get_spec.infra.hw_info.product for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 2184, "body_sha256": "sha256:3114b140fdd0afd0fba3cfde864103dc6f39580dbf6927e95bb6a3c28e575698", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:product", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:product", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info", "path": "docs/guides/data-sources--site_registrations_by_state--properties--items--get_spec--infra--hw_info--product.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "product"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/product/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.get_spec.infra.hw_info.product for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.product

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
- [Property reference](data-sources--site_registrations_by_state--reference.md)
- [items](data-sources--site_registrations_by_state--properties--items.md)
- [items.get_spec](data-sources--site_registrations_by_state--properties--items--get_spec.md)
- [items.get_spec.infra](data-sources--site_registrations_by_state--properties--items--get_spec--infra.md)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--properties--items--get_spec--infra--hw_info.md)
- items.get_spec.infra.hw_info.product

<a id="section"></a>

Type: `"single"`. Computed.

Product Information. Product information.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--product--name"></a>

### name property

Type: `"string"`. Computed.

Name. Product name, eg. For AWS m5a.xlarge. Info taken from /sys/class/dmi/ID/product\_name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--get_spec--infra--hw_info--product--serial"></a>

### serial property

Type: `"string"`. Computed.

Serial number, eg. For AWS 00000000-0000-4000-8000-23460f645a1f. Info taken from
/sys/class/dmi/ID/product\_serial.

<a id="schema-items--get_spec--infra--hw_info--product--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. Vendor name, eg. For AWS Amazon EC2. Info taken from /sys/class/dmi/ID/product\_vendor.

<a id="schema-items--get_spec--infra--hw_info--product--version"></a>

### version property

Type: `"string"`. Computed.

Version name. Info taken from /sys/class/dmi/ID/product\_version.

## Next pages

- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--properties--items--get_spec--infra--hw_info.md)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
