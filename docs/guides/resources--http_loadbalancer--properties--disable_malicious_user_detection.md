---
page_title: "disable_malicious_user_detection"
subcategory: "Load Balancing"
description: "disable_malicious_user_detection for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1557, "body_sha256": "sha256:f1511ffbe2be1ca961c1e0191736a2dd8d05fb0e0aef04a584d346d6194244da", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:disable_malicious_user_detection", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:disable_malicious_user_detection", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--disable_malicious_user_detection.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_malicious_user_detection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/disable_malicious_user_detection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_malicious_user_detection for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_malicious_user_detection

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- disable_malicious_user_detection

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_malicious_user_detection](resources--http_loadbalancer--properties--disable_malicious_user_detection.md#section)
- [enable_malicious_user_detection](resources--http_loadbalancer--properties--enable_malicious_user_detection.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malicious_user_detection = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
