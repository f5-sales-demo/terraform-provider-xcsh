---
page_title: "xcsh_data_group"
subcategory: ""
description: "Manages data group in a given namespace. If one already exists it will give an error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["data group"], "body_bytes": 1514, "body_sha256": "sha256:91c3205d3fb5ab2807030af7790c0fb7bdbb231ee7563ab1b5ce62afe7b71280", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:data_group:reference", "xcsh-docs:resources:data_group:examples", "xcsh-docs:resources:data_group:import", "xcsh-docs:resources:data_group:timeouts"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_group:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/data_group/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2033221010300023-0310201310213312-1111123303023302-1002002103313121-0023323222202020-3323233223120112-1122030002313100-3103222111013031", "registry_path": "docs/resources/data_group.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages data group in a given namespace. If one already exists it will give an error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_data_group

Breadcrumbs:

- xcsh_data_group

Manages data group in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataGroup Resource Example
# Manages data group in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataGroup configuration
resource "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/lifecycle/timeouts/)
