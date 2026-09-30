---
page_title: "items.object.system_metadata.namespace"
subcategory: ""
description: "items.object.system_metadata.namespace for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 2646, "body_sha256": "sha256:9e6d61b49e1787ba23e0ccf383f869391fe4de8ece2b894c5f1a7060cca76a49", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:namespace", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:namespace", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata", "path": "docs/guides/data-sources--site_registrations_by_site--properties--items--object--system_metadata--namespace.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "system_metadata", "namespace"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/system_metadata/namespace/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.system_metadata.namespace for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# items.object.system_metadata.namespace

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- [items](data-sources--site_registrations_by_site--properties--items.md)
- [items.object](data-sources--site_registrations_by_site--properties--items--object.md)
- [items.object.system_metadata](data-sources--site_registrations_by_site--properties--items--object--system_metadata.md)
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

- [items.object.system_metadata](data-sources--site_registrations_by_site--properties--items--object--system_metadata.md)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
