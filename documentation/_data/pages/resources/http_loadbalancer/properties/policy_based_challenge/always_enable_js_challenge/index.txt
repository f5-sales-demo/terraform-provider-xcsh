---
page_title: "policy_based_challenge.always_enable_js_challenge"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["policy based challenge always enable js challenge"], "body_bytes": 1408, "body_sha256": "sha256:1be6f7f3b8b30a83bb9673278863bd7da6f055c3571d75dea1c562365a74a3c9", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:always_enable_js_challenge", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge", "path": "documentation/resources/http_loadbalancer/properties/policy_based_challenge/always_enable_js_challenge/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2020101203121101-3122131030121332-0001233010301221-2210011313322031-2230000010103111-2322322320133211-3211101333203323-0113330131100123", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-022.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "always_enable_js_challenge"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/always_enable_js_challenge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.always_enable_js_challenge

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/)
- policy_based_challenge.always_enable_js_challenge

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable js challenge.

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
always_enable_js_challenge = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
