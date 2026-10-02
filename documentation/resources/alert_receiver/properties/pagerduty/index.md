---
page_title: "pagerduty"
subcategory: ""
description: "PagerDuty configuration to send alert notifications."
xcsh_docs: {"aliases": ["pagerduty"], "body_bytes": 2726, "body_sha256": "sha256:7e565b096ce1c5f6a8b98f30d77e3bc169c3ab9ff69e57e6fc5ae81c67cd6d15", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:pagerduty", "parent_id": "xcsh-docs:resources:alert_receiver:reference", "path": "documentation/resources/alert_receiver/properties/pagerduty/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2030011230032101-0030332021230211-0222131321322313-3023232233132332-1112230010122010-3103113212312200-1332122030201102-0313302230011001", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [{"anchor": "schema-pagerduty--url", "enforcement": "provider-schema", "group": "pagerduty:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["pagerduty"], "schema_version": 1, "sections": [{"aliases": ["routing key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "pagerduty.routing_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "pagerduty.routing_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty:routing_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["pagerduty", "routing_key"], "syntax": "block", "type": "object"}, {"aliases": ["url"], "anchor": "schema-pagerduty--url", "description": "URL to send API requests to.", "document_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pagerduty", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/pagerduty/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "PagerDuty configuration to send alert notifications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# pagerduty

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- pagerduty

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

PagerDuty configuration to send alert notifications.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
pagerduty {
  # Configure direct properties listed below.
}
```

## Direct properties

- [routing_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/): complete subsection reference.

<a id="schema-pagerduty--url"></a>

### url property

Type: `"string"`. Optional.

Pager Duty URL. URL to send API requests to.

Upstream description:

URL to send API requests to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

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

- [pagerduty.routing_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
