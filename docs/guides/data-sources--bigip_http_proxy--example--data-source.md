---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1123, "body_sha256": "sha256:1d4ad7245a6bad07061552831e4253e5ef8d05ac75cab9605eee49bfc8a7c273", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:436b4e693b05709f4b55969967c5b9c433a210d85502a116c164d87cdf218449", "source_path": "examples/data-sources/xcsh_bigip_http_proxy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bigip_http_proxy:example:data-source", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:examples", "path": "docs/guides/data-sources--bigip_http_proxy--example--data-source.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Examples](data-sources--bigip_http_proxy--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bigip_http_proxy/data-source.tf`; digest `sha256:436b4e693b05709f4b55969967c5b9c433a210d85502a116c164d87cdf218449`.

```terraform
# BigIPHTTPProxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPHTTPProxy by name
data "xcsh_bigip_http_proxy" "example" {
  name      = "example-bigip-http-proxy"
  namespace = "staging"
}

output "bigip_http_proxy_id" {
  value = data.xcsh_bigip_http_proxy.example.id
}
```

## Next pages

- [Examples](data-sources--bigip_http_proxy--examples.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
