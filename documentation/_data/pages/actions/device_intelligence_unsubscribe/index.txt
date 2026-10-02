---
page_title: "xcsh_device_intelligence_unsubscribe"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["device intelligence unsubscribe"], "body_bytes": 1321, "body_sha256": "sha256:722af403224e5bf6386ea46e3b02b9f2cc198a38b844227267c4fac6ee6fc4bf", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:device_intelligence_unsubscribe:reference", "xcsh-docs:actions:device_intelligence_unsubscribe:examples", "xcsh-docs:actions:device_intelligence_unsubscribe:lifecycle"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:device_intelligence_unsubscribe:collection", "completeness": "complete", "id": "xcsh-docs:actions:device_intelligence_unsubscribe:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/device_intelligence_unsubscribe/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_unsubscribe", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "actions", "registry_anchor": "canonical-3110012200313320-0120221200221312-0230333311202000-0331023323320332-1101121132213210-3031113003001131-3201203222213110-2210102303020312", "registry_path": "docs/actions/device_intelligence_unsubscribe.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_unsubscribe/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_unsubscribe

Breadcrumbs:

- xcsh_device_intelligence_unsubscribe

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceUnsubscribe Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_device_intelligence_unsubscribe" "example" {
  config {
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/lifecycle/)
