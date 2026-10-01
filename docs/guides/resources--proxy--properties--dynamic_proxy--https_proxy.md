---
page_title: "dynamic_proxy.https_proxy"
subcategory: ""
description: "dynamic_proxy.https_proxy for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1393, "body_sha256": "sha256:3071255c86d18ca14568bf32d5255cccb1529bdae7995f28603f23f69d1217b7", "canonical_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option", "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy", "path": "docs/guides/resources--proxy--properties--dynamic_proxy--https_proxy.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.https_proxy for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- dynamic_proxy.https_proxy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for https proxy.

Upstream description:

Parameters for dynamic HTTPS proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
https_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [more_option](resources--proxy--properties--dynamic_proxy--https_proxy--more_option.md): complete subsection reference.

- [tls_params](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md): complete subsection reference.

## Next pages

- [dynamic_proxy.https_proxy.more_option](resources--proxy--properties--dynamic_proxy--https_proxy--more_option.md)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- [xcsh_proxy](../resources/proxy.md)
