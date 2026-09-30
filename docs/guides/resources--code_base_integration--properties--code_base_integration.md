---
page_title: "code_base_integration"
subcategory: ""
description: "code_base_integration for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 4541, "body_sha256": "sha256:408d6d87df435e35b55b98a27a0515d9dbe405f7818d58ee4bbfa0125267b466", "canonical_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration", "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration:azure_repos", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:github", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:github_enterprise", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:gitlab_enterprise"], "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration", "parent_id": "xcsh-docs:resources:code_base_integration:reference", "path": "docs/guides/resources--code_base_integration--properties--code_base_integration.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["code_base_integration"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "code_base_integration for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# code_base_integration

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md)
- [Property reference](resources--code_base_integration--reference.md)
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

- [azure_repos](resources--code_base_integration--properties--code_base_integration--azure_repos.md): complete subsection reference.

- [bitbucket](resources--code_base_integration--properties--code_base_integration--bitbucket.md): complete subsection reference.

- [bitbucket_server](resources--code_base_integration--properties--code_base_integration--bitbucket_server.md): complete subsection reference.

- [github](resources--code_base_integration--properties--code_base_integration--github.md): complete subsection reference.

- [github_enterprise](resources--code_base_integration--properties--code_base_integration--github_enterprise.md): complete subsection reference.

- [gitlab](resources--code_base_integration--properties--code_base_integration--gitlab.md): complete subsection reference.

- [gitlab_enterprise](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise.md): complete subsection reference.

## Next pages

- [code_base_integration.azure_repos](resources--code_base_integration--properties--code_base_integration--azure_repos.md)
- [code_base_integration.bitbucket](resources--code_base_integration--properties--code_base_integration--bitbucket.md)
- [code_base_integration.bitbucket_server](resources--code_base_integration--properties--code_base_integration--bitbucket_server.md)
- [code_base_integration.github](resources--code_base_integration--properties--code_base_integration--github.md)
- [code_base_integration.github_enterprise](resources--code_base_integration--properties--code_base_integration--github_enterprise.md)
- [code_base_integration.gitlab](resources--code_base_integration--properties--code_base_integration--gitlab.md)
- [code_base_integration.gitlab_enterprise](resources--code_base_integration--properties--code_base_integration--gitlab_enterprise.md)
- [Property reference](resources--code_base_integration--reference.md)
- [xcsh_code_base_integration](../resources/code_base_integration.md)
