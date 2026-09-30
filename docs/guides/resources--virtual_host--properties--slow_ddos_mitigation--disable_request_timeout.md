---
page_title: "slow_ddos_mitigation.disable_request_timeout"
subcategory: ""
description: "slow_ddos_mitigation.disable_request_timeout for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 993, "body_sha256": "sha256:90032d94e1ebfc444e2ae5cbfd7bf04d1c6119d2aaa4b2b703d7317b1c6fa9c5", "canonical_id": "xcsh-docs:resources:virtual_host:properties:slow_ddos_mitigation:disable_request_timeout", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:slow_ddos_mitigation:disable_request_timeout", "parent_id": "xcsh-docs:resources:virtual_host:properties:slow_ddos_mitigation", "path": "docs/guides/resources--virtual_host--properties--slow_ddos_mitigation--disable_request_timeout.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["slow_ddos_mitigation", "disable_request_timeout"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/slow_ddos_mitigation/disable_request_timeout/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "slow_ddos_mitigation.disable_request_timeout for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# slow_ddos_mitigation.disable_request_timeout

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [slow_ddos_mitigation](resources--virtual_host--properties--slow_ddos_mitigation.md)
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

- [slow_ddos_mitigation](resources--virtual_host--properties--slow_ddos_mitigation.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
