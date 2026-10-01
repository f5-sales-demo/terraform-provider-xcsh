---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cloud_elastic_ip."
xcsh_docs: {"aliases": [], "body_bytes": 1123, "body_sha256": "sha256:ca8499ce53e30570aff47352a06747c8fc59bc6d2a087384527e9d36e6aba416", "canonical_id": "xcsh-docs:data-sources:cloud_elastic_ip:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cloud_elastic_ip:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a222dbb0119e3d807c71bf21695db7199a7a0767160af1c0ed21e3f33eb7b545", "source_path": "examples/data-sources/xcsh_cloud_elastic_ip/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cloud_elastic_ip:example:data-source", "parent_id": "xcsh-docs:data-sources:cloud_elastic_ip:examples", "path": "docs/guides/data-sources--cloud_elastic_ip--example--data-source.md", "provider_name": "cloud_elastic_ip", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_elastic_ip/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_cloud_elastic_ip.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_elastic_ipCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md)
- [Examples](data-sources--cloud_elastic_ip--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_elastic_ip/data-source.tf`; digest `sha256:a222dbb0119e3d807c71bf21695db7199a7a0767160af1c0ed21e3f33eb7b545`.

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

## Next pages

- [Examples](data-sources--cloud_elastic_ip--examples.md)
- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md)
