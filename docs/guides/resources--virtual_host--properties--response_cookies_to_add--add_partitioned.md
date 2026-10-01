---
page_title: "response_cookies_to_add.add_partitioned"
subcategory: ""
description: "response_cookies_to_add.add_partitioned for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1078, "body_sha256": "sha256:6feff402126d17dd0ef4264b4e243c85428d0af8ec24228a12f7ca49f11b902d", "canonical_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:add_partitioned", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:add_partitioned", "parent_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "path": "docs/guides/resources--virtual_host--properties--response_cookies_to_add--add_partitioned.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["response_cookies_to_add", "add_partitioned"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/response_cookies_to_add/add_partitioned/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "response_cookies_to_add.add_partitioned for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_cookies_to_add.add_partitioned

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [response_cookies_to_add](resources--virtual_host--properties--response_cookies_to_add.md)
- response_cookies_to_add.add_partitioned

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add partitioned.

Upstream description:

This can be used for messages where no values are needed.

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
add_partitioned = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [response_cookies_to_add](resources--virtual_host--properties--response_cookies_to_add.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
