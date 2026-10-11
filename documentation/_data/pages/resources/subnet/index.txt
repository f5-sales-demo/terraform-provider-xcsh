---
page_title: "xcsh_subnet"
subcategory: ""
description: "Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod. it is created in user or shared namespace. configuration."
xcsh_docs: {"aliases": ["subnet"], "body_bytes": 1638, "body_sha256": "sha256:2ee3612b2c254173b6821229a7d78f128140bab8a6bb8266b10e93a0b8378a32", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:subnet:reference", "xcsh-docs:resources:subnet:examples", "xcsh-docs:resources:subnet:import", "xcsh-docs:resources:subnet:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/subnet/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001", "registry_path": "docs/resources/subnet.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod. it is created in user or shared namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["subnetCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_subnet

Breadcrumbs:

- xcsh_subnet

Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an
interface of a vm/pod. it is created in user or shared namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Subnet Resource Example
# Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Subnet configuration
resource "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/lifecycle/timeouts/)
