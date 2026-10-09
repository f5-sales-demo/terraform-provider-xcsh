---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_security_evidence."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1050, "body_sha256": "sha256:63d4be4f1c90c9015c6fd6d49f3ffb306e26e8a6239c389cc0e21159e1ec7ce6", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_security_evidence:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a712a6cc48d2b8d3894590d0570f0c9aa69c910e418b124dca6ff57b24d59d2a", "source_path": "examples/data-sources/xcsh_app_security_evidence/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_security_evidence:example:data-source", "parent_id": "xcsh-docs:data-sources:app_security_evidence:examples", "path": "documentation/data-sources/app_security_evidence/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "app_security_evidence", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1021122003121201-1300022010132211-1212102231110310-2122003313212300-1012310133112113-1113332321002312-2112112331331110-2221133121001122", "registry_path": "docs/guides/data-sources--app_security_evidence--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_security_evidence/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_app_security_evidence.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_security_evidence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_security_evidence/data-source.tf`; digest `sha256:a712a6cc48d2b8d3894590d0570f0c9aa69c910e418b124dca6ff57b24d59d2a`.

```terraform
# AppSecurityEvidence DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_app_security_evidence" "example" {
  namespace = "example-value"
}

output "app_security_evidence_result" {
  value = data.xcsh_app_security_evidence.example
}
```
