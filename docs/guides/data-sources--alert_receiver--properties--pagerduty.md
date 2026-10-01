---
page_title: "pagerduty"
subcategory: ""
description: "pagerduty for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2033, "body_sha256": "sha256:f9c66020d96739c47a3f497df4f8eb1d557fcfd87034bc623807d3485128833b", "canonical_id": "xcsh-docs:data-sources:alert_receiver:properties:pagerduty", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:pagerduty:routing_key"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:pagerduty", "parent_id": "xcsh-docs:data-sources:alert_receiver:reference", "path": "docs/guides/data-sources--alert_receiver--properties--pagerduty.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["pagerduty"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/pagerduty/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "pagerduty for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# pagerduty

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- pagerduty

<a id="section"></a>

Type: `"single"`. Computed.

PagerDuty configuration to send alert notifications.

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

## Direct properties

- [routing_key](data-sources--alert_receiver--properties--pagerduty--routing_key.md): complete subsection reference.

<a id="schema-pagerduty--url"></a>

### url property

Type: `"string"`. Computed.

Pager Duty URL. URL to send API requests to.

Upstream description:

URL to send API requests to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [pagerduty.routing_key](data-sources--alert_receiver--properties--pagerduty--routing_key.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
