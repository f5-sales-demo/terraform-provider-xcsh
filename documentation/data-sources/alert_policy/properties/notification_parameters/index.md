---
page_title: "notification_parameters"
subcategory: "Monitoring"
description: "Set of notification parameters to decide how and when the alert notifications should be sent to the receivers."
xcsh_docs: {"aliases": ["notification parameters"], "body_bytes": 6990, "body_sha256": "sha256:bece170ea16596c07030ec95760229937dfae5560e017289298b455c61474d4c", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_policy:properties:notification_parameters:custom", "xcsh-docs:data-sources:alert_policy:properties:notification_parameters:default", "xcsh-docs:data-sources:alert_policy:properties:notification_parameters:individual", "xcsh-docs:data-sources:alert_policy:properties:notification_parameters:ves_io_group"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters", "parent_id": "xcsh-docs:data-sources:alert_policy:reference", "path": "documentation/data-sources/alert_policy/properties/notification_parameters/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200", "registry_path": "docs/guides/data-sources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["notification_parameters"], "schema_version": 1, "sections": [{"aliases": ["custom"], "anchor": "section", "description": "Specify list of custom labels to group/aggregate the alerts.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters:custom", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["notification_parameters", "custom"], "syntax": "attribute", "type": "object"}, {"aliases": ["default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters:default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "default"], "syntax": "attribute", "type": "object"}, {"aliases": ["group interval"], "anchor": "schema-notification_parameters--group_interval", "description": "Group Interval is used to specify how long to wait before sending a notification about new alerts that are added to the group for which an initial notification has already been sent. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_interval defaults to \"1m\"", "document_id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "group_interval"], "syntax": "attribute", "type": "string"}, {"aliases": ["group wait"], "anchor": "schema-notification_parameters--group_wait", "description": "Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect more alerts for the same group. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_wait defaults to \"30s\"", "document_id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "group_wait"], "syntax": "attribute", "type": "string"}, {"aliases": ["individual"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters:individual", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "individual"], "syntax": "attribute", "type": "object"}, {"aliases": ["login success", "repeat interval", "succeeded", "success", "successful"], "anchor": "schema-notification_parameters--repeat_interval", "description": "Repeat Interval is used to specify how long to wait before sending a notification again if it has already been sent successfully. Format: , where s - seconds, m - minutes, h - hours, d - days If not specified, group_interval defaults to \"4h\"", "document_id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "repeat_interval"], "syntax": "attribute", "type": "string"}, {"aliases": ["ves io group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters:ves_io_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "ves_io_group"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/notification_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Set of notification parameters to decide how and when the alert notifications should be sent to the receivers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# notification_parameters

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/)
- notification_parameters

<a id="section"></a>

Type: `"single"`. Computed.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

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

## Direct properties

- [custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/custom/): complete subsection reference.

- [default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/default/): complete subsection reference.

<a id="schema-notification_parameters--group_interval"></a>

### group_interval property

Type: `"string"`. Computed.

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval..

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-notification_parameters--group_wait"></a>

### group_wait property

Type: `"string"`. Computed.

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to '30s'.

Upstream description:

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to "30s"

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [individual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/individual/): complete subsection reference.

<a id="schema-notification_parameters--repeat_interval"></a>

### repeat_interval property

Type: `"string"`. Computed.

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to '4h'.

Upstream description:

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to "4h"

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [ves_io_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/ves_io_group/): complete subsection reference.

## Next pages

- [notification_parameters.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/custom/)
- [notification_parameters.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/default/)
- [notification_parameters.individual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/individual/)
- [notification_parameters.ves_io_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/ves_io_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
