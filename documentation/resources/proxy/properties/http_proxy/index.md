---
page_title: "http_proxy"
subcategory: ""
description: "http_proxy for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1678, "body_sha256": "sha256:0ec4b1ea8b5230d0bbeefae698a2d326d3fa0af4bac0bdf727549ef8e00d968e", "child_ids": ["xcsh-docs:resources:proxy:properties:http_proxy:enable_http", "xcsh-docs:resources:proxy:properties:http_proxy:more_option"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:http_proxy", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "documentation/resources/proxy/properties/http_proxy/index.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["http_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/http_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_proxy for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- http_proxy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP Connect Proxy. Parameters for HTTP Connect Proxy.

Upstream description:

Parameters for HTTP Connect Proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_https_choice": "[\"enable_http\"]"
}
```

Terraform syntax:

```terraform
http_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [enable_http](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/enable_http/): complete subsection reference.

- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/): complete subsection reference.

## Next pages

- [http_proxy.enable_http](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/enable_http/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
