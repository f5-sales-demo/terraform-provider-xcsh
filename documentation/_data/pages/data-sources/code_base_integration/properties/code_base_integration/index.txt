---
page_title: "code_base_integration"
subcategory: ""
description: "Choose your code base (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection details."
xcsh_docs: {"aliases": ["code base integration"], "body_bytes": 3750, "body_sha256": "sha256:4fc0fa0cc1c4486aa30aecf35a98dccb21add644a9811aa34fa0e7841f054d69", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:azure_repos", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket_server", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab_enterprise"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration", "parent_id": "xcsh-docs:data-sources:code_base_integration:reference", "path": "documentation/data-sources/code_base_integration/properties/code_base_integration/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122", "registry_path": "docs/guides/data-sources--code_base_integration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration"], "schema_version": 1, "sections": [{"aliases": ["code base integration azure repos"], "anchor": "section", "description": "Configuration parameter for azure repos.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:azure_repos", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "azure_repos"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration bitbucket"], "anchor": "section", "description": "BitBucket Cloud Integration.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "bitbucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration bitbucket server"], "anchor": "section", "description": "Configuration parameter for bitbucket server.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "bitbucket_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration github"], "anchor": "section", "description": "Github Integration.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "github"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration github enterprise"], "anchor": "section", "description": "Configuration parameter for github enterprise.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "github_enterprise"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration gitlab"], "anchor": "section", "description": "GitLab Cloud Integration.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "gitlab"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration gitlab enterprise"], "anchor": "section", "description": "Configuration parameter for gitlab enterprise.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab_enterprise", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "gitlab_enterprise"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/code_base_integration/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Choose your code base (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection details.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/)
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

- [azure_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/azure_repos/): complete subsection reference.

- [bitbucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/): complete subsection reference.

- [bitbucket_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/): complete subsection reference.

- [github](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/): complete subsection reference.

- [github_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/): complete subsection reference.

- [gitlab](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab/): complete subsection reference.

- [gitlab_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/): complete subsection reference.

## Next pages

- [code_base_integration.azure_repos](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/azure_repos/)
- [code_base_integration.bitbucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/)
- [code_base_integration.bitbucket_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/)
- [code_base_integration.github](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/)
- [code_base_integration.github_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/)
- [code_base_integration.gitlab](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab/)
- [code_base_integration.gitlab_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
