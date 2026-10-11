---
page_title: "response_cache.default_response_cache_parameters"
subcategory: "DNS"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["response cache default response cache parameters"], "body_bytes": 1113, "body_sha256": "sha256:98aa2da99ff62dd0da85f1665b29461cf13137d9102c514cf5b81cf87e4fae06", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "parent_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache", "path": "documentation/resources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-2032203223131302-1320313200321200-1233201312003033-3331221203011330-0211113110211120-0010000031210201-1303120232222020-1030012203022232", "registry_path": "docs/guides/resources--dns_load_balancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["response_cache", "default_response_cache_parameters"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_cache.default_response_cache_parameters

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/)
- [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/)
- response_cache.default_response_cache_parameters

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default response cache parameters.

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
default_response_cache_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.
