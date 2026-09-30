---
page_title: "http_proxy"
subcategory: ""
description: "http_proxy for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1171, "body_sha256": "sha256:82f17807f51c4a23eedbde6305bbc25996db8675e9cc58a51d538ec0fe9e3723", "canonical_id": "xcsh-docs:resources:proxy:properties:http_proxy", "child_ids": ["xcsh-docs:resources:proxy:properties:http_proxy:enable_http", "xcsh-docs:resources:proxy:properties:http_proxy:more_option"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:http_proxy", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "docs/guides/resources--proxy--properties--http_proxy.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/http_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_proxy for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# http_proxy

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
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

- [enable_http](resources--proxy--properties--http_proxy--enable_http.md): complete subsection reference.

- [more_option](resources--proxy--properties--http_proxy--more_option.md): complete subsection reference.

## Next pages

- [http_proxy.enable_http](resources--proxy--properties--http_proxy--enable_http.md)
- [http_proxy.more_option](resources--proxy--properties--http_proxy--more_option.md)
- [Property reference](resources--proxy--reference.md)
- [xcsh_proxy](../resources/proxy.md)
