---
page_title: "Data source"
subcategory: "Container"
description: "Data source for xcsh_virtual_k8s."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1266, "body_sha256": "sha256:f236149f09d90c3f00b7de744d37091ff6da4711109d284dfb9970cc41da970a", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_k8s:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:85c7894f9985a87a12ec00003bcdb34e2727cf22575e06597a586ec7fcc0019d", "source_path": "examples/data-sources/xcsh_virtual_k8s/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:virtual_k8s:example:data-source", "parent_id": "xcsh-docs:data-sources:virtual_k8s:examples", "path": "documentation/data-sources/virtual_k8s/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1220312122123020-3330103200030130-2312131021103011-3010313313132313-1102122321123330-0230231102220120-2220001102200031-1013300102120123", "registry_path": "docs/guides/data-sources--virtual_k8s--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_k8s/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_virtual_k8s.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
