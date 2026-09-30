---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_certificate."
xcsh_docs: {"aliases": [], "body_bytes": 1177, "body_sha256": "sha256:f5fb128720a02a2f9cb8a9f8578ad9639bb1a743f07e50c7ce0bb1dfca6145b0", "child_ids": [], "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a89bc2bb9f8d9310b1230ada6b9c044b0eeaf77e938fa8139bb314ba9412a092", "source_path": "examples/resources/xcsh_certificate/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:certificate:example:resource", "parent_id": "xcsh-docs:resources:certificate:examples", "path": "documentation/resources/certificate/examples/resource/index.md", "provider_name": "certificate", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_certificate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_certificate/resource.tf`; digest `sha256:a89bc2bb9f8d9310b1230ada6b9c044b0eeaf77e938fa8139bb314ba9412a092`.

```terraform
# Certificate Resource Example
# Manages a Certificate resource in F5 Distributed Cloud for certificate.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Certificate configuration
resource "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"

  certificate_url = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/examples/)
- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
