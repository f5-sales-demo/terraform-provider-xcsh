---
page_title: "default_pool_list"
subcategory: "Load Balancing"
description: "default_pool_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1107, "body_sha256": "sha256:14281e333f6aed05a71276773fce00862ef40d39ced1e22f6cbe68ac253810ca", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool_list.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool_list

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- default_pool_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Origin Pool List Type. List of Origin Pools.

Upstream description:

List of Origin Pools.

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

Terraform syntax:

```terraform
default_pool_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [pools](resources--http_loadbalancer--properties--default_pool_list--pools.md): complete subsection reference.

## Next pages

- [default_pool_list.pools](resources--http_loadbalancer--properties--default_pool_list--pools.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
