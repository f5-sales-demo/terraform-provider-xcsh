---
page_title: "caching_policy"
subcategory: "Load Balancing"
description: "Caching Policies for the CDN."
xcsh_docs: {"aliases": ["caching policy"], "body_bytes": 2269, "body_sha256": "sha256:862559415a7d4b68e312d1242e89bc3eb6476d33ef233e8b1cfb777fbf1a2ee2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/caching_policy/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["caching_policy"], "schema_version": 1, "sections": [{"aliases": ["custom cache rule"], "anchor": "section", "description": "Caching policies for CDN.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["caching_policy", "custom_cache_rule"], "syntax": "block", "type": "object"}, {"aliases": ["default cache action"], "anchor": "section", "description": "This defines a Default Cache Action.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-caching_policy--default_cache_action--cache_ttl_default", "enforcement": "provider-schema", "group": "caching_policy.default_cache_action:ConflictingObjectAttributes:cache_disabled,cache_ttl_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action", "type": "conflicts"}, {"anchor": "schema-caching_policy--default_cache_action--cache_ttl_default", "enforcement": "provider-schema", "group": "caching_policy.default_cache_action:ConflictingObjectAttributes:cache_ttl_default,cache_ttl_override", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action", "type": "conflicts"}, {"anchor": "schema-caching_policy--default_cache_action--cache_ttl_override", "enforcement": "provider-schema", "group": "caching_policy.default_cache_action:ConflictingObjectAttributes:cache_disabled,cache_ttl_override", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action", "type": "conflicts"}, {"anchor": "schema-caching_policy--default_cache_action--cache_ttl_override", "enforcement": "provider-schema", "group": "caching_policy.default_cache_action:ConflictingObjectAttributes:cache_ttl_default,cache_ttl_override", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "caching_policy.default_cache_action:ConflictingObjectAttributes:cache_disabled,cache_ttl_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action:cache_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "caching_policy.default_cache_action:ConflictingObjectAttributes:cache_disabled,cache_ttl_override", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:default_cache_action:cache_disabled", "type": "conflicts"}], "schema_path": ["caching_policy", "default_cache_action"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/caching_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Caching Policies for the CDN.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# caching_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- caching_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: caching\_policy, disable\_caching; Default: disable\_caching\] Policy configuration for
this feature.

Upstream description:

Caching Policies for the CDN.

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

OneOf alternatives in this subsection:

- [caching_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/#section)
- [disable_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/disable_caching/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
caching_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/custom_cache_rule/): complete subsection reference.

- [default_cache_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/default_cache_action/): complete subsection reference.

## Next pages

- [caching_policy.custom_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/custom_cache_rule/)
- [caching_policy.default_cache_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/caching_policy/default_cache_action/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
