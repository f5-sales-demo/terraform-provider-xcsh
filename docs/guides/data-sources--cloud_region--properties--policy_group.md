---
page_title: "policy_group"
subcategory: ""
description: "policy_group for xcsh_cloud_region."
xcsh_docs: {"aliases": [], "body_bytes": 1332, "body_sha256": "sha256:5b84eee07d995b412f6c38a52f9c50072f7784d6246237ef9069bc86294489a8", "canonical_id": "xcsh-docs:data-sources:cloud_region:properties:policy_group", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cloud_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_region:properties:policy_group", "parent_id": "xcsh-docs:data-sources:cloud_region:reference", "path": "docs/guides/data-sources--cloud_region--properties--policy_group.md", "provider_name": "cloud_region", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_group"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_region/properties/policy_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_group for xcsh_cloud_region.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_group

Breadcrumbs:

- [xcsh_cloud_region](../data-sources/cloud_region.md)
- [Property reference](data-sources--cloud_region--reference.md)
- policy_group

<a id="section"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

## Direct properties

<a id="schema-policy_group--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-policy_group--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-policy_group--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

## Next pages

- [Property reference](data-sources--cloud_region--reference.md)
- [xcsh_cloud_region](../data-sources/cloud_region.md)
