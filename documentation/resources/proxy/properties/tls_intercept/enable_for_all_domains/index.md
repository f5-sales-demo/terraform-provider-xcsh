---
page_title: "tls_intercept.enable_for_all_domains"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls intercept enable for all domains"], "body_bytes": 1017, "body_sha256": "sha256:f4ac362ee3bbf92f7bb969a861e2c653474ea396f68fc481ad266645646af344", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:enable_for_all_domains", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept", "path": "documentation/resources/proxy/properties/tls_intercept/enable_for_all_domains/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0010312210200030-1032102030200221-1013231023111111-0213001030001301-3302211232102022-2010221321321011-1320131130020100-1230311302022102", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_intercept", "enable_for_all_domains"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/enable_for_all_domains/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.enable_for_all_domains

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/)
- tls_intercept.enable_for_all_domains

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable for all domains.

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
enable_for_all_domains = {}
```

This is an empty object or choice marker. It has no direct properties.
