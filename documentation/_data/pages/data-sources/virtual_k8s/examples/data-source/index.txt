---
page_title: "Data source"
subcategory: "Container"
description: "Data source for xcsh_virtual_k8s."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1266, "body_sha256": "sha256:f236149f09d90c3f00b7de744d37091ff6da4711109d284dfb9970cc41da970a", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_k8s:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:85c7894f9985a87a12ec00003bcdb34e2727cf22575e06597a586ec7fcc0019d", "source_path": "examples/data-sources/xcsh_virtual_k8s/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:virtual_k8s:example:data-source", "parent_id": "xcsh-docs:data-sources:virtual_k8s:examples", "path": "documentation/data-sources/virtual_k8s/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1220312122123020-3330103200030130-2312131021103011-3010313313132313-1102122321123330-0230231102220120-2220001102200031-1013300102120123", "registry_path": "docs/guides/data-sources--virtual_k8s--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_k8s/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_virtual_k8s.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_virtual_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_k8s/data-source.tf`; digest `sha256:85c7894f9985a87a12ec00003bcdb34e2727cf22575e06597a586ec7fcc0019d`.

```terraform
# VirtualK8S Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualK8S by name
data "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}

output "virtual_k8s_id" {
  value = data.xcsh_virtual_k8s.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/examples/)
- [xcsh_virtual_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/)
