---
page_title: "xcsh_dns_compliance_checks"
subcategory: ""
description: "Manages DNS Compliance Checks Specification in a given namespace. If one already exists it will give an error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["dns compliance checks"], "body_bytes": 1730, "body_sha256": "sha256:a857e1c89f1d0aa0cb96071a397a1bc33c2ba2068d8c65a70e0055f83fe64017", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_compliance_checks:reference", "xcsh-docs:resources:dns_compliance_checks:examples", "xcsh-docs:resources:dns_compliance_checks:import", "xcsh-docs:resources:dns_compliance_checks:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_compliance_checks:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_compliance_checks:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/dns_compliance_checks/index.md", "product": "distributed-cloud", "provider_name": "dns_compliance_checks", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2320330031033230-1032132011323312-3322232320012321-0210130122222101-3131030333303122-1013130113312022-2213230000130100-1010120232210210", "registry_path": "docs/resources/dns_compliance_checks.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_compliance_checks/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Manages DNS Compliance Checks Specification in a given namespace. If one already exists it will give an error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["dns_compliance_checksCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_compliance_checks

Breadcrumbs:

- xcsh_dns_compliance_checks

Manages DNS Compliance Checks Specification in a given namespace. If one already exists it will give
an error in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSComplianceChecks Resource Example
# Manages DNS Compliance Checks Specification in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSComplianceChecks configuration
resource "xcsh_dns_compliance_checks" "example" {
  name      = "example-dns-compliance-checks"
  namespace = "staging"

  domain_denylist = ["example-value"]
}
```

## Root configuration

Required root properties: `domain_denylist`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/lifecycle/timeouts/)
