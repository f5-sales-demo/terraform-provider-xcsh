---
page_title: "app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity"
subcategory: ""
description: "When L7 policy rules are set up to disallow certain types of requests, the system monitors persistent attempts from a user to send requests which result in policy denies. These settings specify how to use disallowed request activity from a user to determine suspicious behavior."
xcsh_docs: {"aliases": ["app type settings user behavior analysis setting enable detection include forbidden activity"], "body_bytes": 3678, "body_sha256": "sha256:512333e3a8488d5a871a9f7a109f67a252112f55e57d140d947bd38bc11ce712", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_forbidden_activity", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "path": "documentation/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_forbidden_activity/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1121232011330112-3032012310222130-3332310022303323-2003302020201303-0202221021100123-0200112032023313-3203111302100302-0011332203013210", "registry_path": "docs/guides/resources--app_setting--reference--group-001.md", "relationships": [{"anchor": "schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_forbidden_activity--forbidden_requests_threshold", "enforcement": "provider-schema", "group": "app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity:RequiredObjectAttributes:forbidden_requests_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_forbidden_activity", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_forbidden_activity"], "schema_version": 1, "sections": [{"aliases": ["app type settings user behavior analysis setting enable detection include forbidden activity forbidden requests threshold"], "anchor": "schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_forbidden_activity--forbidden_requests_threshold", "description": "The number of forbidden requests beyond which the system will flag this user as malicious.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_forbidden_activity", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_forbidden_activity", "forbidden_requests_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_forbidden_activity/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "When L7 policy rules are set up to disallow certain types of requests, the system monitors persistent attempts from a user to send requests which result in policy denies. These settings specify how to use disallowed request activity from a user to determine suspicious behavior.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_settingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/)
- [app_type_settings.user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

When L7 policy rules are set up to disallow certain types of requests, the system monitors
persistent attempts from a user to send requests which result in policy denies. These settings
specify how to use disallowed request activity from a user to determine suspicious behavior.

Upstream description:

When L7 policy rules are set up to disallow certain types of requests, the system monitors
persistent attempts from a user to send requests which result in policy denies. These settings
specify how to use disallowed request activity from a user to determine suspicious behavior.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forbidden_requests_threshold")}
```

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
include_forbidden_activity {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_forbidden_activity--forbidden_requests_threshold"></a>

### forbidden_requests_threshold property

Type: `"number"`. Optional.

The number of forbidden requests beyond which the system will flag this user as malicious.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

## Next pages

- [app_type_settings.user_behavior_analysis_setting.enable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
