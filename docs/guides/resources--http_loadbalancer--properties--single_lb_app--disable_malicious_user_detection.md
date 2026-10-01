---
page_title: "single_lb_app.disable_malicious_user_detection"
subcategory: "Load Balancing"
description: "single_lb_app.disable_malicious_user_detection for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1121, "body_sha256": "sha256:6bc877f4679fd13e7d2e135a99b86c925b4ed36d4c882063fc728b1d9dda5253", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_malicious_user_detection", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:disable_malicious_user_detection", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app", "path": "docs/guides/resources--http_loadbalancer--properties--single_lb_app--disable_malicious_user_detection.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app", "disable_malicious_user_detection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/disable_malicious_user_detection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.disable_malicious_user_detection for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.disable_malicious_user_detection

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [single_lb_app](resources--http_loadbalancer--properties--single_lb_app.md)
- single_lb_app.disable_malicious_user_detection

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable malicious user detection.

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
disable_malicious_user_detection = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [single_lb_app](resources--http_loadbalancer--properties--single_lb_app.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
