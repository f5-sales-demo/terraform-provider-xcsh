---
page_title: "slack.url"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["slack url"], "body_bytes": 1829, "body_sha256": "sha256:f72881f6d62f9cc67a459a56ae25bcc0db36d98f28130ca427d9d2b98d41e5cf", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:slack:url:blindfold_secret_info", "xcsh-docs:data-sources:alert_receiver:properties:slack:url:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:slack:url", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:slack", "path": "documentation/data-sources/alert_receiver/properties/slack/url/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1010221222222122-1232032003021211-0132111211200300-1230203023100021-3333313011300332-1123102031101110-1303333212103120-3133200331231032", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["slack", "url"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:slack:url:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["slack", "url", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:slack:url:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["slack", "url", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/slack/url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slack.url

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
- [slack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/slack/)
- slack.url

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/slack/url/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/slack/url/clear_secret_info/): complete subsection reference.

## Next pages

- [slack.url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/slack/url/blindfold_secret_info/)
- [slack.url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/slack/url/clear_secret_info/)
- [slack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/slack/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
