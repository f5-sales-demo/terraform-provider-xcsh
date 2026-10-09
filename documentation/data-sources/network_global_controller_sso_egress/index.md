---
page_title: "xcsh_network_global_controller_sso_egress"
subcategory: ""
description: "Global Controller SSO egress IPv4 addresses. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest."
xcsh_docs: {"aliases": ["network global controller sso egress"], "body_bytes": 1572, "body_sha256": "sha256:b6b59d70758113a1f0c729322521bc85fca591ccd1c104f245fe62f54ed1060e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_global_controller_sso_egress:reference", "xcsh-docs:data-sources:network_global_controller_sso_egress:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_global_controller_sso_egress:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_global_controller_sso_egress:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_global_controller_sso_egress/index.md", "product": "distributed-cloud", "provider_name": "network_global_controller_sso_egress", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0312323020200211-1203202133322020-0021231103030131-3121030331030302-1102330101113213-3300113123100001-1202022320202320-2331132003123002", "registry_path": "docs/data-sources/network_global_controller_sso_egress.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_global_controller_sso_egress/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Global Controller SSO egress IPv4 addresses. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_global_controller_sso_egress

Breadcrumbs:

- xcsh_network_global_controller_sso_egress

Global Controller SSO egress IPv4 addresses. Values are bundled from the pinned OpenAPI release;
this data source performs no network request. Ports and traffic direction are not encoded in the
manifest.

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

data "xcsh_network_global_controller_sso_egress" "https" {}

output "global_controller_sso_https" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_global_controller_sso_egress.https.cidr_blocks
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_controller_sso_egress/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_global_controller_sso_egress/examples/)
