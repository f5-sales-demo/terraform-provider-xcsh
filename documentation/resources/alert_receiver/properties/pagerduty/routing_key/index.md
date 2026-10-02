---
page_title: "pagerduty.routing_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["pagerduty routing key"], "body_bytes": 2222, "body_sha256": "sha256:e6046e9471dc8aea6b21ec12f0d56f3f0be5933572dc3b969eff87b2593ab735", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key:blindfold_secret_info", "xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key", "parent_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty", "path": "documentation/resources/alert_receiver/properties/pagerduty/routing_key/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0102322223320031-2012233012020231-3313030010322021-0112222332120010-1020101123230322-1002000011100003-2302021223200211-3202103010331111", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "pagerduty.routing_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pagerduty.routing_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["pagerduty", "routing_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-pagerduty--routing_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "pagerduty.routing_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key:blindfold_secret_info", "type": "requires"}], "schema_path": ["pagerduty", "routing_key", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-pagerduty--routing_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "pagerduty.routing_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key:clear_secret_info", "type": "requires"}], "schema_path": ["pagerduty", "routing_key", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/pagerduty/routing_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# pagerduty.routing_key

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [pagerduty](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/)
- pagerduty.routing_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
routing_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/clear_secret_info/): complete subsection reference.

## Next pages

- [pagerduty.routing_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/blindfold_secret_info/)
- [pagerduty.routing_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/clear_secret_info/)
- [pagerduty](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
