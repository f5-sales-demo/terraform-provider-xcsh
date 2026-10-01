---
page_title: "xcsh_cloud_elastic_ip"
subcategory: ""
description: "xcsh_cloud_elastic_ip for xcsh_cloud_elastic_ip."
xcsh_docs: {"aliases": [], "body_bytes": 1310, "body_sha256": "sha256:1dcaf5b0a7a87d9ff7d7a5ccc52f8e41eafbbc46615b561b89956aff154c9ef2", "canonical_id": "xcsh-docs:data-sources:cloud_elastic_ip:fundamentals", "child_ids": ["xcsh-docs:data-sources:cloud_elastic_ip:reference", "xcsh-docs:data-sources:cloud_elastic_ip:examples"], "collection_id": "xcsh-docs:data-sources:cloud_elastic_ip:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_elastic_ip:fundamentals", "parent_id": null, "path": "docs/data-sources/cloud_elastic_ip.md", "provider_name": "cloud_elastic_ip", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_elastic_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cloud_elastic_ip for xcsh_cloud_elastic_ip.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_elastic_ipCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Property reference](../guides/data-sources--cloud_elastic_ip--reference.md)
- [Examples](../guides/data-sources--cloud_elastic_ip--examples.md)
