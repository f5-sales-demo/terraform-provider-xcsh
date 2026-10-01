---
page_title: "infra.hugepages"
subcategory: ""
description: "infra.hugepages for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 1883, "body_sha256": "sha256:917046328bbf72aabfb86890e2279009fdde9b0ac6dcdca17b41a84185b3163f", "canonical_id": "xcsh-docs:data-sources:registration:properties:infra:hugepages", "child_ids": [], "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hugepages", "parent_id": "xcsh-docs:data-sources:registration:properties:infra", "path": "docs/guides/data-sources--registration--properties--infra--hugepages.md", "provider_name": "registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["infra", "hugepages"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hugepages/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "infra.hugepages for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hugepages

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md)
- [Property reference](data-sources--registration--reference.md)
- [infra](data-sources--registration--properties--infra.md)
- infra.hugepages

<a id="section"></a>

Type: `"list"`. Computed.

Hugepage settings for CE on K8s SMV2 site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-infra--hugepages--free"></a>

### free property

Type: `"number"`. Computed.

Free Hugepages. Total number of free hugepages present.

Upstream description:

Total number of free hugepages present.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hugepages--page_size"></a>

### page_size property

Type: `"number"`. Computed.

Hugepage Size. Size of each hugepage.

Upstream description:

Size of each hugepage.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hugepages--total"></a>

### total property

Type: `"number"`. Computed.

Total Hugepages. Total number of hugepages present.

Upstream description:

Total number of hugepages present.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [infra](data-sources--registration--properties--infra.md)
- [xcsh_registration](../data-sources/registration.md)
