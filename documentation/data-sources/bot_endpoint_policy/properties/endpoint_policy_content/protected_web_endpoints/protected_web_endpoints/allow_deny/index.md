---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.allow_deny"
subcategory: ""
description: "Known Bot Allow And/Or Deny Action Type."
xcsh_docs: {"aliases": ["endpoint policy content protected web endpoints protected web endpoints allow deny"], "body_bytes": 2488, "body_sha256": "sha256:c7b41bd100d91241f848b6070d94088c8fcea3afdeb2ce4506e435f27c5f2f6f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:allow_deny", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/allow_deny/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-0112323300031323-1133310330120023-0022330320121312-0312310301120221-2010010331103201-1322133002212002-3132021321333233-2012132131230200", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "allow_deny"], "schema_version": 1, "sections": [{"aliases": ["allow list", "backend servers", "origin servers", "upstream servers"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--allow_deny--allow_list", "description": "Select Known Bots to allow to proceed to the origin.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:allow_deny", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "allow_deny", "allow_list"], "syntax": "attribute", "type": "list"}, {"aliases": ["deny list"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--allow_deny--deny_list", "description": "Deny list actions will only take effect when the Mitigation Action above is set (e.g., block, redirect, transform). If mitigation action above is set to Continue, Known bots will be flagged.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:allow_deny", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "allow_deny", "deny_list"], "syntax": "attribute", "type": "list"}, {"aliases": ["text block"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--allow_deny--text_block", "description": "All others (not in Allow or Deny). Blocking or denial configuration", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:allow_deny", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "allow_deny", "text_block"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/allow_deny/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Known Bot Allow And/Or Deny Action Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.allow_deny

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.allow_deny

<a id="section"></a>

Type: `"single"`. Computed.

Known Bot Allow And/Or Deny Action Type.

## Direct properties

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--allow_deny--allow_list"></a>

### allow_list property

Type: `["list", "string"]`. Computed.

Select Known Bots to allow to proceed to the origin.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--allow_deny--deny_list"></a>

### deny_list property

Type: `["list", "string"]`. Computed.

Deny list actions will only take effect when the Mitigation Action above is set (e.g., block,
redirect, transform). If mitigation action above is set to Continue, Known bots will be flagged.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--allow_deny--text_block"></a>

### text_block property

Type: `"string"`. Computed.

All others (not in Allow or Deny). Blocking or denial configuration

## Next pages

- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
