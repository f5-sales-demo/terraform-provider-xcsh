---
page_title: "xcsh_dns_compliance_checks"
subcategory: ""
description: "xcsh_dns_compliance_checks for xcsh_dns_compliance_checks."
xcsh_docs: {"aliases": [], "body_bytes": 1631, "body_sha256": "sha256:9c61e04b5b498fb1c34d8dd5c8010db652d4fd58d3202faf59f29181833025c2", "child_ids": ["xcsh-docs:resources:dns_compliance_checks:reference", "xcsh-docs:resources:dns_compliance_checks:examples", "xcsh-docs:resources:dns_compliance_checks:import", "xcsh-docs:resources:dns_compliance_checks:timeouts"], "collection_id": "xcsh-docs:resources:dns_compliance_checks:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_compliance_checks:fundamentals", "parent_id": null, "path": "documentation/resources/dns_compliance_checks/index.md", "provider_name": "dns_compliance_checks", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_compliance_checks/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_dns_compliance_checks for xcsh_dns_compliance_checks.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_compliance_checksCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
