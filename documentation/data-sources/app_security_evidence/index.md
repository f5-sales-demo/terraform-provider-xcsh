---
page_title: "xcsh_app_security_evidence"
subcategory: ""
description: "xcsh_app_security_evidence for xcsh_app_security_evidence."
xcsh_docs: {"aliases": [], "body_bytes": 1266, "body_sha256": "sha256:f84e6a5f4cbb2673f762e57ec41e7698398e8efc07def7baa17c4ff06b64eb0a", "child_ids": ["xcsh-docs:data-sources:app_security_evidence:reference", "xcsh-docs:data-sources:app_security_evidence:examples"], "collection_id": "xcsh-docs:data-sources:app_security_evidence:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_security_evidence:fundamentals", "parent_id": null, "path": "documentation/data-sources/app_security_evidence/index.md", "provider_name": "app_security_evidence", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_security_evidence/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_app_security_evidence for xcsh_app_security_evidence.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_app_security_evidence

Breadcrumbs:

- xcsh_app_security_evidence

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/examples/)
