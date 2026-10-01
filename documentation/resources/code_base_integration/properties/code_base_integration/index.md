---
page_title: "code_base_integration"
subcategory: ""
description: "code_base_integration for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 5548, "body_sha256": "sha256:b222187ee9c6c7fa0ed9039fca3d17b5205384f55cbc5902f15a5bc6f59d89da", "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration:azure_repos", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:github", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:github_enterprise", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise"], "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration", "parent_id": "xcsh-docs:resources:code_base_integration:reference", "path": "documentation/resources/code_base_integration/properties/code_base_integration/index.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["code_base_integration"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "code_base_integration for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/)
- code_base_integration

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Choose your code base (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection
details.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("azure_repos",
    "bitbucket"),
  validators.ConflictingObjectAttributes("azure_repos",
    "bitbucket_server"),
  validators.ConflictingObjectAttributes("azure_repos",
    "github"),
  validators.ConflictingObjectAttributes("azure_repos",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("azure_repos",
    "gitlab"),
  validators.ConflictingObjectAttributes("azure_repos",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket",
    "bitbucket_server"),
  validators.ConflictingObjectAttributes("bitbucket",
    "github"),
  validators.ConflictingObjectAttributes("bitbucket",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket",
    "gitlab"),
  validators.ConflictingObjectAttributes("bitbucket",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "github"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "gitlab"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("github",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("github",
    "gitlab"),
  validators.ConflictingObjectAttributes("github",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("github_enterprise",
    "gitlab"),
  validators.ConflictingObjectAttributes("github_enterprise",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("gitlab",
    "gitlab_enterprise")}
```

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

Terraform syntax:

```terraform
code_base_integration {
  # Configure direct properties listed below.
}
```

## Direct properties

- [azure_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/azure_repos/): complete subsection reference.

- [bitbucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket/): complete subsection reference.

- [bitbucket_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket_server/): complete subsection reference.

- [github](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/github/): complete subsection reference.

- [github_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/github_enterprise/): complete subsection reference.

- [gitlab](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/gitlab/): complete subsection reference.

- [gitlab_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/): complete subsection reference.

## Next pages

- [code_base_integration.azure_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/azure_repos/)
- [code_base_integration.bitbucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket/)
- [code_base_integration.bitbucket_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket_server/)
- [code_base_integration.github](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/github/)
- [code_base_integration.github_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/github_enterprise/)
- [code_base_integration.gitlab](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/gitlab/)
- [code_base_integration.gitlab_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/gitlab_enterprise/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
