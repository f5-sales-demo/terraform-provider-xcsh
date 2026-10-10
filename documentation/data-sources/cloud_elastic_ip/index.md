---
page_title: "xcsh_cloud_elastic_ip"
subcategory: ""
description: "Reads Cloud Elastic IP information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["cloud elastic ip"], "body_bytes": 1359, "body_sha256": "sha256:57f3134e4a320b6d100d3963b3b4f96de0b3c661da1f684a43800f3c779dac7a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_elastic_ip:reference", "xcsh-docs:data-sources:cloud_elastic_ip:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_elastic_ip:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_elastic_ip:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/cloud_elastic_ip/index.md", "product": "distributed-cloud", "provider_name": "cloud_elastic_ip", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0223023303003220-2123332302131330-3202021121113030-2021223333123231-2302313213301131-1233113113023011-3112030123323311-3322022223022102", "registry_path": "docs/data-sources/cloud_elastic_ip.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_elastic_ip/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reads Cloud Elastic IP information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cloud_elastic_ipCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cloud_elastic_ip

Breadcrumbs:

- xcsh_cloud_elastic_ip

Reads Cloud Elastic IP information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudElasticIP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudElasticIP by name
data "xcsh_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "staging"
}

output "cloud_elastic_ip_id" {
  value = data.xcsh_cloud_elastic_ip.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_elastic_ip/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_elastic_ip/examples/)
