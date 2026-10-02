---
page_title: "xcsh_virtual_k8s"
subcategory: "Container"
description: "Manages virtual_k8s will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "upstream servers", "virtual k8s"], "body_bytes": 1743, "body_sha256": "sha256:5451b956884054861960561207342769fad553decefe2f4bd64e154e343fcefa", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:virtual_k8s:reference", "xcsh-docs:resources:virtual_k8s:examples", "xcsh-docs:resources:virtual_k8s:import", "xcsh-docs:resources:virtual_k8s:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_k8s:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_k8s:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/virtual_k8s/index.md", "product": "distributed-cloud", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2231213311220130-3130200112303103-1233222013120311-0300110031301022-3200010102220211-1202030313103211-3121033330120121-3020002013232100", "registry_path": "docs/resources/virtual_k8s.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:optional", "target_id": "xcsh-docs:resources:workload:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_k8s/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages virtual_k8s will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_virtual_k8s

Breadcrumbs:

- xcsh_virtual_k8s

Manages virtual\_k8s will create the object in the storage backend for namespace metadata.namespace
in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `workload`.

- workload: Container workloads in this namespace

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualK8S Resource Example
# Manages virtual_k8s will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualK8S configuration
resource "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/lifecycle/timeouts/)
