---
page_title: "Resource"
subcategory: "Container"
description: "Resource for xcsh_workload."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1200, "body_sha256": "sha256:490127396ec861ccdae425a09d4f8dd1e7ad67eb50fc3f65fc4b4b1540f13684", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:95bc586cf100a6ef22147aba592992dc22645f4b290f9226bf45165f3414e58c", "source_path": "examples/resources/xcsh_workload/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:workload:example:resource", "parent_id": "xcsh-docs:resources:workload:examples", "path": "documentation/resources/workload/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3202301020033123-1000311020333023-2200330011313010-3211330331312122-3231102022332233-0023313320023323-2321332001221132-0001132011323320", "registry_path": "docs/guides/resources--workload--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_workload/resource.tf`; digest `sha256:95bc586cf100a6ef22147aba592992dc22645f4b290f9226bf45165f3414e58c`.

```terraform
# Workload Resource Example
# Manages a Workload resource in F5 Distributed Cloud for workload.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Workload configuration
resource "xcsh_workload" "example" {
  name      = "example-workload"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/examples/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
