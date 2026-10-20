---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_app_firewall."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1008, "body_sha256": "sha256:7e6d87060a281214432658b4052978881caac881fafedcf1da6c3507bb67a551", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:03e5e1d1ea7b4ce0a0005da7cf6bf27eeca5913c46ae5552f44fffd430a9cb75", "source_path": "examples/resources/xcsh_app_firewall/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:resource", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2212130033200011-3131123033222113-1312132301000221-3321311011213200-3103222201202231-3000320132003002-2033331001201113-3331122230233023", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/resource/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Resource for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["app_firewallCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/resource.tf`; digest `sha256:03e5e1d1ea7b4ce0a0005da7cf6bf27eeca5913c46ae5552f44fffd430a9cb75`.

```terraform
# AppFirewall Resource Example
# Manages Application Firewall in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppFirewall configuration
resource "xcsh_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "staging"
}
```
