---
page_title: "xcsh_gcp_vpc_site"
subcategory: "Infrastructure"
description: "Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud VPC environments."
xcsh_docs: {"aliases": ["gcp vpc site"], "body_bytes": 1496, "body_sha256": "sha256:864b5a78703db0c4b8f35cc6c7ccd6aa49c049c289e09e378d0b2d1aefdb8f6a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:reference", "xcsh-docs:data-sources:gcp_vpc_site:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/gcp_vpc_site/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202", "registry_path": "docs/data-sources/gcp_vpc_site.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:required", "target_id": "xcsh-docs:resources:cloud_credentials:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud VPC environments.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_gcp_vpc_site

Breadcrumbs:

- xcsh_gcp_vpc_site

Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud
VPC environments.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: GCP authentication for deployment

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GCPVPCSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GCPVPCSite by name
data "xcsh_gcp_vpc_site" "example" {
  name      = "example-gcp-vpc-site"
  namespace = "staging"
}

output "gcp_vpc_site_id" {
  value = data.xcsh_gcp_vpc_site.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/examples/)
