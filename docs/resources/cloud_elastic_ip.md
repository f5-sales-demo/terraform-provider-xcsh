---
page_title: "xcsh_cloud_elastic_ip"
subcategory: ""
description: "xcsh_cloud_elastic_ip for xcsh_cloud_elastic_ip."
xcsh_docs: {"aliases": [], "body_bytes": 1486, "body_sha256": "sha256:e90ac0c473176580f991565f92dde549194b9e4bb325024ef764a81367d846aa", "canonical_id": "xcsh-docs:resources:cloud_elastic_ip:fundamentals", "child_ids": ["xcsh-docs:resources:cloud_elastic_ip:reference", "xcsh-docs:resources:cloud_elastic_ip:examples", "xcsh-docs:resources:cloud_elastic_ip:import", "xcsh-docs:resources:cloud_elastic_ip:timeouts"], "collection_id": "xcsh-docs:resources:cloud_elastic_ip:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_elastic_ip:fundamentals", "parent_id": null, "path": "docs/resources/cloud_elastic_ip.md", "provider_name": "cloud_elastic_ip", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_elastic_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cloud_elastic_ip for xcsh_cloud_elastic_ip.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_elastic_ipCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/resources--cloud_elastic_ip--reference.md)
- [Examples](../guides/resources--cloud_elastic_ip--examples.md)
- [Import](../guides/resources--cloud_elastic_ip--import.md)
- [Timeouts](../guides/resources--cloud_elastic_ip--timeouts.md)
