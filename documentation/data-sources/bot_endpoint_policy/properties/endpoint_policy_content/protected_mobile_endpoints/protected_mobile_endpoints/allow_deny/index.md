---
page_title: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.allow_deny"
subcategory: ""
description: "Known Bot Allow And/Or Deny Action Type."
xcsh_docs: {"aliases": ["endpoint policy content protected mobile endpoints protected mobile endpoints allow deny"], "body_bytes": 2150, "body_sha256": "sha256:66c2eb36b8db4bd055120baf860ec6b236569688a189e4d9598cbd72905619dd", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:allow_deny", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/allow_deny/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3212123123010102-0100230131010101-3210021113200111-3021221303222100-0323300312130122-1132113110010102-1311111313211233-0310103231133212", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "allow_deny"], "schema_version": 1, "sections": [{"aliases": ["endpoint policy content protected mobile endpoints protected mobile endpoints allow deny allow list"], "anchor": "schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--allow_deny--allow_list", "description": "Select Known Bots to allow to proceed to the origin.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:allow_deny", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "allow_deny", "allow_list"], "syntax": "attribute", "type": "list"}, {"aliases": ["endpoint policy content protected mobile endpoints protected mobile endpoints allow deny deny list"], "anchor": "schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--allow_deny--deny_list", "description": "Deny list actions will only take effect when the Mitigation Action above is set (e.g., block, redirect, transform). If mitigation action above is set to Continue, Known bots will be flagged.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:allow_deny", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "allow_deny", "deny_list"], "syntax": "attribute", "type": "list"}, {"aliases": ["endpoint policy content protected mobile endpoints protected mobile endpoints allow deny text block"], "anchor": "schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--allow_deny--text_block", "description": "All others (not in Allow or Deny). Blocking or denial configuration", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:allow_deny", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "allow_deny", "text_block"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/allow_deny/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Known Bot Allow And/Or Deny Action Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
