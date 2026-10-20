---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_user_identification."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1130, "body_sha256": "sha256:e12bc224752b2b9ba7640e4c3098770d3b0cf244d5a3ca1b557203a2a07f1ef8", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a74347515a80c8e0aaaac40831664823cc6a3928018d3932b931b768606d6cb2", "source_path": "examples/resources/xcsh_user_identification/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:user_identification:example:resource", "parent_id": "xcsh-docs:resources:user_identification:examples", "path": "documentation/resources/user_identification/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0021030033310023-3323121131320120-2331223202320121-1111003123223123-2213223213002321-0121003220303213-3100001300010001-2122000223301302", "registry_path": "docs/guides/resources--user_identification--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/examples/resource/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Resource for xcsh_user_identification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["user_identificationCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/resource.tf`; digest `sha256:a74347515a80c8e0aaaac40831664823cc6a3928018d3932b931b768606d6cb2`.

```terraform
# UserIdentification Resource Example
# Manages user_identification creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UserIdentification configuration
resource "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}
```
