---
page_title: "proxy_config.https_auto_cert.use_mtls.xfcc_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["proxy config https auto cert use mtls xfcc disabled"], "body_bytes": 1381, "body_sha256": "sha256:ed53f05be88f0de1bd9372d1a2b4fcdb87ab95feebfe74b57755e842257b68b5", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls:xfcc_disabled", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:use_mtls", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/xfcc_disabled/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3211201331120032-3121223100220002-2322132131000133-0022022033220231-1113210301102222-3110303311313013-2321030022222123-2221030200322002", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https_auto_cert", "use_mtls", "xfcc_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/xfcc_disabled/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https_auto_cert.use_mtls.xfcc_disabled

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/)
- [proxy_config.https_auto_cert.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/)
- proxy_config.https_auto_cert.use_mtls.xfcc_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.
