---
page_title: "items.get_spec.infra.bond_config"
subcategory: ""
description: "Bond device configuration for VPM registration."
xcsh_docs: {"aliases": ["items get spec infra bond config"], "body_bytes": 2534, "body_sha256": "sha256:135fa207e8de43a6d4885819d8ee74b72e17a00202de215f81f84a634e4fbf97", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:bond_config", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra", "path": "documentation/data-sources/site_registrations/properties/items/get_spec/infra/bond_config/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3302131310323123-2022321211032300-3121131322303332-1231332233301320-2231110010113320-0232323230310331-2222211002210211-1111113333211113", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "bond_config"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra bond config interfaces"], "anchor": "schema-items--get_spec--infra--bond_config--interfaces", "description": "Member Interfaces. Configuration parameter for interfaces", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:bond_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "bond_config", "interfaces"], "syntax": "attribute", "type": "list"}, {"aliases": ["items get spec infra bond config mode"], "anchor": "schema-items--get_spec--infra--bond_config--mode", "description": "Bonding mode for bond device configuration Bond mode is not specified Active-backup bond mode (one interface active, others as backup) IEEE 802.3ad Dynamic link aggregation (LACP). Possible values are `BOND_MODE_UNSPECIFIED`, `ACTIVE_BACKUP`, `LACP_802_3AD`. Defaults to `BOND_MODE_UNSPECIFIED`.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:bond_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "bond_config", "mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra bond config name"], "anchor": "schema-items--get_spec--infra--bond_config--name", "description": "Bond Name. Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:bond_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "bond_config", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/infra/bond_config/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Bond device configuration for VPM registration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.bond_config

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- items.get_spec.infra.bond_config

<a id="section"></a>

Type: `"single"`. Computed.

Bond device configuration for VPM registration.

## Direct properties

<a id="schema-items--get_spec--infra--bond_config--interfaces"></a>

### interfaces property

Type: `["list", "string"]`. Computed.

Member Interfaces. Configuration parameter for interfaces

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

<a id="schema-items--get_spec--infra--bond_config--mode"></a>

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

<a id="schema-items--get_spec--infra--bond_config--name"></a>

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

- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
