---
page_title: "routes.notification_parameters"
subcategory: "Monitoring"
description: "Set of notification parameters to decide how and when the alert notifications should be sent to the receivers."
xcsh_docs: {"aliases": ["routes notification parameters"], "body_bytes": 6362, "body_sha256": "sha256:9a420345fd58a6b8deb57a08153352dbbe469efab0648eae0bf8ee8102b96614", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:custom", "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:default", "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:individual", "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:ves_io_group"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters", "parent_id": "xcsh-docs:resources:alert_policy:properties:routes", "path": "documentation/resources/alert_policy/properties/routes/notification_parameters/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1132102222000120-1312023210212021-2212103031202003-2303010133130331-2001130111201011-0311320332230320-2230201020303103-2220332311321331", "registry_path": "docs/guides/resources--alert_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:default,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:default,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:individual", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:default,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:individual", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:individual,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:individual", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:ves_io_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:default,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:ves_io_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:individual,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:ves_io_group", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "notification_parameters"], "schema_version": 1, "sections": [{"aliases": ["routes notification parameters custom"], "anchor": "section", "description": "Specify list of custom labels to group/aggregate the alerts.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:custom", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "notification_parameters", "custom"], "syntax": "block", "type": "object"}, {"aliases": ["routes notification parameters default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:default", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "default"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes notification parameters group interval"], "anchor": "schema-routes--notification_parameters--group_interval", "description": "Group Interval is used to specify how long to wait before sending a notification about new alerts that are added to the group for which an initial notification has already been sent. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_interval defaults to \"1m\"", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "group_interval"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes notification parameters group wait"], "anchor": "schema-routes--notification_parameters--group_wait", "description": "Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect more alerts for the same group. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_wait defaults to \"30s\"", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "group_wait"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes notification parameters individual"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:individual", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "individual"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes notification parameters repeat interval", "succeeded", "success", "successful"], "anchor": "schema-routes--notification_parameters--repeat_interval", "description": "Repeat Interval is used to specify how long to wait before sending a notification again if it has already been sent successfully. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_interval defaults to \"4h\"", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "repeat_interval"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes notification parameters ves io group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:ves_io_group", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "ves_io_group"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/routes/notification_parameters/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Set of notification parameters to decide how and when the alert notifications should be sent to the receivers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["alert_policyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.notification_parameters

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/)
- routes.notification_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom",
    "default"),
  validators.ConflictingObjectAttributes("custom",
    "individual"),
  validators.ConflictingObjectAttributes("custom",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("default",
    "individual"),
  validators.ConflictingObjectAttributes("default",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("individual",
    "ves_io_group")}
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
  "x-ves-oneof-field-group_by": "[\"custom\",\"default\",\"individual\",\"ves_io_group\"]"
}
```

Terraform syntax:

```terraform
notification_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/custom/): complete subsection reference.

- [default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/default/): complete subsection reference.

<a id="schema-routes--notification_parameters--group_interval"></a>

### group_interval property

Type: `"string"`. Optional.

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "1m"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="schema-routes--notification_parameters--group_wait"></a>

### group_wait property

Type: `"string"`. Optional.

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to '30s'.

Additional upstream details:

Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_wait defaults to "30s"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [individual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/individual/): complete subsection reference.

<a id="schema-routes--notification_parameters--repeat_interval"></a>

### repeat_interval property

Type: `"string"`. Optional.

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to '4h'.

Additional upstream details:

Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "4h"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [ves_io_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/ves_io_group/): complete subsection reference.
