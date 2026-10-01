---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_dns_compliance_checks."
xcsh_docs: {"aliases": [], "body_bytes": 1394, "body_sha256": "sha256:7f4f832cd5f9b1bcaa111c330d093434a108d2f9e67018dd6de65b56cef2e0b2", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_compliance_checks:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:494f38d0fa830fb2beff3f30865ce3a3a079a872bc62acf125c470d1e3df90d4", "source_path": "examples/data-sources/xcsh_dns_compliance_checks/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_compliance_checks:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_compliance_checks:examples", "path": "documentation/data-sources/dns_compliance_checks/examples/data-source/index.md", "provider_name": "dns_compliance_checks", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_compliance_checks/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_dns_compliance_checks.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_compliance_checksCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dns_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_compliance_checks/data-source.tf`; digest `sha256:494f38d0fa830fb2beff3f30865ce3a3a079a872bc62acf125c470d1e3df90d4`.

```terraform
# DNSComplianceChecks Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSComplianceChecks by name
data "xcsh_dns_compliance_checks" "example" {
  name      = "example-dns-compliance-checks"
  namespace = "staging"
}

output "dns_compliance_checks_id" {
  value = data.xcsh_dns_compliance_checks.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/examples/)
- [xcsh_dns_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/)
