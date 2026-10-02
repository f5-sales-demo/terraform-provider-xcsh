---
page_title: "items.get_spec.infra.hw_info.board"
subcategory: ""
description: "Board Details. Board information."
xcsh_docs: {"aliases": ["items get spec infra hw info board"], "body_bytes": 2562, "body_sha256": "sha256:990a64fdd937c5eb0c912dd2d995720ded108a9f6309eae1304bf3e32d15ed36", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:board", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/board/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0020112033033321-3211111000213023-0122113303021132-3102101202022123-1020111131120330-2210031212330030-3320332233003221-0330002302331023", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "board"], "schema_version": 1, "sections": [{"aliases": ["asset tag"], "anchor": "schema-items--get_spec--infra--hw_info--board--asset_tag", "description": "Information from /sys/class/dmi/ID/board_asset_tag.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "board", "asset_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-items--get_spec--infra--hw_info--board--name", "description": "Information from /sys/class/dmi/ID/board_name.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "board", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["serial"], "anchor": "schema-items--get_spec--infra--hw_info--board--serial", "description": "Information from /sys/class/dmi/ID/board_serial.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "board", "serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["vendor"], "anchor": "schema-items--get_spec--infra--hw_info--board--vendor", "description": "Information from /sys/class/dmi/ID/board_vendor.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "board", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["version"], "anchor": "schema-items--get_spec--infra--hw_info--board--version", "description": "Information from /sys/class/dmi/ID/board_version.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "board", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/board/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Board Details. Board information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.board

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/)
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

- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
