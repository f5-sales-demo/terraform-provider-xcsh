---
page_title: "api_protection_rules.api_endpoint_rules.action.deny"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["api protection rules api endpoint rules action deny"], "body_bytes": 1773, "body_sha256": "sha256:fa383eb0c5de3cd4037459a9aa5e78e929443ff4689c13f2eea1c83f3d83007f", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action:deny", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action", "path": "documentation/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/deny/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1113222113302312-2313111022101103-1021202102123323-0112000032231220-1321000223003233-3221023213330312-0121222022020232-2232011000123330", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_endpoint_rules", "action", "deny"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/deny/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_endpoint_rules.action.deny

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/)
- [api_protection_rules.api_endpoint_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/)
- api_protection_rules.api_endpoint_rules.action.deny

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
deny = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_protection_rules.api_endpoint_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
