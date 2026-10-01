---
page_title: "cache_rules.eligible_for_cache"
subcategory: ""
description: "cache_rules.eligible_for_cache for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1959, "body_sha256": "sha256:16bba29825d21b3e89a6c9b783dc0025dfa62b4025a936715d29a534db563cca", "canonical_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "child_ids": ["xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri"], "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "parent_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules", "path": "docs/guides/resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cache_rules", "eligible_for_cache"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_rules.eligible_for_cache for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.eligible_for_cache

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
- [Property reference](resources--cdn_cache_rule--reference.md)
- [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md)
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

- [scheme_proxy_host_request_uri](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_request_uri.md): complete subsection reference.

- [scheme_proxy_host_uri](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_uri.md): complete subsection reference.

## Next pages

- [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_request_uri.md)
- [cache_rules.eligible_for_cache.scheme_proxy_host_uri](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_uri.md)
- [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
