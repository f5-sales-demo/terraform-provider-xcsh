---
page_title: "enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception"
subcategory: "Networking"
description: "enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1611, "body_sha256": "sha256:14fc05ccb60f1809e885fd47937ca67b737475ea4190b245a17294e83cfd94ea", "canonical_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:disable_interception", "child_ids": [], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules:disable_interception", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy:interception_rules", "path": "docs/guides/resources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules--disable_interception.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "policy", "interception_rules", "disable_interception"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/disable_interception/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md)
- [Property reference](resources--network_connector--reference.md)
- [enable_forward_proxy](resources--network_connector--properties--enable_forward_proxy.md)
- [enable_forward_proxy.tls_intercept](resources--network_connector--properties--enable_forward_proxy--tls_intercept.md)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--properties--enable_forward_proxy--tls_intercept--policy.md)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules.md)
- enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable interception.

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
disable_interception = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules.md)
- [xcsh_network_connector](../resources/network_connector.md)
