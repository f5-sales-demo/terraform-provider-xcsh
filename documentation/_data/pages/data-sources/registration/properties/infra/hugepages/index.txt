---
page_title: "infra.hugepages"
subcategory: ""
description: "Hugepage settings for CE on K8s SMV2 site."
xcsh_docs: {"aliases": ["infra hugepages"], "body_bytes": 2140, "body_sha256": "sha256:1d5a0b942805e43a4a681bdce6e1e2f62706d4b391051dea8f522a67d0e4932c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hugepages", "parent_id": "xcsh-docs:data-sources:registration:properties:infra", "path": "documentation/data-sources/registration/properties/infra/hugepages/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2120031103113301-3131220333132032-3213120323123113-0301303301112301-1202231231220101-2130323303003113-0133003202112021-0123331322323300", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hugepages"], "schema_version": 1, "sections": [{"aliases": ["free"], "anchor": "schema-infra--hugepages--free", "description": "Total number of free hugepages present.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hugepages", "free"], "syntax": "attribute", "type": "number"}, {"aliases": ["page size"], "anchor": "schema-infra--hugepages--page_size", "description": "Size of each hugepage.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hugepages", "page_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["total"], "anchor": "schema-infra--hugepages--total", "description": "Total number of hugepages present.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hugepages", "total"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hugepages/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Hugepage settings for CE on K8s SMV2 site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hugepages

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
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

- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
