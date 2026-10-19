---
page_title: "origin_servers.origin_servers.no_preference"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["origin servers origin servers no preference"], "body_bytes": 1140, "body_sha256": "sha256:2334ea6a57fa7c85f7e53ca3f8f8a3a4559f847b22157c00b91d7612780b70bf", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:no_preference", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/origin_servers/no_preference/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3222002000231230-0322310333333030-1212100212023011-3133123030311110-3111101221013101-3103003132020103-3203131120331332-2222122021111030", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers", "no_preference"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/origin_servers/no_preference/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.no_preference

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/)
- origin_servers.origin_servers.no_preference

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no preference.

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

This is an empty object or choice marker. It has no direct properties.
