---
page_title: "items.object.system_metadata.owner_view"
subcategory: ""
description: "ViewRefType represents a reference to a view."
xcsh_docs: {"aliases": ["items object system metadata owner view"], "body_bytes": 2647, "body_sha256": "sha256:261a2d88f42b1d2e99df2c13b3b15ccea5d0c5c4f5a64642e6f4c298ff010e94", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:owner_view", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/system_metadata/owner_view/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1332123201133031-1321012020133231-1031003123111130-0100330112202022-2310013002312223-3032000022033213-3000112333322212-2100122332322303", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "system_metadata", "owner_view"], "schema_version": 1, "sections": [{"aliases": ["items object system metadata owner view kind"], "anchor": "schema-items--object--system_metadata--owner_view--kind", "description": "Kind. Kind of the view object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:owner_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "owner_view", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata owner view name"], "anchor": "schema-items--object--system_metadata--owner_view--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:owner_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "owner_view", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata owner view namespace"], "anchor": "schema-items--object--system_metadata--owner_view--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:owner_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "owner_view", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata owner view uid"], "anchor": "schema-items--object--system_metadata--owner_view--uid", "description": "UID. UID of the view object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:owner_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "owner_view", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/system_metadata/owner_view/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "ViewRefType represents a reference to a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.system_metadata.owner_view

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/)
- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/)
- items.object.system_metadata.owner_view

<a id="section"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

## Direct properties

<a id="schema-items--object--system_metadata--owner_view--kind"></a>

### kind property

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="schema-items--object--system_metadata--owner_view--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--object--system_metadata--owner_view--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--object--system_metadata--owner_view--uid"></a>

### uid property

Type: `"string"`. Computed.

UID. UID of the view object.

## Next pages

- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
