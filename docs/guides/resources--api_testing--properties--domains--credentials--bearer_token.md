---
page_title: "domains.credentials.bearer_token"
subcategory: ""
description: "domains.credentials.bearer_token for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 1253, "body_sha256": "sha256:a3e783b346024fc47a6634aae813e0095c4a8f8f8697df073c275054455e156b", "canonical_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token:token"], "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "parent_id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "path": "docs/guides/resources--api_testing--properties--domains--credentials--bearer_token.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["domains", "credentials", "bearer_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/bearer_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.credentials.bearer_token for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.bearer_token

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md)
- [Property reference](resources--api_testing--reference.md)
- [domains](resources--api_testing--properties--domains.md)
- [domains.credentials](resources--api_testing--properties--domains--credentials.md)
- domains.credentials.bearer_token

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bearer token.

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
bearer_token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [token](resources--api_testing--properties--domains--credentials--bearer_token--token.md): complete subsection reference.

## Next pages

- [domains.credentials.bearer_token.token](resources--api_testing--properties--domains--credentials--bearer_token--token.md)
- [domains.credentials](resources--api_testing--properties--domains--credentials.md)
- [xcsh_api_testing](../resources/api_testing.md)
