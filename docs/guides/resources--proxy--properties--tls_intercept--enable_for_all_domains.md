---
page_title: "tls_intercept.enable_for_all_domains"
subcategory: ""
description: "tls_intercept.enable_for_all_domains for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 898, "body_sha256": "sha256:0fe556847d558224437e68e59cbf262441fbc364b57602e96fc6b3300eb4c6ed", "canonical_id": "xcsh-docs:resources:proxy:properties:tls_intercept:enable_for_all_domains", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:enable_for_all_domains", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept", "path": "docs/guides/resources--proxy--properties--tls_intercept--enable_for_all_domains.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_intercept", "enable_for_all_domains"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/enable_for_all_domains/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_intercept.enable_for_all_domains for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_intercept.enable_for_all_domains

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [tls_intercept](resources--proxy--properties--tls_intercept.md)
- tls_intercept.enable_for_all_domains

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable for all domains.

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
enable_for_all_domains = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_intercept](resources--proxy--properties--tls_intercept.md)
- [xcsh_proxy](../resources/proxy.md)
