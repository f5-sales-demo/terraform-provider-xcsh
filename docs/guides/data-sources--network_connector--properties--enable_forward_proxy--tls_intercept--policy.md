---
page_title: "enable_forward_proxy.tls_intercept.policy"
subcategory: "Networking"
description: "enable_forward_proxy.tls_intercept.policy for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1304, "body_sha256": "sha256:0c1f1dbb5c0c8f20667fba15f3bda8042ad0534eb1c03eff5305ce867436fbbb", "canonical_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules"], "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy", "parent_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept", "path": "docs/guides/data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_forward_proxy.tls_intercept.policy for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# enable_forward_proxy.tls_intercept.policy

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md)
- [Property reference](data-sources--network_connector--reference.md)
- [enable_forward_proxy](data-sources--network_connector--properties--enable_forward_proxy.md)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept.md)
- enable_forward_proxy.tls_intercept.policy

<a id="section"></a>

Type: `"single"`. Computed.

Policy to enable or disable TLS interception.

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

## Direct properties

- [interception_rules](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules.md): complete subsection reference.

## Next pages

- [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules.md)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept.md)
- [xcsh_network_connector](../data-sources/network_connector.md)
