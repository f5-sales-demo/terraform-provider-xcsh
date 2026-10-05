---
page_title: "xcsh_site_image"
subcategory: ""
description: "Resolve the current KVM image using exactly one Site owned by the named SMSv2 configuration. No caller UID or static fallback is supported. Verify the artifact MD5 before use; image resolution does not imply successful boot."
xcsh_docs: {"aliases": ["site image", "succeeded", "success", "successful"], "body_bytes": 1398, "body_sha256": "sha256:1b207e85df2e6c0df3f366a20436a0fff462b3fd447330fc9712bb1aa4ba22a8", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site_image:reference", "xcsh-docs:data-sources:site_image:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_image:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_image:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/site_image/index.md", "product": "distributed-cloud", "provider_name": "site_image", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3131222302011312-0030101332231322-0021111121002033-3033310200203330-1301310123023111-0121222100300001-3102203221303023-0331320320202102", "registry_path": "docs/data-sources/site_image.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_image/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resolve the current KVM image using exactly one Site owned by the named SMSv2 configuration. No caller UID or static fallback is supported. Verify the artifact MD5 before use; image resolution does not imply successful boot.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_image

Breadcrumbs:

- xcsh_site_image

Resolve the current KVM image using exactly one Site owned by the named SMSv2 configuration. No
caller UID or static fallback is supported. Verify the artifact MD5 before use; image resolution
does not imply successful boot.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteImage DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_image" "example" {
  site_name = "example-value"
}

output "site_image_result" {
  value     = data.xcsh_site_image.example
  sensitive = true
}
```

## Root configuration

Required root properties: `site_name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/examples/)
