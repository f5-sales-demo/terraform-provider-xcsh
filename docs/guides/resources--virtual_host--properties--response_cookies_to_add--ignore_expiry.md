---
page_title: "response_cookies_to_add.ignore_expiry"
subcategory: ""
description: "response_cookies_to_add.ignore_expiry for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 971, "body_sha256": "sha256:8c95614620150eda4e69d19a9572b55d0bc17ec7a3fa8b8e3accae85cfeb9b95", "canonical_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_expiry", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_expiry", "parent_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "path": "docs/guides/resources--virtual_host--properties--response_cookies_to_add--ignore_expiry.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["response_cookies_to_add", "ignore_expiry"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/response_cookies_to_add/ignore_expiry/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "response_cookies_to_add.ignore_expiry for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# response_cookies_to_add.ignore_expiry

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [response_cookies_to_add](resources--virtual_host--properties--response_cookies_to_add.md)
- response_cookies_to_add.ignore_expiry

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore expiry.

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
ignore_expiry = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [response_cookies_to_add](resources--virtual_host--properties--response_cookies_to_add.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
