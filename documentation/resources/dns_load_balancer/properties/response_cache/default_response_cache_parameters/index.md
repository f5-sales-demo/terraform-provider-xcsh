---
page_title: "response_cache.default_response_cache_parameters"
subcategory: "DNS"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["response cache default response cache parameters"], "body_bytes": 1388, "body_sha256": "sha256:86910c7aa586f2f2d7c93dc402c5372f046e54e7afb450b35b3577132e1ab10a", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "parent_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache", "path": "documentation/resources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2032203223131302-1320313200321200-1233201312003033-3331221203011330-0211113110211120-0010000031210201-1303120232222020-1030012203022232", "registry_path": "docs/guides/resources--dns_load_balancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["response_cache", "default_response_cache_parameters"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
default_response_cache_parameters = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
