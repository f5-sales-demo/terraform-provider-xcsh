---
page_title: "bot_defense.policy.disable_js_insert"
subcategory: "Load Balancing"
description: "bot_defense.policy.disable_js_insert for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1374, "body_sha256": "sha256:6c30d5dd53b9c7184f34cf34e43dbcd4423a5fc431f5e390cf8b760e88273b76", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:disable_js_insert", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy", "path": "documentation/resources/http_loadbalancer/properties/bot_defense/policy/disable_js_insert/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["bot_defense", "policy", "disable_js_insert"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/disable_js_insert/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.disable_js_insert for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.disable_js_insert

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- bot_defense.policy.disable_js_insert

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
