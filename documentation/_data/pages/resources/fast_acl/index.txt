---
page_title: "xcsh_fast_acl"
subcategory: ""
description: "Manages object, object contains rules to protect site from denial of service It has destination{destination IP, destination port) and references to in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["fast acl"], "body_bytes": 1676, "body_sha256": "sha256:9a2978621153fae0fa3c87abaa4a009ef657bb60975af7f9582010f015777bc9", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:reference", "xcsh-docs:resources:fast_acl:examples", "xcsh-docs:resources:fast_acl:import", "xcsh-docs:resources:fast_acl:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/fast_acl/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102", "registry_path": "docs/resources/fast_acl.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages object, object contains rules to protect site from denial of service It has destination{destination IP, destination port) and references to in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_fast_acl

Breadcrumbs:

- xcsh_fast_acl

Manages object, object contains rules to protect site from denial of service It has
destination\{destination IP, destination port) and references to in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACL Resource Example
# Manages object, object contains rules to protect site from denial of service It has destination{destination IP, destination port) and references to in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FastACL configuration
resource "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/lifecycle/timeouts/)
