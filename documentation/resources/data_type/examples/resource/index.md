---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_data_type."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1040, "body_sha256": "sha256:64cda08206c6122365bc515e52d6c60bdb446b018626247f389f417db671e4df", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:292681f3b6c18ec2ffbe2be1a0e17beb9c57b44e1777befb42567a6c3a8f12ee", "source_path": "examples/resources/xcsh_data_type/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:data_type:example:resource", "parent_id": "xcsh-docs:resources:data_type:examples", "path": "documentation/resources/data_type/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3123010221000010-1300110032110002-0020113210323203-3323003033130211-3223132211302022-1011103232001133-3101011012212000-2123200123233222", "registry_path": "docs/guides/resources--data_type--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/examples/resource/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Resource for xcsh_data_type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["data_typeCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_data_type/resource.tf`; digest `sha256:292681f3b6c18ec2ffbe2be1a0e17beb9c57b44e1777befb42567a6c3a8f12ee`.

```terraform
# DataType Resource Example
# Manages data_type creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataType configuration
resource "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}
```
