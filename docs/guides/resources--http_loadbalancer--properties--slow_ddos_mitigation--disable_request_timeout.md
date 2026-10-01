---
page_title: "slow_ddos_mitigation.disable_request_timeout"
subcategory: "Load Balancing"
description: "slow_ddos_mitigation.disable_request_timeout for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1127, "body_sha256": "sha256:2ee26e71899f8b1e354a56d425c91805807be0f4c69bdd56a54db774fe87f113", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation:disable_request_timeout", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation:disable_request_timeout", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation", "path": "docs/guides/resources--http_loadbalancer--properties--slow_ddos_mitigation--disable_request_timeout.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["slow_ddos_mitigation", "disable_request_timeout"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/slow_ddos_mitigation/disable_request_timeout/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "slow_ddos_mitigation.disable_request_timeout for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slow_ddos_mitigation.disable_request_timeout

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [slow_ddos_mitigation](resources--http_loadbalancer--properties--slow_ddos_mitigation.md)
- slow_ddos_mitigation.disable_request_timeout

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable request timeout.

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
disable_request_timeout = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [slow_ddos_mitigation](resources--http_loadbalancer--properties--slow_ddos_mitigation.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
