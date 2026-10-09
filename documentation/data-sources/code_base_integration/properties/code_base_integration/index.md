---
page_title: "code_base_integration"
subcategory: ""
description: "Choose your code base (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection details."
xcsh_docs: {"aliases": ["code base integration"], "body_bytes": 2241, "body_sha256": "sha256:60695b7ca0ea52b0a32779fea4da8b2d50890d7c00b310818b75d7ba325ed3fe", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:azure_repos", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket_server", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab_enterprise"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration", "parent_id": "xcsh-docs:data-sources:code_base_integration:reference", "path": "documentation/data-sources/code_base_integration/properties/code_base_integration/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122", "registry_path": "docs/guides/data-sources--code_base_integration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration"], "schema_version": 1, "sections": [{"aliases": ["code base integration azure repos"], "anchor": "section", "description": "Configuration parameter for azure repos.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:azure_repos", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "azure_repos"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration bitbucket"], "anchor": "section", "description": "BitBucket Cloud Integration.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "bitbucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration bitbucket server"], "anchor": "section", "description": "Configuration parameter for bitbucket server.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "bitbucket_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration github"], "anchor": "section", "description": "Github Integration.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "github"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration github enterprise"], "anchor": "section", "description": "Configuration parameter for github enterprise.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "github_enterprise"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration gitlab"], "anchor": "section", "description": "GitLab Cloud Integration.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "gitlab"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration gitlab enterprise"], "anchor": "section", "description": "Configuration parameter for gitlab enterprise.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab_enterprise", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "gitlab_enterprise"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/code_base_integration/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Choose your code base (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection details.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
