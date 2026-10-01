---
page_title: "items.ref_value"
subcategory: ""
description: "items.ref_value for xcsh_bot_suggest_values."
xcsh_docs: {"aliases": [], "body_bytes": 1867, "body_sha256": "sha256:de92ff020e876778639a1c7c8da83cb4f0d0d4de8325cb2608662f8cb40ca44d", "canonical_id": "xcsh-docs:data-sources:bot_suggest_values:properties:items:ref_value", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_suggest_values:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_suggest_values:properties:items:ref_value", "parent_id": "xcsh-docs:data-sources:bot_suggest_values:properties:items", "path": "docs/guides/data-sources--bot_suggest_values--properties--items--ref_value.md", "provider_name": "bot_suggest_values", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "ref_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_suggest_values/properties/items/ref_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.ref_value for xcsh_bot_suggest_values.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.ref_value

Breadcrumbs:

- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md)
- [Property reference](data-sources--bot_suggest_values--reference.md)
- [items](data-sources--bot_suggest_values--properties--items.md)
- items.ref_value

<a id="section"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

## Direct properties

<a id="schema-items--ref_value--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

<a id="schema-items--ref_value--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

<a id="schema-items--ref_value--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

## Next pages

- [items](data-sources--bot_suggest_values--properties--items.md)
- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md)
