---
page_title: "irules"
subcategory: ""
description: "irules for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 995, "body_sha256": "sha256:253bbabbf47617cbd45d62e89f976be49711bcd5b6ab05742d87c194517c8e8c", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:irules", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:irules:irules"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:irules", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "docs/guides/resources--bigip_http_proxy--properties--irules.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["irules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/irules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "irules for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# irules

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- irules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IRules Configuration for downstream connections.

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
irules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [irules](resources--bigip_http_proxy--properties--irules--irules.md): complete subsection reference.

## Next pages

- [irules.irules](resources--bigip_http_proxy--properties--irules--irules.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
