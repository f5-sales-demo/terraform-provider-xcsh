---
page_title: "items.object.spec.gc_spec.infra.bond_config"
subcategory: ""
description: "items.object.spec.gc_spec.infra.bond_config for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 2574, "body_sha256": "sha256:e987f960b2b9a3eca848f91447df6431499170b512a72f0022bce6ebeac33c70", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:bond_config", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:bond_config", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "path": "docs/guides/data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra--bond_config.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "bond_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/bond_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.infra.bond_config for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.bond_config

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- [items](data-sources--site_registrations_by_site--properties--items.md)
- [items.object](data-sources--site_registrations_by_site--properties--items--object.md)
- [items.object.spec](data-sources--site_registrations_by_site--properties--items--object--spec.md)
- [items.object.spec.gc_spec](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec.md)
- [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra.md)
- items.object.spec.gc_spec.infra.bond_config

<a id="section"></a>

Type: `"single"`. Computed.

Bond device configuration for VPM registration.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--bond_config--interfaces"></a>

### interfaces property

Type: `["list", "string"]`. Computed.

Member Interfaces. Configuration parameter for interfaces

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

<a id="schema-items--object--spec--gc_spec--infra--bond_config--mode"></a>

### mode property

Type: `"string"`. Computed.

\[Enum: BOND\_MODE\_UNSPECIFIED|ACTIVE\_BACKUP|LACP\_802\_3AD\] Bonding mode for bond device
configuration Bond mode is not specified Active-backup bond mode (one interface active, others as
backup) IEEE 802.3ad Dynamic link aggregation (LACP). Possible values are
\`BOND\_MODE\_UNSPECIFIED\`, \`ACTIVE\_BACKUP\`, \`LACP\_802\_3AD\`. Defaults to
\`BOND\_MODE\_UNSPECIFIED\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BOND_MODE_UNSPECIFIED",
    "ACTIVE_BACKUP",
    "LACP_802_3AD"),
}
```

<a id="schema-items--object--spec--gc_spec--infra--bond_config--name"></a>

### name property

Type: `"string"`. Computed.

Bond Name. Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

## Next pages

- [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_site--properties--items--object--spec--gc_spec--infra.md)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
