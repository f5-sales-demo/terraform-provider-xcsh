---
page_title: "xcsh_site_image"
subcategory: ""
description: "Resolve the current KVM image using exactly one Site owned by the named SMSv2 configuration. No caller UID or static fallback is supported. Verify the artifact MD5 before use; image resolution does not imply successful boot."
xcsh_docs: {"aliases": ["site image", "succeeded", "success", "successful"], "body_bytes": 1411, "body_sha256": "sha256:a967cc064080cfd5b8d7f734a5e5f46ae2ca99899f4c76d38497e6e45531b711", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site_image:reference", "xcsh-docs:data-sources:site_image:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_image:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_image:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/site_image/index.md", "product": "distributed-cloud", "provider_name": "site_image", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3131222302011312-0030101332231322-0021111121002033-3033310200203330-1301310123023111-0121222100300001-3102203221303023-0331320320202102", "registry_path": "docs/data-sources/site_image.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_image/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Resolve the current KVM image using exactly one Site owned by the named SMSv2 configuration. No caller UID or static fallback is supported. Verify the artifact MD5 before use; image resolution does not imply successful boot.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/examples/)
