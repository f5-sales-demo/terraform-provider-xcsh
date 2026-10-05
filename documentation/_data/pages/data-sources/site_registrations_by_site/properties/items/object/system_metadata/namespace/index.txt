---
page_title: "items.object.system_metadata.namespace"
subcategory: ""
description: "The namespace this object belongs to. This is populated by the service based on the metadata.namespace field when an object is created."
xcsh_docs: {"aliases": ["items object system metadata namespace"], "body_bytes": 3099, "body_sha256": "sha256:508f77ed82b186854f1c5fa67927f50c23af2e5606c5d1602847c8f901f72776", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:namespace", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/system_metadata/namespace/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3110302010321032-0123102333010321-3201310113313103-0020331023020101-1022332022102301-2211210020003013-3323131121012133-3101312022132011", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "system_metadata", "namespace"], "schema_version": 1, "sections": [{"aliases": ["items object system metadata namespace kind"], "anchor": "schema-items--object--system_metadata--namespace--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route').", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:namespace", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "namespace", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata namespace name"], "anchor": "schema-items--object--system_metadata--namespace--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:namespace", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "namespace", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata namespace namespace"], "anchor": "schema-items--object--system_metadata--namespace--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:namespace", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "namespace", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata namespace tenant"], "anchor": "schema-items--object--system_metadata--namespace--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:namespace", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "namespace", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata namespace uid"], "anchor": "schema-items--object--system_metadata--namespace--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:namespace", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "namespace", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/system_metadata/namespace/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "The namespace this object belongs to. This is populated by the service based on the metadata.namespace field when an object is created.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.system_metadata.namespace

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/system_metadata/)
- items.object.system_metadata.namespace

<a id="section"></a>

Type: `"list"`. Computed.

The namespace this object belongs to. This is populated by the service based on the
metadata.namespace field when an object is created.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

## Direct properties

<a id="schema-items--object--system_metadata--namespace--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="schema-items--object--system_metadata--namespace--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-items--object--system_metadata--namespace--namespace"></a>

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

<a id="schema-items--object--system_metadata--namespace--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="schema-items--object--system_metadata--namespace--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

## Next pages

- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/system_metadata/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
