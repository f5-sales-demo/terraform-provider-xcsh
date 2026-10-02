---
page_title: "email"
subcategory: ""
description: "Email Configuration."
xcsh_docs: {"aliases": ["email"], "body_bytes": 3005, "body_sha256": "sha256:304a0d4eddd2de7c450909befc3476c0d7ac320efc8049108cac97c11d35fa7c", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:email", "parent_id": "xcsh-docs:resources:alert_receiver:reference", "path": "documentation/resources/alert_receiver/properties/email/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1323022032232303-1330320332312203-2011200000313323-0131203103332132-0201310302130131-1010310000301301-3032213102033322-2120230210032032", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["email"], "schema_version": 1, "sections": [{"aliases": ["email"], "anchor": "schema-email--email", "description": "Email ID of the user.", "document_id": "xcsh-docs:resources:alert_receiver:properties:email", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["email", "email"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/email/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Email Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# email

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- email

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: email, opsgenie, pagerduty, slack, sms, webhook\] Email Configuration.

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

OneOf alternatives in this subsection:

- [email](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/email/#section)
- [opsgenie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/#section)
- [pagerduty](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/#section)
- [slack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/#section)
- [sms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/sms/#section)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
email {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-email--email"></a>

### email property

Type: `"string"`. Optional.

Email. Email ID of the user.

Upstream description:

Email ID of the user.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "email",
    "formatDescription": "RFC 5322 email address, max 254 characters",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
    "validation": {
      "rfc": "RFC 5322"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
