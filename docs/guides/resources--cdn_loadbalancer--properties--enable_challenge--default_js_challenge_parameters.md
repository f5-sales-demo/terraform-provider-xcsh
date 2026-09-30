---
page_title: "enable_challenge.default_js_challenge_parameters"
subcategory: "Load Balancing"
description: "enable_challenge.default_js_challenge_parameters for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1029, "body_sha256": "sha256:8864fe514f9fdfc627f4b8e6551786100e8e84374b2820e40c827414babd34af", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge:default_js_challenge_parameters", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_challenge", "path": "docs/guides/resources--cdn_loadbalancer--properties--enable_challenge--default_js_challenge_parameters.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_challenge", "default_js_challenge_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/enable_challenge/default_js_challenge_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_challenge.default_js_challenge_parameters for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# enable_challenge.default_js_challenge_parameters

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [enable_challenge](resources--cdn_loadbalancer--properties--enable_challenge.md)
- enable_challenge.default_js_challenge_parameters

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default js challenge parameters.

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
default_js_challenge_parameters = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [enable_challenge](resources--cdn_loadbalancer--properties--enable_challenge.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
