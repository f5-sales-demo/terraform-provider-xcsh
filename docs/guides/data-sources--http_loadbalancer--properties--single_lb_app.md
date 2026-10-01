---
page_title: "single_lb_app"
subcategory: "Load Balancing"
description: "single_lb_app for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2133, "body_sha256": "sha256:d6db2a3cf8fbd62d53872142dde906fc2e0957688199a7251b2e1ae1353fbbba", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:disable_discovery", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:disable_malicious_user_detection", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_malicious_user_detection"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--single_lb_app.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/single_lb_app/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- single_lb_app

<a id="section"></a>

Type: `"single"`. Computed.

Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_choice": "[\"disable_discovery\",\"enable_discovery\"]",
  "x-ves-oneof-field-malicious_user_detection_choice": "[\"disable_malicious_user_detection\",\"enable_malicious_user_detection\"]"
}
```

## Direct properties

- [disable_discovery](data-sources--http_loadbalancer--properties--single_lb_app--disable_discovery.md): complete subsection reference.

- [disable_malicious_user_detection](data-sources--http_loadbalancer--properties--single_lb_app--disable_malicious_user_detection.md): complete subsection reference.

- [enable_discovery](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery.md): complete subsection reference.

- [enable_malicious_user_detection](data-sources--http_loadbalancer--properties--single_lb_app--enable_malicious_user_detection.md): complete subsection reference.

## Next pages

- [single_lb_app.disable_discovery](data-sources--http_loadbalancer--properties--single_lb_app--disable_discovery.md)
- [single_lb_app.disable_malicious_user_detection](data-sources--http_loadbalancer--properties--single_lb_app--disable_malicious_user_detection.md)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- [single_lb_app.enable_malicious_user_detection](data-sources--http_loadbalancer--properties--single_lb_app--enable_malicious_user_detection.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
