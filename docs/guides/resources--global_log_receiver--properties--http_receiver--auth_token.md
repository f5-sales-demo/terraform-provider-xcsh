---
page_title: "http_receiver.auth_token"
subcategory: ""
description: "http_receiver.auth_token for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1257, "body_sha256": "sha256:d5d7ad03143caa804be1ec0159bace3abf62f5bafa7dcdf6c6f32a8d5af1e75a", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token:token"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver", "path": "docs/guides/resources--global_log_receiver--properties--http_receiver--auth_token.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_receiver", "auth_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/http_receiver/auth_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_receiver.auth_token for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver.auth_token

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [http_receiver](resources--global_log_receiver--properties--http_receiver.md)
- http_receiver.auth_token

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Access Token. Authentication Token for access.

Upstream description:

Authentication Token for access.

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
auth_token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [token](resources--global_log_receiver--properties--http_receiver--auth_token--token.md): complete subsection reference.

## Next pages

- [http_receiver.auth_token.token](resources--global_log_receiver--properties--http_receiver--auth_token--token.md)
- [http_receiver](resources--global_log_receiver--properties--http_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
