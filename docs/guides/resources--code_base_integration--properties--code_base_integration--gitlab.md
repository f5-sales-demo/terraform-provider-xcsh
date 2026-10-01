---
page_title: "code_base_integration.gitlab"
subcategory: ""
description: "code_base_integration.gitlab for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 1273, "body_sha256": "sha256:72556d890f34fc0ecfbcc0f80b4842c5450e26abac1bac681db7208757cd865e", "canonical_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab", "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab:access_token"], "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab", "parent_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration", "path": "docs/guides/resources--code_base_integration--properties--code_base_integration--gitlab.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["code_base_integration", "gitlab"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/gitlab/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "code_base_integration.gitlab for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.gitlab

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md)
- [Property reference](resources--code_base_integration--reference.md)
- [code_base_integration](resources--code_base_integration--properties--code_base_integration.md)
- code_base_integration.gitlab

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

GitLab Cloud Integration.

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
gitlab {
  # Configure direct properties listed below.
}
```

## Direct properties

- [access_token](resources--code_base_integration--properties--code_base_integration--gitlab--access_token.md): complete subsection reference.

## Next pages

- [code_base_integration.gitlab.access_token](resources--code_base_integration--properties--code_base_integration--gitlab--access_token.md)
- [code_base_integration](resources--code_base_integration--properties--code_base_integration.md)
- [xcsh_code_base_integration](../resources/code_base_integration.md)
