---
page_title: "dynamic_proxy.disable_dns_masquerade"
subcategory: ""
description: "dynamic_proxy.disable_dns_masquerade for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 997, "body_sha256": "sha256:8e72b7ea911b103876887b58c42a4785f1409705020a2f53a7f3c18aad57ab54", "canonical_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:disable_dns_masquerade", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:disable_dns_masquerade", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy", "path": "docs/guides/resources--proxy--properties--dynamic_proxy--disable_dns_masquerade.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "disable_dns_masquerade"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/disable_dns_masquerade/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.disable_dns_masquerade for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.disable_dns_masquerade

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- dynamic_proxy.disable_dns_masquerade

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable dns masquerade.

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
disable_dns_masquerade = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- [xcsh_proxy](../resources/proxy.md)
