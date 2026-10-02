---
page_title: "xcsh_gcp_vpc_site"
subcategory: "Infrastructure"
description: "Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud VPC environments."
xcsh_docs: {"aliases": ["gcp vpc site"], "body_bytes": 1895, "body_sha256": "sha256:1d52c3ac0734c63331b9532008d77f21bdff1cf92cedf019e5279ee6a0dd5947", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:reference", "xcsh-docs:resources:gcp_vpc_site:examples", "xcsh-docs:resources:gcp_vpc_site:import", "xcsh-docs:resources:gcp_vpc_site:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/gcp_vpc_site/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122", "registry_path": "docs/resources/gcp_vpc_site.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:required", "target_id": "xcsh-docs:resources:cloud_credentials:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud VPC environments.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
# GCPVPCSite Resource Example
# Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud VPC environments.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GCPVPCSite configuration
resource "xcsh_gcp_vpc_site" "example" {
  name      = "example-gcp-vpc-site"
  namespace = "staging"

  gcp_region    = "example-value"
  instance_type = "example-value"
  ssh_key       = "example-value"
}
```

## Root configuration

Required root properties: `gcp_region`, `instance_type`, `name`, `namespace`, `ssh_key`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/lifecycle/timeouts/)
