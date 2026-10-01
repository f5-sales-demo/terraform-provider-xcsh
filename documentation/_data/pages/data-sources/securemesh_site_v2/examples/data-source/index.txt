---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1354, "body_sha256": "sha256:8fa8b17f9ac1b642b30941c85404c832fd2866b51f2c330265fe772c04728d30", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2abaac1742f086ab988ef94e8d836fe6f538c549072b6b6dd846213e69e59239", "source_path": "examples/data-sources/xcsh_securemesh_site_v2/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:securemesh_site_v2:example:data-source", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:examples", "path": "documentation/data-sources/securemesh_site_v2/examples/data-source/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_securemesh_site_v2/data-source.tf`; digest `sha256:2abaac1742f086ab988ef94e8d836fe6f538c549072b6b6dd846213e69e59239`.

```terraform
# SecuremeshSiteV2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSiteV2 by name
data "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}

output "securemesh_site_v2_id" {
  value = data.xcsh_securemesh_site_v2.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/examples/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
