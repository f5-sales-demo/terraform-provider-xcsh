---
page_title: "proxy_config.https.pass_through"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["proxy config https pass through"], "body_bytes": 1447, "body_sha256": "sha256:3cc1827b3077082f8ca00bf30f9ac58df54f8339fd78dc79c6fb20498f03f717", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:pass_through", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/https/pass_through/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2310222311302101-2121110111210313-3033122233131333-3320311202312031-1102212112100311-2120323112302121-3313233100202102-3132232232313210", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https", "pass_through"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/pass_through/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https.pass_through

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/)
- proxy_config.https.pass_through

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
