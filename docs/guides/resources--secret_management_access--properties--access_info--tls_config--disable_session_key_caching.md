---
page_title: "access_info.tls_config.disable_session_key_caching"
subcategory: ""
description: "access_info.tls_config.disable_session_key_caching for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 1287, "body_sha256": "sha256:79c84c2d110578dcecd012806cc0c5d9f4572173bb5d371e61f11ab6ef9275ff", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:disable_session_key_caching", "child_ids": [], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:disable_session_key_caching", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "path": "docs/guides/resources--secret_management_access--properties--access_info--tls_config--disable_session_key_caching.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "tls_config", "disable_session_key_caching"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/disable_session_key_caching/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.tls_config.disable_session_key_caching for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.disable_session_key_caching

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- [access_info.tls_config](resources--secret_management_access--properties--access_info--tls_config.md)
- access_info.tls_config.disable_session_key_caching

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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
disable_session_key_caching = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [access_info.tls_config](resources--secret_management_access--properties--access_info--tls_config.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
