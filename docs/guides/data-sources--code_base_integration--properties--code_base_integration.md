---
page_title: "code_base_integration"
subcategory: ""
description: "code_base_integration for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 2743, "body_sha256": "sha256:03e80cf34af84bc169ce8ee2b313d982f12ed6d3d925324d0d3b4f637a3cf79e", "canonical_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration", "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:azure_repos", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket_server", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab_enterprise"], "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration", "parent_id": "xcsh-docs:data-sources:code_base_integration:reference", "path": "docs/guides/data-sources--code_base_integration--properties--code_base_integration.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["code_base_integration"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/code_base_integration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "code_base_integration for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# code_base_integration

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md)
- [Property reference](data-sources--code_base_integration--reference.md)
- code_base_integration

<a id="section"></a>

Type: `"single"`. Computed.

Choose your code base (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection
details.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"azure_repos\",\"bitbucket\",\"bitbucket_server\",\"github\",\"github_enterprise\",\"gitlab\",\"gitlab_enterprise\"]"
}
```

## Direct properties

- [azure_repos](data-sources--code_base_integration--properties--code_base_integration--azure_repos.md): complete subsection reference.

- [bitbucket](data-sources--code_base_integration--properties--code_base_integration--bitbucket.md): complete subsection reference.

- [bitbucket_server](data-sources--code_base_integration--properties--code_base_integration--bitbucket_server.md): complete subsection reference.

- [github](data-sources--code_base_integration--properties--code_base_integration--github.md): complete subsection reference.

- [github_enterprise](data-sources--code_base_integration--properties--code_base_integration--github_enterprise.md): complete subsection reference.

- [gitlab](data-sources--code_base_integration--properties--code_base_integration--gitlab.md): complete subsection reference.

- [gitlab_enterprise](data-sources--code_base_integration--properties--code_base_integration--gitlab_enterprise.md): complete subsection reference.

## Next pages

- [code_base_integration.azure_repos](data-sources--code_base_integration--properties--code_base_integration--azure_repos.md)
- [code_base_integration.bitbucket](data-sources--code_base_integration--properties--code_base_integration--bitbucket.md)
- [code_base_integration.bitbucket_server](data-sources--code_base_integration--properties--code_base_integration--bitbucket_server.md)
- [code_base_integration.github](data-sources--code_base_integration--properties--code_base_integration--github.md)
- [code_base_integration.github_enterprise](data-sources--code_base_integration--properties--code_base_integration--github_enterprise.md)
- [code_base_integration.gitlab](data-sources--code_base_integration--properties--code_base_integration--gitlab.md)
- [code_base_integration.gitlab_enterprise](data-sources--code_base_integration--properties--code_base_integration--gitlab_enterprise.md)
- [Property reference](data-sources--code_base_integration--reference.md)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md)
