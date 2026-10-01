---
page_title: "tls_intercept.volterra_trusted_ca"
subcategory: ""
description: "tls_intercept.volterra_trusted_ca for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 985, "body_sha256": "sha256:d46f5c8d316871fa8ea79fbb10230bc93e0ab3c11b470bb7d5f685f6751ab4f1", "canonical_id": "xcsh-docs:resources:proxy:properties:tls_intercept:volterra_trusted_ca", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:volterra_trusted_ca", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept", "path": "docs/guides/resources--proxy--properties--tls_intercept--volterra_trusted_ca.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_intercept", "volterra_trusted_ca"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/volterra_trusted_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_intercept.volterra_trusted_ca for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.volterra_trusted_ca

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [tls_intercept](resources--proxy--properties--tls_intercept.md)
- tls_intercept.volterra_trusted_ca

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

Upstream description:

This can be used for messages where no values are needed.

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
volterra_trusted_ca = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_intercept](resources--proxy--properties--tls_intercept.md)
- [xcsh_proxy](../resources/proxy.md)
