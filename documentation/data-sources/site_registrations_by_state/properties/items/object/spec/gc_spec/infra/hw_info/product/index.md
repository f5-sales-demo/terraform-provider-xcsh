---
page_title: "items.object.spec.gc_spec.infra.hw_info.product"
subcategory: ""
description: "items.object.spec.gc_spec.infra.hw_info.product for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 2938, "body_sha256": "sha256:a5afe0dc398f774173582959ed764b4d9e84d249c619ebf501744500060a86d8", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:product", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hw_info/product/index.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "product"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hw_info/product/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.infra.hw_info.product for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# items.object.spec.gc_spec.infra.hw_info.product

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hw_info/)
- items.object.spec.gc_spec.infra.hw_info.product

<a id="section"></a>

Type: `"single"`. Computed.

Product Information. Product information.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--product--name"></a>

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

<a id="schema-items--object--spec--gc_spec--infra--hw_info--product--serial"></a>

### serial property

Type: `"string"`. Computed.

Serial number, eg. For AWS 00000000-0000-4000-8000-23460f645a1f. Info taken from
/sys/class/dmi/ID/product\_serial.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--product--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. Vendor name, eg. For AWS Amazon EC2. Info taken from /sys/class/dmi/ID/product\_vendor.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--product--version"></a>

### version property

Type: `"string"`. Computed.

Version name. Info taken from /sys/class/dmi/ID/product\_version.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hw_info/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
