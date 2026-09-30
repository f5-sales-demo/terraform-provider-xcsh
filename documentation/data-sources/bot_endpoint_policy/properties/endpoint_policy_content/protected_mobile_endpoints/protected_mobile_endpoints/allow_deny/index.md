---
page_title: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.allow_deny"
subcategory: ""
description: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.allow_deny for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2449, "body_sha256": "sha256:3e815834abf068fde87161c20faad3d3db3345b09a4544293fefa3cc7ab7193c", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:allow_deny", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/allow_deny/index.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "allow_deny"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/allow_deny/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.allow_deny for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.allow_deny

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.allow_deny

<a id="section"></a>

Type: `"single"`. Computed.

Known Bot Allow And/Or Deny Action Type.

## Direct properties

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--allow_deny--allow_list"></a>

### allow_list property

Type: `["list", "string"]`. Computed.

Select Known Bots to allow to proceed to the origin.

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--allow_deny--deny_list"></a>

### deny_list property

Type: `["list", "string"]`. Computed.

Deny list actions will only take effect when the Mitigation Action above is set (e.g., block,
redirect, transform). If mitigation action above is set to Continue, Known bots will be flagged.

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--allow_deny--text_block"></a>

### text_block property

Type: `"string"`. Computed.

All others (not in Allow or Deny). Blocking or denial configuration

## Next pages

- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
