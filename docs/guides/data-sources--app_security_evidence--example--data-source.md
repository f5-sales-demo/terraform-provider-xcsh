---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_security_evidence."
xcsh_docs: {"aliases": [], "body_bytes": 1099, "body_sha256": "sha256:3b0cfae088f7b94c4b4227f656575deb883caee9d6ddc760e4617c1af0055272", "canonical_id": "xcsh-docs:data-sources:app_security_evidence:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:app_security_evidence:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a712a6cc48d2b8d3894590d0570f0c9aa69c910e418b124dca6ff57b24d59d2a", "source_path": "examples/data-sources/xcsh_app_security_evidence/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_security_evidence:example:data-source", "parent_id": "xcsh-docs:data-sources:app_security_evidence:examples", "path": "docs/guides/data-sources--app_security_evidence--example--data-source.md", "provider_name": "app_security_evidence", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_security_evidence/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_app_security_evidence.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md)
- [Examples](data-sources--app_security_evidence--examples.md)
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

## Next pages

- [Examples](data-sources--app_security_evidence--examples.md)
- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md)
