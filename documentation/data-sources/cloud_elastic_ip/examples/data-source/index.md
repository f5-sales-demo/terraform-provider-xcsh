---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cloud_elastic_ip."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1089, "body_sha256": "sha256:2bf479eb9a22ab44cb17610022317a999c14a5d0261a79fe802c96fa9419f361", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_elastic_ip:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a222dbb0119e3d807c71bf21695db7199a7a0767160af1c0ed21e3f33eb7b545", "source_path": "examples/data-sources/xcsh_cloud_elastic_ip/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cloud_elastic_ip:example:data-source", "parent_id": "xcsh-docs:data-sources:cloud_elastic_ip:examples", "path": "documentation/data-sources/cloud_elastic_ip/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cloud_elastic_ip", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3203123010020303-2030133310010000-0012121012233011-0331031301032120-0222122233201013-1132201131121321-3123323320313312-1332210203331331", "registry_path": "docs/guides/data-sources--cloud_elastic_ip--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_elastic_ip/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_cloud_elastic_ip.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cloud_elastic_ipCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cloud_elastic_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_elastic_ip/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_elastic_ip/examples/)
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
