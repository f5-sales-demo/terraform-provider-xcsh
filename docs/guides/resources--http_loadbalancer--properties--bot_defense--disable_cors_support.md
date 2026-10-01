---
page_title: "bot_defense.disable_cors_support"
subcategory: "Load Balancing"
description: "bot_defense.disable_cors_support for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1030, "body_sha256": "sha256:1a70a608c47fec5fc60e4cd911675b56df891d679a889435be5f9971a396da8f", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:disable_cors_support", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:disable_cors_support", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense", "path": "docs/guides/resources--http_loadbalancer--properties--bot_defense--disable_cors_support.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "disable_cors_support"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/disable_cors_support/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.disable_cors_support for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.disable_cors_support

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [bot_defense](resources--http_loadbalancer--properties--bot_defense.md)
- bot_defense.disable_cors_support

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
disable_cors_support = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [bot_defense](resources--http_loadbalancer--properties--bot_defense.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
