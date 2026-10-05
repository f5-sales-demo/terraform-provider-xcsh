---
page_title: "xcsh_cloud_elastic_ip"
subcategory: ""
description: "Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["cloud elastic ip"], "body_bytes": 1675, "body_sha256": "sha256:f7d5db4a7ad77a0776cb6474a65ff5feeccb5c6c447a2bbccbc1305fc7159b23", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_elastic_ip:reference", "xcsh-docs:resources:cloud_elastic_ip:examples", "xcsh-docs:resources:cloud_elastic_ip:import", "xcsh-docs:resources:cloud_elastic_ip:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_elastic_ip:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_elastic_ip:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/cloud_elastic_ip/index.md", "product": "distributed-cloud", "provider_name": "cloud_elastic_ip", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1012233313021212-0003222033112111-1303213323303313-1020020302213310-1211110033020313-2122312122003031-2320202301023301-3102202323122130", "registry_path": "docs/resources/cloud_elastic_ip.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_elastic_ip/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cloud_elastic_ipCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cloud_elastic_ip

Breadcrumbs:

- xcsh_cloud_elastic_ip

Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudElasticIP Resource Example
# Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudElasticIP configuration
resource "xcsh_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "staging"

  item_count = 1
}
```

## Root configuration

Required root properties: `item_count`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_elastic_ip/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_elastic_ip/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_elastic_ip/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_elastic_ip/lifecycle/timeouts/)
