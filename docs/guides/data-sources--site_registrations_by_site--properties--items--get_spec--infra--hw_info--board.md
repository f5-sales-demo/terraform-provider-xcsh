---
page_title: "items.get_spec.infra.hw_info.board"
subcategory: ""
description: "items.get_spec.infra.hw_info.board for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 2161, "body_sha256": "sha256:3d39fec7a98f5197aecbc11e961e9184c15ff809bf4ddafc719a32134bec840d", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:board", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:board", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info", "path": "docs/guides/data-sources--site_registrations_by_site--properties--items--get_spec--infra--hw_info--board.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "board"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/board/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.get_spec.infra.hw_info.board for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.board

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- [items](data-sources--site_registrations_by_site--properties--items.md)
- [items.get_spec](data-sources--site_registrations_by_site--properties--items--get_spec.md)
- [items.get_spec.infra](data-sources--site_registrations_by_site--properties--items--get_spec--infra.md)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_site--properties--items--get_spec--infra--hw_info.md)
- items.get_spec.infra.hw_info.board

<a id="section"></a>

Type: `"single"`. Computed.

Board Details. Board information.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--board--asset_tag"></a>

### asset_tag property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_asset\_tag.

<a id="schema-items--get_spec--infra--hw_info--board--name"></a>

### name property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--get_spec--infra--hw_info--board--serial"></a>

### serial property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_serial.

<a id="schema-items--get_spec--infra--hw_info--board--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_vendor.

<a id="schema-items--get_spec--infra--hw_info--board--version"></a>

### version property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_version.

## Next pages

- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_site--properties--items--get_spec--infra--hw_info.md)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
