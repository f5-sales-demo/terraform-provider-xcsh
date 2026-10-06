---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_usb_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 993, "body_sha256": "sha256:50700df1288f8c2dbdbda63bc89c41f828733d7bf17481fca5f834b3f03d796b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:usb_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d81199b0fb106ffd6cc8d506207e603f1169d2cbf87d6b31f0ac720b8b44794e", "source_path": "examples/resources/xcsh_usb_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:usb_policy:example:resource", "parent_id": "xcsh-docs:resources:usb_policy:examples", "path": "documentation/resources/usb_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "usb_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1021311230322121-1223221313333000-3302100222013302-0322101321332312-2111102333322331-3020210003100310-1200012222211011-1102301103010312", "registry_path": "docs/guides/resources--usb_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/usb_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_usb_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["usb_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_usb_policy/resource.tf`; digest `sha256:d81199b0fb106ffd6cc8d506207e603f1169d2cbf87d6b31f0ac720b8b44794e`.

```terraform
# UsbPolicy Resource Example
# Manages new USB policy object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UsbPolicy configuration
resource "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}
```
