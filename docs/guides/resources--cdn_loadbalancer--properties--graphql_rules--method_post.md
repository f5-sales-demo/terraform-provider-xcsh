---
page_title: "graphql_rules.method_post"
subcategory: "Load Balancing"
description: "graphql_rules.method_post for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1030, "body_sha256": "sha256:bc5c5559218d8eff124166609b003f7160da2adfafc0165b2a3712dd6fe1c5eb", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:method_post", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:method_post", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules", "path": "docs/guides/resources--cdn_loadbalancer--properties--graphql_rules--method_post.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["graphql_rules", "method_post"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/graphql_rules/method_post/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "graphql_rules.method_post for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# graphql_rules.method_post

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [graphql_rules](resources--cdn_loadbalancer--properties--graphql_rules.md)
- graphql_rules.method_post

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for method post.

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
method_post = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [graphql_rules](resources--cdn_loadbalancer--properties--graphql_rules.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
