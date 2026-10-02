---
page_title: "items.system_metadata.owner_view"
subcategory: ""
description: "ViewRefType represents a reference to a view."
xcsh_docs: {"aliases": ["items system metadata owner view"], "body_bytes": 2425, "body_sha256": "sha256:91cae8ae41797d0becf0e3b04d2b37520abc93aec14a4fe6aaa89581e8759886", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:system_metadata:owner_view", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:system_metadata", "path": "documentation/data-sources/site_registrations_by_site/properties/items/system_metadata/owner_view/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-0131111212003031-0320232030112333-0331230220122203-2133101121301011-1121230312102310-2302113330020300-3133000132212130-2103020320312211", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "system_metadata", "owner_view"], "schema_version": 1, "sections": [{"aliases": ["kind"], "anchor": "schema-items--system_metadata--owner_view--kind", "description": "Kind. Kind of the view object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:system_metadata:owner_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "owner_view", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-items--system_metadata--owner_view--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:system_metadata:owner_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "owner_view", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-items--system_metadata--owner_view--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:system_metadata:owner_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "owner_view", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["uid"], "anchor": "schema-items--system_metadata--owner_view--uid", "description": "UID. UID of the view object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:system_metadata:owner_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "owner_view", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/system_metadata/owner_view/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "ViewRefType represents a reference to a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.system_metadata.owner_view

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/system_metadata/)
- items.system_metadata.owner_view

<a id="section"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

## Direct properties

<a id="schema-items--system_metadata--owner_view--kind"></a>

### kind property

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="schema-items--system_metadata--owner_view--name"></a>

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

<a id="schema-items--system_metadata--owner_view--namespace"></a>

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

<a id="schema-items--system_metadata--owner_view--uid"></a>

### uid property

Type: `"string"`. Computed.

UID. UID of the view object.

## Next pages

- [items.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/system_metadata/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
