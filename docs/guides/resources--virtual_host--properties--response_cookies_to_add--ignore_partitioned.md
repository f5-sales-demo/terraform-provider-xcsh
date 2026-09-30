---
page_title: "response_cookies_to_add.ignore_partitioned"
subcategory: ""
description: "response_cookies_to_add.ignore_partitioned for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 991, "body_sha256": "sha256:d011e6da638ef914fee3aa908b2818bb5c3d4bc9a565d8c73dea6921d5801190", "canonical_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_partitioned", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_partitioned", "parent_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "path": "docs/guides/resources--virtual_host--properties--response_cookies_to_add--ignore_partitioned.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["response_cookies_to_add", "ignore_partitioned"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/response_cookies_to_add/ignore_partitioned/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "response_cookies_to_add.ignore_partitioned for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# response_cookies_to_add.ignore_partitioned

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [response_cookies_to_add](resources--virtual_host--properties--response_cookies_to_add.md)
- response_cookies_to_add.ignore_partitioned

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore partitioned.

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
ignore_partitioned = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [response_cookies_to_add](resources--virtual_host--properties--response_cookies_to_add.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
