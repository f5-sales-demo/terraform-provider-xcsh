---
page_title: "xcsh_network_dnslb_health_checks"
subcategory: ""
description: "DNS Load Balancer health-check probe IPv4 addresses. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest."
xcsh_docs: {"aliases": ["network dnslb health checks"], "body_bytes": 1596, "body_sha256": "sha256:1128ff45ea6f1724bc80ba3867511c432207532d1ca4eaf7fa6452c9bc62e352", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_dnslb_health_checks:reference", "xcsh-docs:data-sources:network_dnslb_health_checks:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_dnslb_health_checks:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_dnslb_health_checks:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_dnslb_health_checks/index.md", "product": "distributed-cloud", "provider_name": "network_dnslb_health_checks", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1333320111300032-2212003230021231-1213001111023002-3323113032220220-3102221212223011-3012000302133013-0331320230013312-1210133113122312", "registry_path": "docs/data-sources/network_dnslb_health_checks.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_dnslb_health_checks/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "DNS Load Balancer health-check probe IPv4 addresses. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_dnslb_health_checks

Breadcrumbs:

- xcsh_network_dnslb_health_checks

DNS Load Balancer health-check probe IPv4 addresses. Values are bundled from the pinned OpenAPI
release; this data source performs no network request. Ports and traffic direction are not encoded
in the manifest.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_dnslb_health_checks" "https_probe" {}

# Match this explicit ingress port to the monitored endpoint.
output "https_health_check_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_dnslb_health_checks.https_probe.cidr_blocks
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_dnslb_health_checks/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_dnslb_health_checks/examples/)
