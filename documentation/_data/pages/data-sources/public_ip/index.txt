---
page_title: "xcsh_public_ip"
subcategory: ""
description: "xcsh_public_ip for xcsh_public_ip."
xcsh_docs: {"aliases": [], "body_bytes": 1406, "body_sha256": "sha256:79211e99c1f1aa0867a406e7209e9095a5da36263e4f3cd3516252f0a096f015", "child_ids": ["xcsh-docs:data-sources:public_ip:reference", "xcsh-docs:data-sources:public_ip:examples"], "collection_id": "xcsh-docs:data-sources:public_ip:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:public_ip:fundamentals", "parent_id": null, "path": "documentation/data-sources/public_ip/index.md", "provider_name": "public_ip", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/public_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_public_ip for xcsh_public_ip.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_public_ip

Breadcrumbs:

- xcsh_public_ip

Manages a Public IP resource in F5 Distributed Cloud for get public\_ip will get the object from the
storage backend for namespace metadata.namespace. configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# PublicIP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PublicIP by name
data "xcsh_public_ip" "example" {
  name      = "example-public-ip"
  namespace = "staging"
}

output "public_ip_id" {
  value = data.xcsh_public_ip.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/examples/)
