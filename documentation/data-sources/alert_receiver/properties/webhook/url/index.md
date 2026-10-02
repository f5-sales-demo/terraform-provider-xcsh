---
page_title: "webhook.url"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["webhook url"], "body_bytes": 1853, "body_sha256": "sha256:8bc0fe7cd057d4f5dbcaf2a2a77dbb49c3af8dbb66c1c17cd280b1ce0bf344a3", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:url:blindfold_secret_info", "xcsh-docs:data-sources:alert_receiver:properties:webhook:url:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:url", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook", "path": "documentation/data-sources/alert_receiver/properties/webhook/url/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3320213310121301-2102313133202001-1222223320030002-0210110101002111-2033020200031033-0021031212201202-1201202021022300-3332013032320210", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "url"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:url:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "url", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:url:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "url", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.url

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/)
- webhook.url

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/url/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/url/clear_secret_info/): complete subsection reference.

## Next pages

- [webhook.url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/url/blindfold_secret_info/)
- [webhook.url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/url/clear_secret_info/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
