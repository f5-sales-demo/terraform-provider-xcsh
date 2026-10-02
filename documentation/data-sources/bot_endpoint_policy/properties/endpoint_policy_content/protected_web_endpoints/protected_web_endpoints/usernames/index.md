---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.usernames"
subcategory: ""
description: "Add the field name where you store username on the Request Body."
xcsh_docs: {"aliases": ["endpoint policy content protected web endpoints protected web endpoints usernames"], "body_bytes": 2253, "body_sha256": "sha256:75be7b30465fc87a39bb0cd3c248afca809b9c3d08e7c6245a58a8c439bae08f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:usernames", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/usernames/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0222112233032131-2201211101203102-2213020310200030-2113323003110331-3222032123312323-2132033121033333-0203200013133322-1021013221213231", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "usernames"], "schema_version": 1, "sections": [{"aliases": ["encryption type"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--usernames--encryption_type", "description": "Encryption Type for username reporting - PlainText: plaintext - Hashed: hashed. Possible values are `PlainText`, `Hashed`. Defaults to `PlainText`.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:usernames", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "usernames", "encryption_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["username reporting"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--usernames--username_reporting", "description": "Field. Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:usernames", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "usernames", "username_reporting"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/usernames/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Add the field name where you store username on the Request Body.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.usernames

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.usernames

<a id="section"></a>

Type: `"list"`. Computed.

Add the field name where you store username on the Request Body.

## Direct properties

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--usernames--encryption_type"></a>

### encryption_type property

Type: `"string"`. Computed.

\[Enum: PlainText|Hashed\] Encryption Type for username reporting - PlainText: plaintext - Hashed:
hashed. Possible values are \`PlainText\`, \`Hashed\`. Defaults to \`PlainText\`.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--usernames--username_reporting"></a>

### username_reporting property

Type: `"string"`. Computed.

Field. Human-readable name for the resource

## Next pages

- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
