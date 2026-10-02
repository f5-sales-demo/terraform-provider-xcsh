---
page_title: "cache_rules.eligible_for_cache"
subcategory: ""
description: "List of OPTIONS for Cache Action."
xcsh_docs: {"aliases": ["cache rules eligible for cache"], "body_bytes": 2412, "body_sha256": "sha256:0e1e3f23e61337733e6b6dc117d9b313b8925b155bbbfc19448ab5f9573230c1", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "parent_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules", "path": "documentation/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3223232310102321-3003333113231003-0113233100223033-1002000121033202-2010331330331313-3230301030101011-2002313003222331-0320232110002100", "registry_path": "docs/guides/resources--cdn_cache_rule--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cache_rules.eligible_for_cache:ConflictingObjectAttributes:scheme_proxy_host_request_uri,scheme_proxy_host_uri", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cache_rules.eligible_for_cache:ConflictingObjectAttributes:scheme_proxy_host_request_uri,scheme_proxy_host_uri", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules", "eligible_for_cache"], "schema_version": 1, "sections": [{"aliases": ["scheme proxy host request uri"], "anchor": "section", "description": "Cache TTL Enable Values.", "document_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--cache_ttl", "enforcement": "provider-schema", "group": "cache_rules.eligible_for_cache.scheme_proxy_host_request_uri:RequiredObjectAttributes:cache_ttl", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "type": "requires"}], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_request_uri"], "syntax": "block", "type": "object"}, {"aliases": ["scheme proxy host uri"], "anchor": "section", "description": "Cache TTL Enable Values.", "document_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_ttl", "enforcement": "provider-schema", "group": "cache_rules.eligible_for_cache.scheme_proxy_host_uri:RequiredObjectAttributes:cache_ttl", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "type": "requires"}], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of OPTIONS for Cache Action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.eligible_for_cache

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/)
- cache_rules.eligible_for_cache

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eligible for cache.

Upstream description:

List of OPTIONS for Cache Action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("scheme_proxy_host_request_uri",
    "scheme_proxy_host_uri")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-eligible_for_cache": "[\"scheme_proxy_host_request_uri\",\"scheme_proxy_host_uri\"]"
}
```

Terraform syntax:

```terraform
eligible_for_cache {
  # Configure direct properties listed below.
}
```

## Direct properties

- [scheme_proxy_host_request_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_request_uri/): complete subsection reference.

- [scheme_proxy_host_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/): complete subsection reference.

## Next pages

- [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_request_uri/)
- [cache_rules.eligible_for_cache.scheme_proxy_host_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
