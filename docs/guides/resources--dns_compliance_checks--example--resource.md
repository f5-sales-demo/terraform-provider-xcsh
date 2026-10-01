---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dns_compliance_checks."
xcsh_docs: {"aliases": [], "body_bytes": 1172, "body_sha256": "sha256:c84af71ebcdecc42bd2c082ac41a924e45497919d9ed4f5d8659b28749b98362", "canonical_id": "xcsh-docs:resources:dns_compliance_checks:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_compliance_checks:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4b55554787f8c66271c8f849e8f82effc0413dcd79554cefe77c85a2e413eba1", "source_path": "examples/resources/xcsh_dns_compliance_checks/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_compliance_checks:example:resource", "parent_id": "xcsh-docs:resources:dns_compliance_checks:examples", "path": "docs/guides/resources--dns_compliance_checks--example--resource.md", "provider_name": "dns_compliance_checks", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_compliance_checks/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_dns_compliance_checks.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_compliance_checksCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md)
- [Examples](resources--dns_compliance_checks--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_compliance_checks/resource.tf`; digest `sha256:4b55554787f8c66271c8f849e8f82effc0413dcd79554cefe77c85a2e413eba1`.

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

## Next pages

- [Examples](resources--dns_compliance_checks--examples.md)
- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md)
