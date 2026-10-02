---
page_title: "items.get_spec.infra.bond_config"
subcategory: ""
description: "Bond device configuration for VPM registration."
xcsh_docs: {"aliases": ["items get spec infra bond config"], "body_bytes": 2615, "body_sha256": "sha256:e30983b71920ecfa6e85ada7de43ce4c689d679db700137f288bf49a90cf1ad3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:bond_config", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra", "path": "documentation/data-sources/site_registrations_by_state/properties/items/get_spec/infra/bond_config/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1010031133200233-2032033102132113-0311123133113020-0330211122202233-1213021231220012-2111022032103301-0020131213033000-2013231120311120", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "bond_config"], "schema_version": 1, "sections": [{"aliases": ["interfaces"], "anchor": "schema-items--get_spec--infra--bond_config--interfaces", "description": "Member Interfaces. Configuration parameter for interfaces", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:bond_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "bond_config", "interfaces"], "syntax": "attribute", "type": "list"}, {"aliases": ["mode"], "anchor": "schema-items--get_spec--infra--bond_config--mode", "description": "Bonding mode for bond device configuration Bond mode is not specified Active-backup bond mode (one interface active, others as backup) IEEE 802.3ad Dynamic link aggregation (LACP). Possible values are `BOND_MODE_UNSPECIFIED`, `ACTIVE_BACKUP`, `LACP_802_3AD`. Defaults to `BOND_MODE_UNSPECIFIED`.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:bond_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "bond_config", "mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-items--get_spec--infra--bond_config--name", "description": "Bond Name. Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:bond_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "bond_config", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/get_spec/infra/bond_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Bond device configuration for VPM registration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.bond_config

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
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

- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
